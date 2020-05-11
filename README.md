# GO-challenge

## Backend

1. Create basic API REST to render domains. `domain/yourdomain`
2. Once that we have the `domain` we can create:
2.1 GET request to [SSLLABS](https://api.ssllabs.com/) to get the servers info.
2.2 GET request to [ipInfo](https://ipinfo.io/) to get the country and the organization name (API Key generated and Env variable created). I tried with [golang-packages]("github.com/likexian/whois-go") but there are some issues.
2.3 GET request to [metadata](https://home.urlmeta.org/) to get the image and the title of the webpage.

3. Create Cockroach database
3.1 Download and install from [CockroachLabs](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb-gorm.html), also I saw this [video](https://www.youtube.com/watch?v=6x9b0t-j1mM) and followed this [tutorial](https://kb.objectrocket.com/cockroachdb/how-to-retrieve-cockroachdb-record-using-golang-web-app-561)
3.2 Inside the project db folder run in a new Powershell `cockroach start --insecure --listen-addr=localhost:26257 --http-addr=localhost:8080`.
3.3 Open a new terminal and inside the project db folder run `cockroach sql --insecure`.
3.4 SQL commands.

```sql
CREATE USER IF NOT EXISTS juanc;
CREATE DATABASE domains;

GRANT ALL ON DATABASE domains TO juanc;

SET DATABASE = domains;
CREATE TABLE tbldomains(id SERIAL PRIMARY KEY, name VARCHAR, count INT);
INSERT INTO tbldomains (name, count) VALUES('pushdev',5);
SELECT * FROM tbldomains;
UPDATE tbldomains SET count = 5 WHERE id = '554066634883203073';
DROP TABLE tbldomains;
DELETE FROM tbldomains WHERE name='pushdev';
```

## Frontend

## How to run 

* `go run environment.go main.go`