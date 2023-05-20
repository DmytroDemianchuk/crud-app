# CRUD Application for Managing Music Lists
## I used the following concepts during development:
- Building a Web Application with Go, following the REST API design.
- The Clean Architecture approach in building the structure of an application. Dependency injection technique.
- Working with the <a href="https://github.com/gorilla/mux">gorilla/mux</a>.framework.
- Working with Postgres database. Run from Docker. Generation of migration files.
- Working with the database using the sqlx library.
- Writing SQL queries.

### To start application

Download github repository in your desktop

```
https://github.com/DmytroDemianchuk/crud-app.git
```

Download docker image from the internet

```
docker pull postgres
```

Create docker container

```
docker run --name=crud-app -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```

Definition migrating to database

```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```

Start application

```
source .env && go build -o app cmd/main.go && ./app
```

### Example of creating a music
```
{
    "name": "Good Morning",  
    "artist": "Kanye West",  
    "album": "Graduation",  
    "genre": "Hip hop",
    "released_year": 2007
}
```