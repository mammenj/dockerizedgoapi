# Customized version of Learning Cloud Native Go - myapp

### Endpoints
- To see the app is up
    GET /


- Domain API
    GET /api/v1/books
    POST /api/v1/books
    GET /api/v1/books/{id}
    PUT /api/v1/books/{id}
    DELETE /api/v1/books/{id}

    GET /users
    GET /users/{id}
    POST /users
    
    POST /token/{userid}/{password}
<!--  -->
### Authentication
- Added JWT authentication

### Authorization
- Added Casbin for declartive authorization

### TO DO
- External authentication
- Role based access to API