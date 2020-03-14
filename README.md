# Group Preparation Service (GPS)

## API:
  ### Работа с университетами
  * **GET** /universities - получение списка всех доступных в базе университетов
  * **GET** /universities (attributes: **id** - id или название университета) - получение конкретного университета
  
  * **GET** /universities/pages (attributes: **page** - номер страницы) - получение **page**-ой страницы со списком университетов
  
  ### Работа с мероприятиями
  * **GET** /events (**subject_id** - id или название предмета) - получение списка всех мероприятий данного предмета
  * **GET** /subjects/**subject_id**/events/pages (attributes: **page** - номер страницы) - получение **page** страницы со списком университетов
  * **GET** /event (attribute: **id**) - поиск мероприятия по частичному или полному совпадению названия или по id
  
  * **POST** /subjects/**subject_id**/events (attributes: **subject_id** - id предмета, **name** - название мероприятия) - добавление нового мероприятия
  * **PUT** /events (attributes: **id** - id мероприятия, название которого нужно изменить, **name** - новое название для мероприятия) - изменение название мероприятия
  
  * **DELETE** /events (attribute: **id**) - удаление мероприятия по id
  
  * **POST** /events/**event_id**/messages (parameters: **event_id** - id мероприятия, attributes: **text** - текст сообщения, **attachments** - файлы-вложения) - отправка сообщения в мероприятие
  * **GET** /events/**event_id**/messages (parameters: **event_id** - id мероприятия, attributes: **id** - id или название мероприятия) - получение списка всех сообщений в мероприятии
  
### Работа с комнатами
  * **GET** /room (attribute: **id** - id или частичное/полное название комнаты) - получение доступа к комнате
  * **POST** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **name** - полное название комнаты) - создание комнаты
  * **PUT** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **id** - id комнаты, **name** - новое название для комнаты) - изменение названия комнаты
  * **DELETE** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **id** - id комнаты) - удаление комнаты
  
  * **POST** /rooms/message (attributes: **text** - текст сообщения, **attachments** - файлы-вложения) - отправка сообщения в мероприятие
  * **GET** /rooms/messages (attributes: **id** - id или название мероприятия) - получение списка всех сообщений в мероприятии
  
### Работа с тестами
  * **GET** /events/**event_id**/tests (parameters: **event_id** - id мероприятия) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста
  
### Работа с вопросами
  * **GET** /tests/**test_id**/questions (parameters: **event_id** - id мероприятия) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста
  
 ### Работа с ответами
  * **GET** /events/**event_id**/tests (parameters: **event_id** - id мероприятия) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста

### Работа с авторизацией
  * /login (attributes: **email**, **password**) - вход 
  * /logout - выход
  * /signup (attributes: **email**, **password**, **university**) - регистрация
 
