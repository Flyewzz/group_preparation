# Group Preparation Service (GPS)

## API:
  ### Университеты
  * **GET** /universities (attributes: **name** - название университета, частично или полностью, **page** - номер страницы) - получение списка всех доступных в базе университетов или поиск по атрибутам, если **name** не пуст. Если **page** отсутствует или имеет нулевое значение, то выдается список всех университетов
  * **GET** /university (attributes: **id** - id университета) - получение конкретного университета по id
  * **POST** /university (attributes: **name** - краткое название университета, **full_name** - полное название университета) - создание университета
  * **DELETE** /university (attributes: **id** - id университета) - удаление конкретного университета
  * **DELETE** /universities - удаление всех университетов
  
### Предметы
  * **GET** /university/**id**/subjects (parameters: **id** - id университета, attributes: **name** - название предмета, частично или полностью, **semester** - семестр предмета, **page** - номер страницы) - получение списка всех предметов определенного университета или поиск по названию и/или семестру, если если **name** или **semester** не пуст. Если **page** отсутствует или имеет нулевое значение, то выдается список всех предметов определенного университета
  * **POST** /university/**id**/subject (parameters: **id** - id университета, attributes: **name** - название, **semester** - семестр) - добавление нового предмета
  * **GET** /subject (attributes: **id** - id предмета) - получение конкретного предмета по id
  * **DELETE** /subject (attributes: **id** - id предмета) - удаление предмета
  * **DELETE** /university/**id**/subjects (parameters: **id** - id университета) - удаление всех предметов определенного университета
  
### Материалы
  * **GET** /subject/**id**/materials (parameters: **id** - id предмета, attributes: **name** - название материала, частично или полностью, **type_id** - тип работы, **page** - номер страницы) - получение списка всех материалов определенного предмета или поиск по названию и/или типу работу, если если **name** или **type_id** не пуст. Если **page** отсутствует или имеет нулевое значение, то выдается список всех материалов определенного предмета
  * **POST** /subject/**id**/material (parameters: **id** - id предмета, attributes: **name** - название материала, **type_id** - id типа работы, **attachments** - список файлов) - добавление нового материала и файлов этого материала (опционально)
  * **GET** /material (attributes: **id** - id материала) - получение конкретного материала по id
  * **DELETE** /material (attributes: **id** - id материала) - удаление предмета
  * **DELETE** /subjects/**id**/materials (parameters: **id** - id предмета) - удаление всех материалов определенного предмета
  
### Файлы материалов
  * **GET** /material/{id}/files (parameters: **id** - id материала) - получение списка всех файлов определенного материала
  * **GET** /material/file/downloading (attributes: **id** - id файла) - скачивание определенного файла по его id

# ⬇️ **Not realized** ⬇️
  
### Комнаты
  * **GET** /room (attribute: **id** - id или частичное/полное название комнаты) - получение доступа к комнате
  * **POST** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **name** - полное название комнаты) - создание комнаты
  * **PUT** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **id** - id комнаты, **name** - новое название для комнаты) - изменение названия комнаты
  * **DELETE** /events/**event_id**/rooms (parameters: **event_id** - id мероприятия, attribute: **id** - id комнаты) - удаление комнаты
  
  * **POST** /rooms/message (attributes: **text** - текст сообщения, **attachments** - файлы-вложения) - отправка сообщения в мероприятие
  * **GET** /rooms/messages (attributes: **id** - id или название мероприятия) - получение списка всех сообщений в мероприятии
  
### Тесты
  * **GET** /events/**event_id**/tests (parameters: **event_id** - id мероприятия) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста
  
### Вопросы
  * **GET** /tests/**test_id**/questions (parameters: **event_id** - id мероприятия) - получение всех вопросов данного теста
  * **GET** /question (attributes: **id** - id или название теста) - получение конкретного вопроса
  
  * **POST** /tests/**test_id**/question (parameters: **test_id** - id теста, attributes: **name** - название вопроса) - создание вопроса
  
  * **POST** /questions/**question_id**/right (parameters: **question_id** - id вопроса, attributes: **answers** - id правильных вариантов ответа) - установка правильных вариантов ответа на вопрос
  
 ### Ответы
  * **GET** /questions/**question_id**/answers (parameters: **question_id** - id вопроса) - получение всех тестов для данного мероприятия
  * **GET** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **id** - id или название теста) - получение конкретного теста
  * **POST** /events/**event_id**/test (parameters: **event_id** - id мероприятия, attributes: **name** - название теста) - создание теста

### Авторизация и регистрация
  * **POST** /signin (attributes: **email**, **password**) - вход 
  * **POST** /signup (attributes: **email**, **password**) - регистрация
