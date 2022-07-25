### Указать доступные порты черезе `.` в конфиг файле `.env`
```shell
 GRPC_PORT=44300.44400.44500
```
### Формат запроса без TLS
```shell
 grpcurl -plaintext -d '{"video":"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' localhost:44300 Balancer/GetUrl
```
### При использование консоли PowerShell, нужно экранировать `"`, используем `\ `, пример запроса:
```shell
 grpcurl -plaintext -d '{\"video\":\"http://s1.origin-cluster/video/123/xcg2djHckad.m3u8\"}' localhost:44300 Balancer/GetUrl
```
### Тестовый клинет находится в `client.go`, использует настройки `.env`
```shell
 go run client.go
```
