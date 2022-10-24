# movie_app

The user will create an account by registering. The logged in user will be able to add, delete and edit movies. If it couldn't login, it will get Unauthorized error. Users who are not logged in will only be able to get a list of all movies.

The application will use HTTP status codes to communicate the success or failure of an operation.

## Getting Started

1. git clone https://github.com/denizcamalan/movie_app.git

2. cd  movie_app

3. Run application
```
docker-compese up -d
```
4.  Browse Swagger UI [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

5. Autherizion
```
barrer {mytoken}
``` 

6. Run Tests
```
go test ./...
```
