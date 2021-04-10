## Используем бандлы

К этому моменту у вас мог возникнуть вопрос: а как можно поделиться набором правил?

Как установить себе уже существующий набор правил?

Нужно ли копировать себе чужие файлы, устраивая своеобразный вендоринг?

- Нет, не обязательно! Мы можем использовать Go модули для этого.

Чтобы не возникало путаницы с пакетами и модулями, модули с правилами называются бандлами.

Бандл - это обычный Go пакет, оформленный как Go модуль, содержащий в себе `gorules` файлы.

### Подключаем внешние правила

В репозитории `ruleguard` можно найти небольшой бандл с правилами, который можно установить
выполнив одну простую команду:

```bash
go get -v github.com/quasilyte/go-ruleguard/rules
```

Теперь осталось импортировать бандл в нашем файле правил.

```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	corerules "github.com/quasilyte/go-ruleguard/rules"
)

func init() {
	dsl.ImportRules("", corerules.Bundle)
}

func suggestConcat(m dsl.Matcher) {
	m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
		Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
		Report("could use a simple concat for $x and $y instead of the Sprintf call")

	m.Match(`strings.Join([]string{$a, $b}, "")`).
		Report("could rewrite a strings.Join call as $a+$b")

	m.Match(`strings.Join([]string{$a, $b}, $glue)`).
		Report(`could rewrite a strings.Join call as $a+$glue+$b`)
}
```

Здесь добавился импорт на сам пакет со внешними правилами (бандл) и вызов `ImportRules` внутри `init()` для регистрации этих правил.

Первый аргумент в `ImportRules`, который в нашем случае пустая строка - это префикс, добавляемый для диагностик, определяемых в бандле.

Коллизии имён вызывают ошибку парсинга файлов правил. Поэтому при импорте можно добавлять префиксы. Это удобно и для разрешения коллизий,
и для более гранулярной организации имён диагностик. Пустую строку можно использовать если ваших собственных правил мало или их нет вовсе.

Например, вот это - тоже валидный файл правил:

```go
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
	corerules "github.com/quasilyte/go-ruleguard/rules"
)

func init() {
	dsl.ImportRules("", corerules.Bundle)
}
```

Данный файл - минимальный пример использования внешних правил без добавления каких-либо своих.

Создадим файл `target.go` для испытаний подключенного бандла:

```go
package example

import (
  "strings"
  "bytes"
)

func f(s string, b []byte, ints []int) {
	_ = strings.Count(s, "/") >= 0
	_ = bytes.Count(b, []byte("/")) >= 0
  
  sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})
}
```

Теперь попробуем запустить наш `rules.go` на новом файле:

```bash
$ ruleguard -rules rules.go target.go
target.go:10:6: badCond: statement always true (diag.go:10)
target.go:11:6: badCond: statement always true (diag.go:11)
target.go:13:2: sortFuncs: suggestion: sort.Ints(ints) (refactor.go:12)
```

Диагностики `badCond` и `sortFuncs` пришли к нам из бандла `go-ruleguard/rules`.

Давайте установим непустой префикс при регистрации правил:

```diff
- dsl.ImportRules("", corerules.Bundle)
+ dsl.ImportRules("corerules", corerules.Bundle)
```

Результат будет выглядет так:

```bash
$ ruleguard -rules rules.go target.go
target.go:10:6: corerules/badCond: statement always true (diag.go:10)
target.go:11:6: corerules/badCond: statement always true (diag.go:11)
target.go:13:2: corerules/sortFuncs: suggestion: sort.Ints(ints) (refactor.go:12)
```

Если вам нужны не все диагностики, а только часть, стоит обратить внимание на параметры `-enable` и `-disable`, но о
них подробнее мы поговорим позднее.

### Создание собственного бандла

Чтобы ваш пакет правил можно было легко импортировать, нужно сделать следующее:

* Оформить пакет в качестве модуля, если это ещё не было сделано
* Добавить в один из файлов пакета переменную `Bundle`

Допустим, наш файл может стать бандлом, если мы перепишем его следующим образом:

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

var Bundle = dsl.Bundle{}

func suggestConcat(m dsl.Matcher) {
	m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
		Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
		Report("could use a simple concat for $x and $y instead of the Sprintf call")

	m.Match(`strings.Join([]string{$a, $b}, "")`).
		Report("could rewrite a strings.Join call as $a+$b")

	m.Match(`strings.Join([]string{$a, $b}, $glue)`).
		Report(`could rewrite a strings.Join call as $a+$glue+$b`)
}
```
