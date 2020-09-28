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
            "firstName": "Mauricio",
            "lastName": "Bergallo",
            "identificationNumber": "31174646",
            "email": "bergallo.mauricio@gmail.com",
            "birthDate": "1984-09-12T00:00:00Z"
        },
        {
            "id": "9b8e22e1-94ac-4b54-bb86-2b8260c15c77",
            "firstName": "Cintia",
            "lastName": "D'Allotta",
            "identificationNumber": "308762323",
            "email": "cinty.dallotta@gmail.com",
            "birthDate": "1984-03-04T00:00:00Z"
        }
    ]
}
```

## Creating new Students
```
POST: http://localhost:8080/api/students
Payload:
{
    "firstName": "Capo",
    "lastName": "Vieja",
    "identificationNumber": "13130994",
    "email": "capo.vieja@gmail.com",
    "birthDate": "1990-10-11T00:00:00Z"
}
```