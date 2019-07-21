### Go Concurrency Sample

This sample code explain that how to use go concurrency feature with mysql

#### Important Notice

Please make the necessary changes to the db/db.go file before running the application

db/db.go

``` 
line: 10 db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1)/dbname")
```

#### How to use it
``` 
go run main.go 
```

