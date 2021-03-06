## Терминология

**Диагностика**, **проверка**, **чекер** - в нашем случае используем как синонимы. Но не включаем в список **правило**,
потому что в `ruleguard` каждая диагностика (проверка/чекер) может состоять из нескольких правил.

**Матчинг** - процесс сопоставления правила некоторому фрагменту кода. Сюда же входит применение фильтров.
Если фрагмент кода успешно подходит по AST шаблонам, не не по фильтрам, то успешного матчинга не происходит.

**Срабатывание правила** - то, что происходит при успешном матчинге.

**Паттерн**, **шаблон** - чаще всего имеется в виду AST шаблон, совместимый с [gogrep](https://github.com/mvdan/gogrep).
Это так же могут быть шаблоны типов, типа `[]$t` или `map[$k]$v` внутри фильтров `Where()`.
