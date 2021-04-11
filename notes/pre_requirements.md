## Предварительная подготовка участников

* В системе установлен Go 1.15 или выше
* Установлена последняя версия ruleguard
* Есть понимание того, что такое [Go modules](https://github.com/golang/go/wiki/Modules); умеем делать `go mod init` и установку модулей

Если у вас Windows или MacOS, стоит тщательно проверить, что всё запускается и работает.
Для этого можно использовать скрипт `check_setup.go` из корня этого репозитория:

```bash
$ git clone https://github.com/quasilyte/gophercon2021-ruleguard.git
$ cd gophercon2021-ruleguard/check_install
$ go run main.go
* running check go version step
* running install ruleguard to this module step
* running run ruleguard step
* running install rules bundle step
* running run ruleguard with bundle rules step
* running run gorule tests step
Everything looks good!
```

Лучше всего для наших образовательных целей ставить `go-ruleguard` из исходников. Но если вы предпочитаете бинарные релизы, то для версии `v0.3.4` есть бинарники для следующих платформ:

* [linux/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-linux-amd64.zip)
* [linux/arm64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-linux-arm64.zip)
* [darwin/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-darwin-amd64.zip)
* [windows/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-windows-amd64.zip)

Также стоит установить пакет `github.com/quasilyte/go-ruleguard/dsl`. Он является отдельным модулем и нужен для компиляции правил.

Собрать `ruleguard` из исходников можно, например, так:

```
$ git clone https://github.com/quasilyte/go-ruleguard.git
$ cd go-ruleguard
$ go build -o bin/ruleguard ./cmd/ruleguard
$ ./bin/ruleguard --help
```

> Я не рекомендую делать `-o ruleguard` так как в репозитории уже есть директория с названием `ruleguard`; поэтому кладём бинарник в директорию `./bin`.

Исполняемый файл `bin/ruleguard` можно положить куда-нибудь под системный `$PATH`.

Для Windows может потребоваться делать `-o bin/ruleguard.exe` (с суффиксом `.exe`).
