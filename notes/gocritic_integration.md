## Запуск ruleguard из линтера gocritic

[go-critic](https://github.com/go-critic/go-critic) имеет диагностику под названием [ruleguard](https://go-critic.github.io/overview.html#ruleguard).

Эта диагностика реализована с помощью [библиотечного API ruleguard](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/ruleguard).

```bash
$ gocritic doc | grep ruleguard
ruleguard [style experimental]

$ gocritic doc ruleguard
ruleguard checker documentation
URL: https://github.com/go-critic/go-critic/checkers
Tags: [style experimental]

Runs user-defined rules using ruleguard linter.

Reads a rules file and turns them into go-critic checkers.

See https://github.com/quasilyte/go-ruleguard.

Checker parameters:
  -@ruleguard.debug string
    	enable debug for the specified named rules group (default "")
  -@ruleguard.failOnError bool
    	If true, panic when the gorule files contain a syntax error. If false, log and skip rules that contain an error (default false)
  -@ruleguard.rules string
    	comma-separated list of gorule file paths. Glob patterns such as 'rules-*.go' may be specified (default "")
```

Основным параметром является `-@ruleguard.rules`, он указывает, из каких файлов нужно читать определения правил.

Следующим образом можно запустить `gocritic` со включенным `ruleguard`:

```bash
gocritic check -enable ruleguard -@ruleguard.rules rules.go ./target
```

Параметр `-@ruleguard.debug` эквивалентен параметру `-debug-group` в `ruleguard`. Отладка с его помощью возможна, но я рекомендую тестировать и отлаживать правила через отдельный `ruleguard`, а `gocritic` (или `golangci-lint`) использовать только для финальной интеграции.

Чаще всего, если возникает какая-то проблема с работой `ruleguard` из `gocritic`, корень лежит где-то [здесь](https://github.com/go-critic/go-critic/blob/master/checkers/ruleguard_checker.go).
