# Go-TodoList

Implementation of a simple todolist webapp using Go programming language using gorilla/mux and go-sql-driver libararies

Clone the project. It uses Go Modules for dependency management, so you'll only will have to change for proper module
path. Everything else should be resolved automatically. If not, do it yourself)

Install gorilla/mux

```
go get -u github.com/gorilla/mux
```

Install go-sql-driver

```
go get -u github.com/go-sql-driver/mysql
```

Modify constants in main.go file for correct mysql configuration and run main.go

```
go run main.go
```

Or use this command to stop local mysql service

```bigquery
sudo systemctl stop mysql
```

Then run it in Docker using

```bigquery
docker-compose up
```

And then run at the root of the project directory

```bigquery
go run .
```