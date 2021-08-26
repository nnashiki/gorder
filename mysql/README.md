# docker-mysql8

init db

```
docker run --rm --name mysql8 \
    -d \
    --rm \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_USER=myuser \
    -e MYSQL_PASSWORD=password \
    -e MYSQL_DATABASE=mydb \
    -v `pwd`/mysql/db_init:/docker-entrypoint-initdb.d \
    -v `pwd`/mysql/cnf:/etc/mysql/conf.d \
    -p 3306:3306 mysql:8.0 
```

use db

```
mysql -u myuser -p -h 127.0.0.1 mydb
```
