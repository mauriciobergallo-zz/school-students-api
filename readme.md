# Test the Application
```
go test
```

# Test the Application with Coverage Report in HTML

```
go test -coverprofile cover.out
go tool cover -html cover.out
```

# Endpoints

## Listing the Students

```
GET: http://localhost:8080/api/students
Payload:
{
    "message": [
        {
            "id": "e638ef30-5304-4369-9c11-3486ba6daabc",
            "firstName": "John",
            "lastName": "Doe",
            "identificationNumber": "12345678",
            "email": "john.doe@mail.com",
            "birthDate": "1984-05-22T00:00:00Z"
        },
        {
            "id": "9b8e22e1-94ac-4b54-bb86-2b8260c15c77",
            "firstName": "Jane",
            "lastName": "Doe",
            "identificationNumber": "45678912",
            "email": "jane.doe@gmail.com",
            "birthDate": "1984-04-07T00:00:00Z"
        }
    ]
}
```

## Creating new Students
```
POST: http://localhost:8080/api/students
Payload:
{
    "firstName": "Steve",
    "lastName": "Perez",
    "identificationNumber": "13130994",
    "email": "steve.perez@mail.com",
    "birthDate": "1990-10-11T00:00:00Z"
}
```