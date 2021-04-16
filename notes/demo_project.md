## Практика

Попробуем запустить проверки на демо-репозитории.

> **Важно**: в системе должен быть установлен `ruleguard`. Команда `ruleguard` должна быть доступна через ваш `$PATH`.

Все инструкции написаны для `bash`. Если у вас что-то другое, то вам придётся самостоятельно находить аналоги для указанных команд.

Для начала скачаем репозиторий:

```bash
$ git clone https://github.com/go-ruleguard/demo-project.git
$ cd demo-project
```

Если у вас в окружении установлен `$GOPATH`, мы можете захотеть отключить его на время экспериментов, чтобы не подтянуть пакет `dsl` из
своего `GOPATH` (если он у вас там уже есть).

```bash
# Уберём GOPATH для текущей bash сессии
unset GOPATH
```

Мы также хотим быть на 100% уверенными, что Go модули будут использоваться. Если у вас в окружении выставлена переменная `GO111MODULE`, то стоит
или убрать её, либо установить в `"on"`.

```bash
# Для всевозможных инсталляций форсированно включим режим модулей.
export GO111MODULE=on
```

Теперь мы готовы работать со скачанным проектом.

Выполним проверку проекта с помощью правил из этого же проекта:

```bash
$ ruleguard --c=0 --rules=rules/rules.go ./mandelbrot
mandelbrot/main.go:36:9: imagePt: zero point should be written as image.Point{} (rules.go:30)
36		min := image.Pt(0, 0)
mandelbrot/main.go:40:11: imageColors: suggestion: color.Black (rules.go:8)
40		black := color.Gray16{0}
mandelbrot/main.go:41:48: imageZP: image.ZP is deprecated, use image.Point{} instead (rules.go:24)
41		draw.Draw(b, bounds, image.NewUniform(black), image.ZP, draw.Src)
```

Если у вас установлен `golangci-lint`, то можно запустить и через него:

```basj
$ golangci-lint run ./mandelbrot
mandelbrot/main.go:36:9: ruleguard: zero point should be written as image.Point{} (gocritic)
	min := image.Pt(0, 0)
	       ^
mandelbrot/main.go:40:11: ruleguard: suggestion: color.Black (gocritic)
	black := color.Gray16{0}
	         ^
mandelbrot/main.go:41:48: ruleguard: image.ZP is deprecated, use image.Point{} instead (gocritic)
	draw.Draw(b, bounds, image.NewUniform(black), image.ZP, draw.Src)
	                                              ^
```

А вот запуск через `gocritic` (если он у вас установлен):

```bash
$ gocritic check -enable ruleguard -@ruleguard.rules rules/rules.go ./mandelbrot
./mandelbrot/main.go:41:48: ruleguard: image.ZP is deprecated, use image.Point{} instead
./mandelbrot/main.go:40:11: ruleguard: suggestion: color.Black
./mandelbrot/main.go:36:9: ruleguard: zero point should be written as image.Point{}
```

Для правил, которые лежат в `demo-project/rules` написаны тесты.

Запустить их можно следующим образом:

```bash
$ go test -v ./rules
=== RUN   TestRules
--- PASS: TestRules (1.00s)
PASS
```

Что можно попробовать:

* Добавить новые тестовые кейсы в `rules/testdata/`
* Добавить новые правила или доработать имеющиеся
* Запустить `ruleguard` с параметром `--fix` и применить автофиксы
* Вспомнить, как выглядит визуализация множества Мандельброта

<img src="https://user-images.githubusercontent.com/6286655/114301199-7b71e600-9acc-11eb-9815-114bab1dd99e.png" width="25%" height="25%">
