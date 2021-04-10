## quickfixes

Помните наши замечательные правила для конкатенации строк?

Мы можем их улучшить, добавив автоматическое исправление.

Автоматические исправления могут работать как из консольного режима, так и при интеграции с LSP-совместимыми редакторами, типа VS Code.

Чтобы определить правило переписывания, мы вызываем метод `Suggest()` со строкой-шаблоном.
Эта строка может содержать переменный шаблона из `Match`, интерполяция будет аналогична той, что присутствует в `Report()` строках.

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestConcat(m dsl.Matcher) {
  m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
    Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
    Report("could use a simple concat for $x and $y instead of the Sprintf call").
    Suggest(`$x + $y`)

  m.Match(`strings.Join([]string{$a, $b}, "")`).
    Report("could rewrite a strings.Join call as $a+$b").
    Suggest(`$a + $b`)
    
  m.Match(`strings.Join([]string{$a, $b}, $glue)`).
    Report(`could rewrite a strings.Join call as $a+$glue+$b`).
    Suggest(`$a + $glue + $b`)
}
```

> Для простейших случаев можно использовать `Suggest()` без `Report()`. Для таких правил автоматически создаётся текст предупреждения.
> Но если `Report()` всё-таки используется, то сообщение будет браться именно оттуда.

Чтобы воспользоваться автоматическим исправлением, утилите `ruleguard` нужно передать флаг `-fix`:

```bash
$ ruleguard -rules rules.go -fix target.go
```

Все изменения происходят inplace в файлах, где были найдены проблемы.

Если для вашего правила сложно выразить корректный `Suggest()`, лучше ограничиться понятным сообщением в `Report()`.
Выдача некорректных quickfix'ов сломает код.

Исправления лучше всего добавлять для простых диагностик, где есть очевидный метод исправления,
а сама диагностика не выдаёт ложных срабатываний.
