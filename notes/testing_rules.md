## Тестируем правила

Мы написали следующий файл с правилами:

```
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

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

Как теперь это покрыть тестами?

Для начала, создадим рядом с файлом правил директорию `testdata`, а в ней создадим пустой файл `target.go`:

```
* rules.go
  * testdata/
    *target.go
```

В `target.go` добавим знакомый для нас пример:

```go
package example

import "strings"

func f(i int, b []byte, s1, s2 string) {
  _ = fmt.Sprintf("%s%s", s1, s2)
  _ = fmt.Sprintf("%s%s", i, s2)
  _ = fmt.Sprintf("%s%s", s1, b)
  _ = fmt.Sprintf("%s%s%s", s1, s2, "")
}
```

Теперь рядом с `rules.go` создадим файл `rules_test.go` и заполним его следующим содержимым:

```go
package gorules_test

import (
	"testing"

	"github.com/quasilyte/go-ruleguard/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRules(t *testing.T) {
	testdata := analysistest.TestData()
	if err := analyzer.Analyzer.Flags.Set("rules", "rules.go"); err != nil {
		t.Fatalf("set rules flag: %v", err)
	}
	analysistest.Run(t, testdata, analyzer.Analyzer, "./...")
}
```

Запускаются тесты привычным образом:

```bash
$ go test .
--- FAIL: TestRules (0.87s)
```

Ага, тесты проваливаются. А почему? Потому что линтер выдаёт предупреждение, а мы этого не ожидаем.

Нужно отредактировать наш `target.go` файл так, чтобы строка кода, на которую будет ругаться `ruleguard`, была размечена магическим комментарием:

```diff
- _ = fmt.Sprintf("%s%s", s1, s2)
+ _ = fmt.Sprintf("%s%s", s1, s2) // want `\Qcould use a simple concat for s1 and s2 instead of the Sprintf call`
```

> `// want` ожидает регулярку; мне нравится всегда добавлять `\Q` в начало, чтобы не думать о необходимости экранировать
> специальные символы внутри текста предупреждения. Особенно это удобно, когда в тексте встречаются скобочки.

Теперь тесты должны проходить успешно:

```
$ go test .
ok  	testrules	0.690s
```

Файл `rules_test.go` будет идентичен почти всегда, разве что в случае наличия нескольких файлов с правилами нужно будет перечислить их через запятую.

В `testdata` можно хранить несколько файлов и/или пакетов.
