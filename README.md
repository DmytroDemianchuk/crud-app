# CRUD Application for Managing Music Lists
## I used the following concepts during development:
- Building a Web Application with Go, following the REST API design.
- The Clean Architecture approach in building the structure of an application. Dependency injection technique.
- Working with the <a href="https://github.com/gorilla/mux">gorilla/mux</a>.framework.
- Working with Postgres database. Run from Docker. Generation of migration files.
- Working with the database using the sqlx library.
- Writing SQL queries.

## Requirements
- go 1.20
- docker & docker-compose

## Run Project

Definition migrating to database

```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

Use `make run` to build&run project


## API:
### POST /auth/sign-up

Creates new user 

##### Example Input: 
```
{
	"name": "user",
	"email": "user@user.com",
    "password": "password"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input: 
```
{
	"email": "user@user.com",
    "password": "password"
} 
```

##### Example Response: 
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /musics

Creates new musics

##### Example Input: 
```
{
    "name": "Good Morning",  
    "artist": "Kanye West",  
    "album": "Graduation",  
    "genre": "Hip hop",
    "released_year": 2007
}
```

### GET /musics

Returns all user bookmarks

##### Example Response: 
```
{
    "name": "Good Morning",  
    "artist": "Kanye West",  
    "album": "Graduation",  
    "genre": "Hip hop",
    "released_year": 2007
} 
```


### DELETE /musics

Deletes bookmark by ID:

##### Example Input: 
```
{
	"id": "1"
} 
```