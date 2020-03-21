# Group Preparation Service (GPS)

## API:
  ### Работа с университетами
  * **GET** /universities - получение списка всех доступных в базе университетов
  * **GET** /university (attributes: **id** - id или название университета) - получение конкретного университета
  
  * **POST** /university (attributes: **name** - название университета) - создание университета
  
  * **GET** /universities (attributes: **page** - номер страницы) - получение **page**-ой страницы со списком университетов
  
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
  * **GET** /tests/**test_id**/questions (parameters: **event_id** - id мероприятия) - получение всех вопросов данного теста
  * **GET** /question (attributes: **id** - id или название теста) - получение конкретного вопроса
  
  * **POST** /tests/**test_id**/question (parameters: **test_id** - id теста, attributes: **name** - название вопроса) - создание вопроса
  
  * **POST** /questions/**question_id**/right (parameters: **question_id** - id вопроса, attributes: **answers** - id правильных вариантов ответа) - установка правильных вариантов ответа на вопрос
  
 ### Работа с ответами
  * **GET** /questions/**question_id**/answers (parameters: **question_id** - id вопроса) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста

### Работа с авторизацией
  * /login (attributes: **email**, **password**) - вход 
  * /logout - выход
  * /signup (attributes: **email**, **password**, **university**) - регистрация
 
