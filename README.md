# Тестовое задание стажера в AvitoMX
JSON API сервис для хранения данных о товарах и продавцах, полученных из файла формата .xlsx.
## Запуск сервиса 
Клонировать проект:  
```git clone https://github.com/Dukastlik/avitomx-api```  
Перейти в директорию avitomx-api:  
```cd avitomx-api```  
Запустить проект с помощью `docker-compose up`

**Сервис доступен по http://localhost:8080**

## Интерфейс
1. Метод `add` передает в сервис идентификатор продавца и ссылку на файл с информацией о продажах, чтобы воспользоваться методом нужно:  
Отправить по адресу http://localhost:8080/add `POST` запрос с json вида:
```
{
   "MerchantId": <MerchntId>,
    "FileLink": <FileLink>
 }
 ```
Для тестирования можно воспользоваться curl:  
```curl -d @"test_add.json" -X POST http://localhost:8080/add```  
В ответ сервис пришлет json c информацией о добавленных данных:  
```
{
    "Products created": 5,
    "Products updated": 0,
    "Products deleted": 0,
    "Products invalid": 2
}
```  
 2. Метод `stat` достает из БД список товаров, соответствующих параметрам запроса, чтобы воспользоваться методом нужно:
Отправить на http://localhost:8080/stat `GET` запрос вида:
```
http://localhost:8080/stat?merchID=<merchID>&offerID=<oofferID>&prod_name=<prod_name>
```  

Для тестирования можно воспользоваться curl:  
```curl -d @"test_stat.json" -X GET http://localhost:8080/stat?merchID=23&offerID=11&prod_name=phone```  
В ответ сервис пришлет собранную статистику, например:  
```
[
    {
        "merchID": 23,
        "offerID": 11,
        "name": "bread",
        "price": 112,
        "quantity": 123,
        "available": true
    },
    {
        "merchID": 23,
        "offerID": 12,
        "name": "brew",
        "price": 344,
        "quantity": 10,
        "available": true
    }
]
```

