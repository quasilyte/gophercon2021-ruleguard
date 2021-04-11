## Настройка включенных правил

В случае, если используются только свои правила, самым интуитивным способом отключения правила может являться унесение правила
под комментарий (или его временное удаление).

Однако если вы используете внешние правила, сделать это будет невозможно.

Выборочное выключение правил происходит с помощью флага `-disable`, который принимает список названий групп правил, которые нужно выключить.

```bash
$ ruleguard -rules rules.go -disable 'dupArg,corerules/badCond' target.go
```

* TODO: Как отключать через `gocritic` и `golangci-lint`
* TODO: `-enable` флаг и для чего он может быть полезен