# service


Общее описание:
Нужно сделать сервис для больницы, при помощи которого можно делать запись на прием к врачу.


Кейсы использования:
	Есть админ, со своей учетной записью. Он может добавлять/изменять/удалять в базе врачей. Есть пользователь, который может создать аккаунт, при помощи своих данных записаться на прием (при этом должна быть возможность получить все свободные слоты записи), вывести все свои записи на прием. У каждого врача есть 10 слотов записи в день (9:00-10:00, 10:00-11:00 и т. д.).


Эндпоинты:

POST /register
```	Body: 
{
 name: “name”,
	 surname: “surname”,
	 login: “login”,
	 password: “password”
}
	Регистрирует пользователя

POST /doctor
Body:
{
 name: “name”,
	 surname: “surname”
}
BodyReponse:
{
 id:”id”,
 name:”name”,
 surname:”surname”
}
Создает доктора
PUT /doctor
Body:
{
id: “id”
 name: “name”,
	 surname: “surname”
}
Меняет данные доктора по id
DELETE /doctor/{id}
	Удаляет доктора по id
Get /doctors/all
	{
 [doctorName: “doctorName”, id: “id”]
}
Получает список всех докторов
GET /getall
BodyResponse:
{
 [time:”time”, doctorName: “doctorName”, timeTableId: “timeTableId”]
}
Получает список всех свободных для записи 
POST /setme
Body:
{
 login: “login”,
 password: “password”,
 timeTableId: “timeTableId”
}
BodyResponse:
{
 login: “login”,
 timeTableId: “timeTableId”,
 time: “time”
}
Записывает пользователя с логином {login} на время с id = {timeTableId}
GET /getall/{login}
BodyResponse:
{
 [time:”time”, doctorName: “doctorName”, timeTableId: “timeTableId”]
}
	Получает записи клиента с логином {login}
```


Технологии:
	Язык сервиса - GoLang
	База даных - PostgreSQL
	Средство деплоя - docker и docker-compose

При проектировании БД сохранить скрипт создания таблиц и скинуть сюда. Запросы можно делать в DataGrip.
PostgreSQL поднять в докер контейнере
Запросы к сервису можно делать с помощью postman

##Дополнения:

Генерация и сохранение cookie при авторизации.

Реализован чат для авторизованных пользователей, его держит отдельный сервер, общение происходит с помощью сокетов. 
