# [GopherCon Russia](https://www.gophercon-russia.ru/) 2021 -- [ruleguard](https://github.com/quasilyte/go-ruleguard) workshop

> В директории [notes/](notes) можно найти краткое текстовое содержание по многим пунктам из плана.

## План

### Предварительная подготовка участников

* В системе установлен Go 1.15 или выше
* Установлена последняя версия ruleguard
* Есть понимание того, что такое [Go modules](https://github.com/golang/go/wiki/Modules); умеем делать `go mod init` и установку модулей

Если у вас Windows или MacOS, стоит тщательно проверить, что всё запускается и работает.
Для этого можно использовать скрипт `check_setup.go` из корня этого репозитория:

```bash
$ git clone https://github.com/quasilyte/gophercon2021-ruleguard.git
$ cd gophercon2021-ruleguard/check_install
$ go run main.go
```

Лучше всего для наших образовательных целей ставить `go-ruleguard` из исходников. Но если вы предпочитаете бинарные релизы, то для версии `v0.3.4` есть бинарники для следующих платформ:

* [linux/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-linux-amd64.zip)
* [linux/arm64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-linux-arm64.zip)
* [darwin/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-darwin-amd64.zip)
* [windows/amd64](https://github.com/quasilyte/go-ruleguard/releases/download/v0.3.4/ruleguard-windows-amd64.zip)

Также стоит установить пакет `github.com/quasilyte/go-ruleguard/dsl`. Он является отдельным модулем и нужен для компиляции правил.

### Вступление

* Введение в проблематику, "какие задачи мы собрались решать?"
* Введение в стандартную архитектуру анализаторов для Go
* Отличия архитектуры ruleguard от других анализаторов
* Вводим и разбираем основные понятия

### Основная часть

* Запуск правил, gorules файлы
* Пишем свои правила
* Знакомимся с отладкой фильтров
* Запускаем тесты для правил
* Подключаем внешние правила, знакомимся с [ruleguard bundles](https://quasilyte.dev/blog/post/ruleguard-modules/)
* Как лучше хранить правила в репозитории, добавлять ли их в `go.mod`
* Разбираем quickfixes, дорабатываем правила через `Suggest()`
* Пробуем более сложные фильтры: программируемые, со сворачиванием константных выражений и прочие

### Интеграции

* Запуск ruleguard через [go-critic](https://github.com/go-critic/go-critic)
* Запуск ruleguard через [golangci-lint](https://github.com/golangci/golangci-lint)
* golangci-lint: решаем типичные проблемы интеграции
* Разные наборы правил для разных частей проекта

### Дополнительная часть

* Как ruleguard импортирует пакеты
* Как упрощать идеи диагностик, чтобы они становились реальными для реализации
* Процесс добавления новых возможностей в ruleguard
* [Ограничения analysistest](https://github.com/golang/go/issues/37054)

### Практическая часть

* Находим проект для проверки (желательно, чтобы там уже был golagnci-lint конфиг)
* Включаем ruleguard, пишем правила
* Запускаем правила на выбранном проекте
* Пишем тесты для правил

## Справочные материалы

* [ruleguard by example](https://go-ruleguard.github.io/by-example/)
* [Документация по формату gorule файлов](https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md)
* [godoc референс на пакет DSL](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl)
* [Введение в ruleguard](https://habr.com/ru/post/481696/)

## Контактная информация

* [Группа go-critic в телеграме](https://t.me/go_critic_ru) (там же обсуждаем ruleguard)
