# Bookings and Reservations

This is the repository for my bookings and reservations project

    Built in Go version 1.20.5

Dependencies:

-   [chi router](github.com/go-chi/chi)
-   [alex edwards SCS session management](github.com/alexedwards/scs/v2)
-   [nosurf](github.com/justinas/nosurf)
-   [pgx](https://github.com/jackc/pgx)
-   [simple mail](https://github.com/xhit/go-simple-mail)
-   [Go validator](https://github.com/asaskevich/govalidator)

In order to build and run this application, it is necessary to install Soda (go install github.com/gobuffalo/pop/... ), create a postgres database, fill in the correct values in database.yml, and then run:

`soda migrate`

Then, fill in the correct values in .env
Then run run.bat file

Project Content:

1. Routing & Middleware
2. State management with sessions
3. Working with forms
4. Handlers
5. Testing functions
6. Error handling
7. Persisting Data with PostgreSQL
8. Sending mail using Go
9. Authentication
10. Setting up secure back end administration
