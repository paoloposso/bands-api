# Bands Auth API
The Bands Auth API serves as the auth service for the bands APP to be constructed.
The use cases are exposed via HTTP API, following the REST pattern.

So far it has the following features / use cases:

## User Register
Action: POST /api/user

Receives the basic user data:
- Name
- BirthDate
- Password

The password is encoded using bcrypt to be stored in the database.

All fields are required.

## User Login
Action: POST /api/user/login

Receives the email and the password in plain text and, if the user is correctly authenticated, returns a generated token.
Data:
- E-mail
- Password

## Enironment Variables
These are the Env Vars necessary to run the project, with sample data:
JWT_SECRET=XXAXXXXXASD453543
JWT_EXPIRY_MINUTES=30
PORT=5000
MONGO_URL=mongodb://localhost:27017
MONGO_DB=bands-db-sample
MONGO_TIMEOUT=5

## swagger
install swag via ```go get -u github.com/swaggo/swag/cmd/swag``` (already installed in this project)
run ```swag init``` to generate swagger files
swagger documentation can be accessed via /swagger/index.html
example: http://localhost:5000/swagger/index.html