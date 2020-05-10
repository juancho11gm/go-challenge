# GO-challenge

## Backend

### Create Basic REST API

### Create database

//new terminal
cockroach start --insecure --listen-addr=localhost:26257 --http-addr=localhost:8080

//new terminal
cockroach sql --insecure

CREATE USER IF NOT EXISTS juanc;
CREATE DATABASE domains;
GRANT ALL ON DATABASE domains TO juanc;
SET DATABASE = domains;
CREATE TABLE tbldomains(id INT PRIMARY KEY, name VARCHAR, count INT);

INSERT INTO tbldomains (id, name, count) VALUES(1,'pushdev',5);

SELECT * FROM tbldomains;
DROP TABLE tbldomains;

## Frontend