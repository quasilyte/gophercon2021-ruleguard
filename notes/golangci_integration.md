## Запуск ruleguard из golangci-lint

### Включаем ruleguard

[golangci-lint](https://github.com/golangci/golangci-lint) имеет в своём наборе линтер [gocritic](https://github.com/go-critic/go-critic). А `gocritic` интегрирует `ruleguard`. Следовательно, мы можем использовать `ruleguard` через `golangci-lint`.

![image](https://user-images.githubusercontent.com/6286655/114297491-9ab34800-9ab9-11eb-8845-2e7df8a43895.png)

Чтобы пользоваться `ruleguard`, нужно выполнить 3 шага:

* Убедиться, что `gocritic` включен (по умолчанию он выключен)
* `ruleguard` чекер из `gocritic` тоже включен
* Установлен параметр `ruleguard.rules`

Всё это выполняется редактированием вашего `.golangci.yml` конфига:

```yaml
### ...

linters:
  enable:
    - gocritic
linters-settings:
  gocritic:
    enabled-checks:
      - ruleguard
    settings:
      ruleguard:
        rules: "rules.go"
        
### ...
```

Если для вас привычнее JSON-нотация, можно записать этот конфиг в стиле [yaml5](https://github.com/quasilyte/yaml5):

```json
{
  "linters": {
    "enable": [
      "gocritic",
    ],
  },  
  "linters-settings": {
    "gocritic": {
      "enabled-checks": [
        "ruleguard",
      ],
      "settings": {
        "ruleguard": {
          "rules": "rules.go",
        },
      },
    },
  },
}
```

Как вы видите, мы ссылаемся на файл `rules.go`. Допустим, его содержимое выглядит следующим образом:

```go
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func boolExprSimplify(m dsl.Matcher) {
	m.Match(`!($x != $y)`).Suggest(`$x == $y`)
	m.Match(`!($x == $y)`).Suggest(`$x != $y`)
}
```

Тогда при запуске `golangci-lint` мы можем получать примерно такие сообщения:

```bash
$ golangci-lint run example.go 
example.go:5:9: ruleguard: can rewrite as xs[0] == ys[0] (gocritic)
        return !(xs[0] != ys[0])
               ^
```

Если вы активно редактируете файл с правилами `rules.go`, вам может потребоваться время от времени сбрасывать кэш `golangci-lint`.

```bash
golangci-lint cache clean
```

По этой и некоторым другим причинам я рекомендую тестировать и разрабатывать правила с использованием standalone бинарника `ruleguard`, а не
через интеграцию с `gocritic` или `golangci-lint`.

### Решение проблем интеграции

#### Относительный путь в параметре rules

При использовании относительного пути в конфиге `.golangci.yml` для параметра `rules` вы получите путь, относительный директории запуска `golangci-lint`.

Одним из решений является препроцессинг `.golangci.yml` файла. Например, [вот пример с использованием mustache](https://github.com/golangci/golangci-lint/issues/1662#issuecomment-778898705).

```yaml
    settings:
      ruleguard:
        rules: "{{RULEGUARD_DIR}}/myrules.go"
```
