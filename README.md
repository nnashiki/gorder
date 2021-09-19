# gorder
Go で order system  を作る


参考: https://selegee.com/7230/

# run

```shell
$ go run server.go

$ curl -X POST http://localhost:1323/api/user -d 'name=nashiki' -d 'email=nashiki@example.com'      
$ curl http://localhost:1323/api/user/1
$ curl "http://localhost:1323/api/user?name=nashiki"
$ curl -X PATCH http://localhost:1323/api/user/1 -d "email=nashiki2@example.com"


$ curl -X POST http://localhost:1323/api/order -d 'creator_id=1' -d 'contractor_id=2'
$ curl http://localhost:1323/api/order/1
$ curl -X PATCH http://localhost:1323/api/order/1 -d 'creator_id=3'
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

# TEST

参考:
https://qiita.com/nirasan/items/b357f0ad9172ab9fa19b#%E3%83%86%E3%82%B9%E3%83%88%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E3%83%95%E3%82%A3%E3%82%AF%E3%82%B9%E3%83%81%E3%83%A3%E3%83%BC%E3%82%92%E8%AA%AD%E3%81%BF%E8%BE%BC%E3%82%80
https://tech.raksul.com/2017/10/11/running-tdd-with-bdd-style-stories/
https://qiita.com/seya/items/582c2bdcca4ad50b03b7#%E8%BF%BD%E8%A8%98
https://qiita.com/theoden9014/items/ac8763381758148e8ce5

```shell
$ go test -v ./test
```
