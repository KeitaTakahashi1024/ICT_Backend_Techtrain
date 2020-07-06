CREATE DATABASE mysql_users;
USE mysql_users;

CREATE TABLE users(
  id              INTEGER PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(20) NOT NULL,
  token           VARCHAR(40) NOT NULL
);

INSERT INTO users (name, token) VALUES ("安倍 晋三", "0921");
INSERT INTO users (name, token) VALUES ("習 近平", "0615");
INSERT INTO users (name, token) VALUES ("文 在寅", "0124");