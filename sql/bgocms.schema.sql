DROP DATABASE bgocms;
CREATE DATABASE bgocms;
\c bgocms

CREATE TABLE users(
  id          SERIAL PRIMARY KEY,
  uname       VARCHAR NOT NULL UNIQUE,
  pass        VARCHAR NOT NULL
);

CREATE TABLE sessions(
  id          SERIAL PRIMARY KEY,
  user_id     INT references users(id),
  token       VARCHAR NOT NULL,
  expires_at  FLOAT NOT NULL
);

CREATE TABLE articles(
  id          SERIAL PRIMARY KEY,
  author      INT references users(id),
  title       TEXT,
  body        TEXT,
  created     TIME NULL,
  last_edited TIME NULL
);

