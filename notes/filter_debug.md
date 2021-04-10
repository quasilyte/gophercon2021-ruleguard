## Отладка фильтров

Возьмём диагностику `suggestConcat` из предыдущей записи:

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestConcat(m dsl.Matcher) {
  m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
    Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
    Report("could use a simple concat for $x and $y instead of the Sprintf call")
}
```

В ней присутствуют два выражения внутри `Where`. Поэтому есть как минимум 3 причины, по которой предупреждения не будет выдаваться:

1. Не сработал матчинг шаблона `fmt.Sprintf("%s%s", $x, $y)`
2. Не сработал фильтр по `m["x"]`
3. Не сработал фильтр по `m["y"]`

Чтобы упростить отладку, в `ruleguard` есть флаг `-debug-group`.

Возьмём следующий файл для анализа (`target.to`):

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

А теперь запустим `ruleguard` с флагом отладки:

```bash
$ ruleguard -rules rules.go -debug-group suggestConcat ./target.go
target.go:7: [rules.go:6] rejected by m["x"].Type.Is(`string`)
  $x int: i
  $y string: s2
target.go:8: [rules.go:6] rejected by m["y"].Type.Is(`string`)
  $x string: s1
  $y []byte: b
target.go:6:6: suggestConcat: could use a simple concat for s1 and s2 instead of the Sprintf call (rules.go:6)
```

Что нам это всё говорит?

* На строку 7 предупреждение не выдано потому что `$x` имеет тип `int`
* На строку 8 предупреждение не выдано потому что `$y` имеет тип `[]byte`
* Строка 9 проигнорирована, потому что не подошла под шаблон в `Match`
* Только на 6-ую строку было выдано предупреждение (пройдены все фильтры)

Заметим, что отладка запускается для определённой диагностика ("группы правил"), поэтому если в одном файле у вас много подобных групп правил,
вы всё равно сможете удобно отлаживать интересующую вас диагностику.
