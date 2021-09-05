# gorder
Go で order system  を作る


参考: https://selegee.com/7230/

# run

```shell
$ go run server.go

$ curl -X POST http://localhost:1323/api/user -d 'name=nashiki' -d 'email=nashiki@example.com'      
$ curl http://localhost:1323/api/user/1
$ curl -X PATCH http://localhost:1323/api/user/1 -d "email=nashiki2@example.com"


$ curl -X POST http://localhost:1323/api/order  -d "user_id=1"
$ curl http://localhost:1323/api/order/2
```

# ER図の作成
docker run --rm --network=host -v $PWD:/work k1low/tbls doc mysql://myuser:password@127.0.0.1:3306/mydb

# tips
- gorm を利用した部分的なアップデート
  - https://gorm.io/docs/update.html#Updates-multiple-columns

# SQLboiler

```shell
$ sqlboiler mysql -c mysql/sqlboiler.toml -o models --no-tests
$ go run boiler.go
```