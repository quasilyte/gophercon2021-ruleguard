## Пишем свои правила

Всё начинается с идеи.

Давайте попробуем находить некоторые фрагменты кода, в которых можно предлагать использовать конкатенацию строк вместо более сложных конструкций.

Первым примером такой операции может быть вызов `fmt.Sprintf` с аргументами-строками:

```diff
- fmt.Sprintf("%s%s", prefix, text)
+ prefix + text
```

Назовём эту диагностику `suggestConcat` (предложение использования конкатенации строк).

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestConcat(m dsl.Matcher) {
  // TODO
}
```

Любое правило стоит создавать с `Match` вызова. В его аргументах мы передаём что, собственно, должно попадать под рассмотрение.
В этих аргументах мы пишем примеры именно плохого кода.

Поскольку мы хотим матчить вызовы `fmt.Sprintf` с конкретной формат-строкой, мы так и записываем:

```go
m.Match(`fmt.Sprintf("%s%s", $x, $y))
```

Поскольку аргументы для `Sprintf` могут быть любые, мы вставляем туда переменные шаблона - `$x` и `$y`.

Дополнительно мы не хотим предлагать `$x + $y` если типы захваченных значений не являются строками. Это бы просто приводило к неправильным
предупреждениям (например, иногда нужно вызвать метод `String()`, формат `%s` делает это за нас).

На этом этапе мы добавляем вызов `Where` как раз с таким условием:

```go
m.Match(`fmt.Sprintf("%s%s", $x, $y)).
  Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`))
```

`m["x"]` адресует переменную шаблона `$x`; а `m["y"]` - переменную `$y`. Название могут быть любые, кроме `_` и `$$` (у этих названий есть особое назначение).

Предикат `Is(type-pattern)` проверяет совпадение типа выражения с указанным паттерном. В нашем случае паттерн максимально прост - обычное название типа, `string`.

Операнды фильтра можно соединять через `&&` и `||`, работать будет ровно так, как вы и ожидаете. Группировать через скобочки - можно.
Также работают отрицания через `!`, а для некоторых полей будут работать операции сравнения (`>`, `<`, `==`, `!=` и так далее).

После фильтрации мы хотим выполнить какое-то действие. Например, выдать предупреждение. Добавим вызов `Report`:

```go
m.Match(`fmt.Sprintf("%s%s", $x, $y)).
  Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
  Report("could use a simple concat for $x and $y instead of the Sprintf call")
```

Внутри строки-аргумента `Report` можно использовать переменные из шаблона `Match`. В них будут подставляться захваченные AST элементы.
Мы можем называть это интерполяцией переменных шаблона.

Например, если правило сработало на код `fmt.Sprintf("%s%s", foo.bar, y[0])`, то будет выдано такое предупреждение:

```
could use a simple concat for foo.bar and y[0] instead of the Sprintf call
```

Теперь наш файл с правилами выглядит как-то так:

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestConcat(m dsl.Matcher) {
  m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
    Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
    Report("could use a simple concat for $x and $y instead of the Sprintf call")
}
```

Это неплохое начало. Но мы можем идти дальше. Например, вот такой код тоже можно заменить на конкатенацию: `strings.Join([]string{s1, s2}, "")`.
Вы не поверите, но такой код иногда встречается.

Добавим для этого ещё одно правило в группу `suggestConcat`:

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func suggestConcat(m dsl.Matcher) {
  m.Match(`fmt.Sprintf("%s%s", $x, $y)`).
    Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
    Report("could use a simple concat for $x and $y instead of the Sprintf call")

  m.Match(`strings.Join([]string{$a, $b}, "")`).
    Report("could rewrite a strings.Join call as $a+$b")
}
```

Здесь нам не нужно проверять типы `$a` и `$b` так как иначе Go код бы просто не скомпилировался.

Можно продолжить:

```go
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

Всё это делается параллельно с тестированием и отладкой фильтров.
