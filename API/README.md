
# Quickstart - API



This is a golang API developed for ABD4 project at [ETNA](https://etna.io/)



## Getting Started



These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.



### Prerequisites



This API need Golang v1.9 or higher to run. You can download [here](https://golang.org/dl/).

Once is installed, try the following in your favorite command tool


``` go env ```

Pay attention to [GOPATH, GOBIN and GOROOT](https://www.programming-books.io/essential/go/10-gopath-goroot-gobin) or [official GOPATH documentation](https://github.com/golang/go/wiki/GOPATH).



### Installing



First of all, you need to setup your Golang environnement.

The minimum to setup is your GOPATH. Follow instructions [here](https://github.com/golang/go/wiki/SettingGOPATH)



Once you have your GOPATH, go into GOPATH/src/ folder and type the following:



```

git clone https://github.com/Kumatetsu/ABD4.git

cd ABD4/API

go get github.com/gorilla/mux

go get github.com/dgrijalva/jwt-go

go get github.com/stretchr/testify/assert

// to use with BoltDb embedded key/value database

go get github.com/boltdb/bolt/...

// to use with MongoDB higly scalable document oriented database on distant server

go get gopkg.in/mgo.v2

go get gopkg.in/olivere/elastic.v5

go build ABD4/API

```



## Running du docker contenant elasticsearch 5.5.3

```

cd docker-es

- if not exists, create dir `data` and `data2`

- launch the docker:

docker-compose up

- kill the docker:

docker-compose down [--remove-orphans, kill all instances if changes has been done in docker-compose.yml]

```



Warning: if you want to run a `local` docker set or `unset the sniffing option` of elastic library.

```

docker-compose.yml ->

services:

my_service:

environment:

http.publish_host=127.0.0.1

...

...

...

...

```



## Commandes Es pour voir les infos

```

- voir les différents nodes (normalement 2):

curl http://127.0.0.1:9200/_nodes/http?pretty=1

- voir le cluster:

curl http://127.0.0.1:9200/

- voir le cluster health:

curl http://127.0.0.1:9200/_cluster/health

```



## Running the tests



In GOPATH/src/ABD4/API



```

go test -v

```



### Break down into end to end tests



Tests will launch an instance of the app then try request with mocked data.

The process will create a temporary .dat file under /test folder. This file is erase in process.

The process will log server behaviour under /test folder. The abd4.log file is created/appened.



Expected console output:



![good_output](https://user-images.githubusercontent.com/16307418/46264432-11e8e800-c51d-11e8-8c8f-758280c16fc3.png)


## Built With



*  [Golang](https://golang.org/) - The base language

*  [gorilla/mux](https://github.com/gorilla/mux) - Rest Routing packages

*  [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) - Json Web Token packages

*  [olivere/elastic.v5](https://godoc.org/gopkg.in/olivere/elastic.v5) - Package elastic provides an interface to the Elasticsearch server

*  [gopkg.in/mgo.v2](https://godoc.org/gopkg.in/mgo.v2) - Package mgo offers a rich MongoDB driver for Go

*  [boltdb/bolt](https://github.com/boltdb/bolt) - NoSql embedded database packages

*  [stretchr/testify/assert](https://github.com/stretchr/testify) - testing assertion base on golang testing tools



## Authors



*  **Aurélien Castellarnau** - *Initial work* -



See also the list of [contributors](https://github.com/kumatetsu/ABD4/contributors) who participated in this project.



## License



This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details



## Based road



Response format:

```

{"status": int, "data": json, ?"message": string, ?"detail": string}

```



* GET /user

  -    protected: false

   -  action: FindAll from IUserManager

   - return: [] model.User[json]

   - status: 200 | 500

* DELETE /user

   - protected: false

  - action: IUserManager.RemoveAll + AppContext.RemoveIndex

  - return: message[string]

  - status: 202 | 500

* OPTIONS /auth/register

* POST /auth/register - body: {"name": string, "email": string, "password": string, "permission": string} -

  - protected: false

  - action: IUserManager.Create + AppContext.IndexData

  - return: model.User[json] or error

  - status: 201 | 500

* OPTIONS /auth/login

* POST /auth/login - body: {"email": string, "password":string} -

  - protected: false

  - action: IUserManager.FindOneBy + Authentication

  - return: jwt_token[string] or error

  - status: 200 | 400 | 401 | 500

* GET /transaction

  - protected: false

  - action: ITransactionManager.FindAll

  - return: [] model.Transaction[json]

  - status: 200 | 500

* OPTIONS /transaction

* POST /transaction - body: see exemple_reservation.json -

  - protected: false

  - action: ITransactionManager.Create + AppContext.IndexData

  - return: model.Transaction[json]

  - status: 201 | 500

* DELETE /transaction

  - protected: false

  - action: ITransactionManager.RemoveAll + AppContext.RemoveIndex

  - return: message[string]

  - status: 202 | 500

* GET /elastic/index/all

  - protected: false

  - action: elasticsearch.CreateIndexation

  - return: indexes[string]

  - status: 200 | 500

* GET /elastic/index/{index}

  - protected: false

  - action: mux.Vars + elasticsearch.Index

  - return: index[string]

  - status: 200 | 400 | 500

* GET /elastic/rmindex/all

  - protected: false

  - action: for on context.INDEXES, apply elasticsearch.RemoveIndex

  - return: message[string]

  - status: 200 | 500

* GET /elastic/rmindex/{index}

  - protected: false

  - action: mux.Vars + elasticsearch.RemoveIndex

  - return: message[string]

  - status: 200 | 500

* GET /elastic/reindex/all

  - protected: false

  - action: elasticsearch.CreateIndexation [reindex==true]

  - return: message[string]

  - status: 200 | 500

* GET /elastic/reindex/{index}

  - protected: false

  - action: mux.Vars + elasticsearch.RemoveIndex + elasticsearch.Index

  - return: message[string]

  - status: 200 | 400 | 500

* GET /elastic/indexdata

  - protected: false

  - handler: GetIndexationData

  - return: message[string]

  - status: 200 | 500

GET /elastic/indexdata/{index}

  - protected: false

  - handler: GetIndexData

  - return: message[string]

  - status: 200 | 400 | 500

