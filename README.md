# Запуск
Запуск  всего проекта через make:

Поднять postgres в docker контейнере
```sh
$ make start-postgres
```
Установка зависимостей (goose) :
```sh
$ make install_deps
```
Накатка миграций :
```sh
$ make local-migration-up
```
Запуск сервера приложения
```sh
$ make start-service
```
Теперь сервис обрабатывает входящие запросы по адрессу `localhost:8080`

Из сообщений, которые сервис сможет обработать, только `/ping` и `/api/tenders/new` 

