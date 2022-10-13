<div align="center">

# AB test support service

# Test cases are stored in memory while the application is running. If the application is stopped, all tests are deleted

</div>

## Running the application

```
docker-compose up --build
```

### Running the application without Docker

```
go run .
```

### Running tests (app must be running)

```
go test .
```

## HTTP Methods

```
"GET" / — Checking the server connection

    example: 
        "GET" :8080/
```

```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8080/ping
```

```
"GET" /all — Get all tests

    example: 
        "GET" :8080/all
```

```
"POST" /create — Create a new test case. Need JSON body

    example: 
        "POST" :8080/create
```

```json
{
    "Key" : "button_color",
    "Options" : {
        "#FF0000" : 33.3,
        "#00FF00" : 33.3,
        "#0000FF" : 33.3
    }
}
```

```
"PATCH" /forUser — Checks if an option is selected for the user. If not, assigns the option
    options:
        hypothesis — Name of the test case
        user — Users Id

    example: 
        "DELETE" :8080/one?hypothesis=button_color,user=5456d87545xx0
```

```
"DELETE" /one — Delete one test case
    options:
        hypothesis — Name of the test case

    example: 
        "DELETE" :8080/one?hypothesis=button_color
```

```
"DELETE" / — Delete all test cases

    example: 
        "DELETE" :8080/
```
