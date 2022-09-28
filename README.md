1. Create DB `userdb_task`; user: `postgres`; pass: `postgres`
2. Create table `users` using following SQL command:
```
 CREATE TABLE users
(
  username VARCHAR NOT NULL CONSTRAINT users_pk PRIMARY KEY,
  login VARCHAR,
  pass VARCHAR
)
```
3. Run `main.go`
```
go run main.go
```
