# GO-challenge

## Backend

1. Create basic API REST to render domains. `domain/yourdomain`
2. Once that we have the `domain` we can create:
* GET request to [SSLLABS](https://api.ssllabs.com/) to get the servers info. You can see the docs [here](https://github.com/ssllabs/ssllabs-scan/blob/master/ssllabs-api-docs-v3.md).
* GET request to [ipInfo](https://ipinfo.io/) to get the country and the organization name (API Key generated and Env variable created). I tried with [golang-packages](https://github.com/likexian/whois-go) but there are some issues.
* GET request to [metadata](https://urlmeta.org/) to get the image and the title of the webpage.

3. Create Cockroach database
* Download and install from [CockroachLabs](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb-gorm.html), also I saw this [video](https://www.youtube.com/watch?v=6x9b0t-j1mM) and followed this [tutorial](https://kb.objectrocket.com/cockroachdb/how-to-retrieve-cockroachdb-record-using-golang-web-app-561).

### SQL commands.

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

1. Install Vue. With nodejs run 

* `npm i -g @vue/cli`
* `vue create domains-app`
* `cd domains-app`
* `npm i vue-resource --save`
* [How to install bootstrap](https://bootstrap-vue.org/docs)
 
## How to run 

* Inside the project db folder run in a new Powershell `cockroach start --insecure --listen-addr=localhost:26257 --http-addr=localhost:8080`.
* Open a new terminal and inside the project db folder run `cockroach sql --insecure`.
* `go run environment.go main.go`. Open `http://localhost:8081/`
* `npm run serve -- --port 3000`. Open `http://localhost:3000/`

Enjoy!

### Type the domain

![image](https://user-images.githubusercontent.com/36536646/81629276-5eb6c500-93c8-11ea-81e9-17a1e32172a3.png)

### See the results

![image](https://user-images.githubusercontent.com/36536646/81629313-7aba6680-93c8-11ea-9af0-ab61a4b4782a.png)

![image](https://user-images.githubusercontent.com/36536646/81629354-91f95400-93c8-11ea-871a-f07f7cc570f6.png)

### Search history

![image](https://user-images.githubusercontent.com/36536646/81629427-bb19e480-93c8-11ea-8398-569c72582ce8.png)

