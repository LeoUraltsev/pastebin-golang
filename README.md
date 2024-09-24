# PASTEBIN Golang

Клон Pastebin на golang

---

### TODO:
* running in Docker
* swagger dock
* Обработать возвращение ошибок при запросах (сейчас считаю что всё отрабатывает всегда ОК)

### Запуск приложения

Создать файл `.env` и заполнить поля из файла `.env.example` 
```Shell
go run cmd\pastebin\main.go
```


### Endpoint's
**Main api v1**
```
host:port/api/v1
```

**Paste**
GET:

/ - получение последних x публичных

/{urlpaste} - получение по ссылке
Response
```JSON
{
    "status": "OK", // 200 - успешно создано
    "paste": 
    {
        "title": "Заголовок",
        "data": "Текстовая информация"
    }
}
```


POST:

**/paste** - создание
Request
```JSON
{
    "url": "myurl", //необязательное поле, если занято вернёт ошибку. Если не указано будет сгенерировано автоматически (hash)
    "title": "Hello World", // Заголовок
    "data": "Hello pastebin golang", // Набор текстовой информации
    "status": "public", //public - доступно всем, unlisted - доступ только по ссылке
    "expiration_time": 60 // время в минутах
}
```
Response
```JSON
{
    "status": "OK", // 201 - успешно создано
    "url": "address/myurl"
}
```