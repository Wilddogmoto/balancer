Запуск проекта происxодит через `main.go` настройки проекта лежат в `.env`. Проект создает несколько gRPC серверов
для приема запросов формата `grpcurl -plaintext -d '{"increment": "22","video":"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' localhost:44300 Balancer/GetUrl` для балансеровки вызовов и перенаправления на CDN сервера. 
`"increment"` добавил для наглядности теста

`client.go` делает запросы на сервера по указанным портам и балансирует вызовы между ними, при этом общаяя логика отправки каждого 10го запроса по оригинальному адрессу сохраняется.
### Указать доступные порты черезе `.` в конфиг файле `.env`
```
 GRPC_PORT=44300.44400.44500
```
### Формат запроса без TLS
```
 grpcurl -plaintext -d '{"increment": "22","video":"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' localhost:44300 Balancer/GetUrl
```
### При использование консоли PowerShell, нужно экранировать `"`, используем `\ `, пример запроса:
```
 grpcurl -plaintext -d '{\"increment\": 1,\"video\":\"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8\"}' localhost:44300 Balancer/GetUrl
```
### Тестовый клиент находится в `client.go`, использует настройки `.env`. запускаем его после запуска `main.go`
```
 go run client.go
```

### Тестирование с помощью GHz запуск через `ghz.go`
```
 go run ghz.go
```
