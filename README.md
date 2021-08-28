# gorder
Go で order system  を作る


参考: https://selegee.com/7230/

# run

```shell
$ go run server.go
$ curl -X POST http://localhost:1323/api/product  -d "name=hoge&price=100"
$ curl http://localhost:1323/api/product/1

$ curl -X POST http://localhost:1323/api/user  -d "name=nashiki"      
$ curl http://localhost:1323/api/user/1
```
