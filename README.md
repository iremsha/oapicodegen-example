# oapicodegen-example

## Начало работы

Указать в Preferences - Go - Go Modules - Environments:

`GOPRIVATE=*.something.org`

Или через консоль:

```
go env -w GOPRIVATE='*.something.org'
git config --global url."ssh://git@gitlab.something.org/".insteadOf "https://gitlab.something.org/"
```

Для работы с переменными окружения создаем файл из шаблона:

```
cp example.config.yaml config.yaml
```

Для использования линтера:

```
brew install golangci-lint
make lint
```

Локальный запуск:
```
make run - поднпимает docker
make go - собирает и запускает приложение
```

Для работы с миграциями:

```
go install github.com/golang-migrate/migrate/v4
make migrate-create title=new_migration_name
```

Для работы с Docker:
```
make run
make clean - останавливает контейнеры и удаляет все связанное с docker
```

Для работы с приложением:
```
make go
make test
```

Устанавливаем gomock:
```
go get -u github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

Генерируем mock:
```
mockgen -source=internal/handler/bank.go -destination=internal/handler/mock_bank.go -package=handler
```