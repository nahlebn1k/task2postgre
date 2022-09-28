1. Create DB userdb_task
2. Create table `users` using following SQL command:
```
 CREATE TABLE users
(
  username VARCHAR NOT NULL CONSTRAINT users_pk PRIMARY KEY,
  login VARCHAR,
  pass VARCHAR
)
```
