# REST на Golang
 
REST-приложение на Golang. Реализованы основные HTTP-запросы (GET(id), GET(all), UPDATE(POST) и Add(PUT)-запросы) позволяющие произвести почти CRUD-операции

Данное ПО написано на:

1. Golang 1.19 (backend)
2. Используется СУБД Postgres

Данное ПО распространяется по лицензии Линукс и с принципом AS IS - т.е. без гарантий и ответственности...

Если используете Postman, то в закладке Headers необходимо добавить {'Accept': 'application/json'}, а в Authorization {'user': 'password'}

Примеры:
GET
localhost:8080/get

GET(all)
localhost:8080/all

UPDATE(POST)
localhost:8080/update?id=1&phone=891012345678

Add(PUT)
localhost:8080/add?user_name=Kirill&phone=891112345678
