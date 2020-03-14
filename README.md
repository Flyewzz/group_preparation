# Group Preparation Service (GPS)

## API:
  ### Работа с университетами
  * **GET** /universities => получение списка всех доступных в базе университетов
  * **GET** /universities (attributes: **id** - id или название университета) => получение конкретного университета
  
  * **GET** /universities/pages (attributes: **page** - номер страницы) => получение **page** страницы со списком университетов
  
  ### Работа с мероприятиями
  * **GET** /events (**subject_id** - id или название предмета) => получение списка всех мероприятий данного предмета
  * **GET** /events (attribute: **id**) => поиск мероприятия по частичному или полному совпадению названия или по id
  
  * **POST** /subjects/**subject_id**/events (attributes: **subject_id** - id предмета, **name** - название мероприятия) => добавление нового мероприятия
  * **PUT** /events (attributes: **id** - id мероприятия, название которого нужно изменить, **name** - новое название для мероприятия) - изменение название мероприятия
  
  * **DELETE** /events (attribute: **id**) => удаление мероприятия по id
  
  * **POST** /events
  
### Работа с комнатами
  * **GET** /events/**univ_id**/rooms (parameters: **univ_id** - id университета, attribute: **id** - id или частичное/полное название комнаты) - получение доступа к комнате
  * **POST** /events/**univ_id**/rooms (parameters: **univ_id** - id университета, attribute: **name** - полное название комнаты) - создание комнаты
  * **PUT** /events/**univ_id**/rooms (parameters: **univ_id** - id университета, attribute: **id** - id комнаты, **name** - новое название для комнаты) - изменение названия комнаты
  * **DELETE** /events/**univ_id**/rooms (parameters: **univ_id** - id университета, attribute: **id** - id комнаты) - удаление комнаты
