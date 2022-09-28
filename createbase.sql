CREATE DATABASE userdb_task;
\c userdb_task
CREATE TABLE users
(
  username VARCHAR NOT NULL CONSTRAINT users_pk PRIMARY KEY,
  login VARCHAR,
  pass VARCHAR
);
