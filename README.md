# GopherCon Russia 2021 -- ruleguard workshop

## План

**Предварительная подготовка участников:**

* В системе установлен Go 1.15 или выше
* Установлена последняя версия ruleguard
* Есть понимание того, что такое [Go modules](https://github.com/golang/go/wiki/Modules); умеем делать `go mod init` и установку модулей

Если у вас Windows или MacOS, стоит тщательно проверить, что всё запускается и работает.
Для этого можно использовать скрипт `check_setup.go` из корня этого репозитория:

```bash
$ git clone https://github.com/quasilyte/gophercon2021-ruleguard.git
$ cd gophercon2021-ruleguard
$ go run check_setup.go
```

**Вступление:**

* Введение в проблематику, "какие задачи мы собрались решать?"
* Введение в стандартную архитектуру анализаторов для Go
* Отличия архитектуры ruleguard от других анализаторов
* Вводим и разбираем основные понятия

**Основная часть:**

* Запуск правил, gorules файлы
* Пишем свои правила
* Знакомимся с отладкой фильтров
* Запускаем тесты для правил
* Подключаем внешние правила, знакомимся с ruleguard bundles
* Как лучше хранить правила в репозитории, добавлять ли их в `go.mod`
* Разбираем quickfixes, дорабатываем правила через `Suggest()`
* Пробуем более сложные фильтры: программируемые, со сворачиванием константных выражений и прочие

**Интеграции:**

* Запуск ruleguard через [go-critic](https://github.com/go-critic/go-critic)
* Запуск ruleguard через [golangci-lint](https://github.com/golangci/golangci-lint)
* golangci-lint: решаем типичные проблемы интеграции
* Разные наборы правил для разных частей проекта

**Дополнительная часть:**

* Как ruleguard импортирует пакеты
* Как упрощать идеи диагностик, чтобы они становились реальными для реализации
* Процесс добавления новых возможностей в ruleguard
* [Ограничения analysistest](https://github.com/golang/go/issues/37054)

**Практическая часть:**

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
