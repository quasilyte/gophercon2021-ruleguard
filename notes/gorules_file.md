## Структура gorules файла

Файл целиком называется `gorules` файл. Мы могли бы называть его просто "rules",
но тогда это будет конфликтовать с любой другой системой правил, которую вы можете использовать в вашем проекте.

Каждый файл правил импортирует пакет `dsl`.

![](/notes/images/gorules1.png)

На верхнем уровне находятся привычные для Go вещи: импорты, функции и, если необходимо, константы.

Если функция принимает тип `dsl.Matcher`, то она считается **группой правил**.

![](/notes/images/gorules2.png)

Каждая группа правил определяет некоторое количество **правил**.

Формально, можно создать группу с пустым множеством правил, но смысла в этом мало.

Правила внутри группы применяются в порядке определения, от первого до последнего.
Первое **сработавшее** правило останавливает обработку группу правил.
Таким образом можно размещать более конкретные правила перед более общими,
реализуя метод разрешения конфликтов.

![](/notes/images/gorules3.png)

Каждое правило содержит образец для сопоставления внутри `Match` выражения.
Чаще всего, это AST шаблон. Однако с недавних пор через `MatchComment` стало возможно
захватывать комментарии через регулярные выражения.

![](/notes/images/gorules4.png)

Каждый успешный матч проходит фильтрацию через выражения, указанные в `Where`.
Фильтры - опциональны, но они часто помогают уменьшить количество ложных срабатываний.

`Match` указывает форму, а `Where` проверяет содержимое.

Правило успешно срабатывает если входной элемент удовлетворяет и `Match`, и `Where`.

![](/notes/images/gorules5.png)

Финальным и обязательным элементом правил является действие, которое происходит при
успешном срабатывании.

Это может быть `Report` или `Suggest` (либо оба).

* `Report` - просто выводит предупреждение.
* `Suggest` - предлагает паттерн для автозамены проблемы (и тоже выдаёт предупреждение).

`Report` может быть использован вместе с `Suggest` для переопределения автоматически сгенерированного
текста, создаваемого `Suggest` шаблоном.

![](/notes/images/gorules6.png)
