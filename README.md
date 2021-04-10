# [GopherCon Russia](https://www.gophercon-russia.ru/) 2021 -- [ruleguard](https://github.com/quasilyte/go-ruleguard) workshop

> В директории [notes/](notes) можно найти краткое текстовое содержание по многим пунктам из плана.
> Ссылки в виде 📖 ведут на записи, относящиеся теме.

## План

Кратко: учимся использовать `ruleguard` и интегрируем его в наш пайплайн.

### Предварительная подготовка участников

См. [notes/pre_requirements.md](notes/pre_requirements.md).

### Вступление

* Введение в проблематику, "какие задачи мы собрались решать?"
* Введение в стандартную архитектуру анализаторов для Go
* Отличия архитектуры ruleguard от других анализаторов
* Вводим и разбираем основные понятия

### Основная часть

* [📖](notes/gorules_file.md) Структура gorules файлов
* [📖](notes/run_rules.md) Запуск правил
* [📖](notes/writing_rules.md) Пишем свои правила
* [📖](notes/filter_debug.md) Знакомимся с отладкой фильтров
* [📖](notes/testing_rules.md) Запускаем тесты для правил
* Подключаем внешние правила, знакомимся с [ruleguard bundles](https://quasilyte.dev/blog/post/ruleguard-modules/)
* Обсуждаем, как лучше хранить правила в репозитории, добавлять ли их в `go.mod`
* Разбираем quickfixes, дорабатываем правила через `Suggest()`
* Пробуем более сложные фильтры: программируемые, со сворачиванием константных выражений и прочие

### Интеграции

* [📖](notes/gocritic_integration.md) Запуск ruleguard через [go-critic](https://github.com/go-critic/go-critic)
* [📖](notes/golangci_integration.md) Запуск ruleguard через [golangci-lint](https://github.com/golangci/golangci-lint)
* [📖](https://github.com/quasilyte/gophercon2021-ruleguard/blob/master/notes/golangci_integration.md#%D1%80%D0%B5%D1%88%D0%B5%D0%BD%D0%B8%D0%B5-%D0%BF%D1%80%D0%BE%D0%B1%D0%BB%D0%B5%D0%BC-%D0%B8%D0%BD%D1%82%D0%B5%D0%B3%D1%80%D0%B0%D1%86%D0%B8%D0%B8) golangci-lint: решаем типичные проблемы интеграции
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

Вдохновение и идеи для правил можно брать здесь: [go-ruleguard.github.io/random](https://go-ruleguard.github.io/random/)

## Справочные материалы

* [ruleguard by example](https://go-ruleguard.github.io/by-example/)
* [Документация по формату gorule файлов](https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md)
* [godoc референс на пакет DSL](https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl)
* [Введение в ruleguard](https://habr.com/ru/post/481696/)

## Контактная информация

* [Группа go-critic в телеграме](https://t.me/go_critic_ru) (там же обсуждаем ruleguard)
