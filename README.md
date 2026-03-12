# loglint

**loglint** — статический анализатор для Go, который проверяет лог-записи на соответствие стандартам качества и безопасности. Линтер интегрируется в `golangci-lint` и поддерживает популярные логгеры: `log/slog` и `go.uber.org/zap`.

## репозитории

- **loglint**: https://github.com/aidashovv/loglint — исходный код анализатора и правила проверки
- **golangci-lint (форк)**: https://github.com/aidashovv/golangci-lint — форк с интегрированным линтером для использования

## структура проекта

```
loglint/
├── pkg/
│   └── analyzer/
│       ├── analyzer.go       # основной анализатор
│       ├── checker.go        # реализация правил
│       ├── runner.go         # обход AST
│       └── analyzer_test.go  # тесты
├── testdata/                 # тестовые файлы
└── README.md
```
