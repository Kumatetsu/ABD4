# Quickstart - API

This is a golang API developed for ABD4 project at [ETNA](https://etna.io/)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

This API need Golang v1.9 or higher to run. You can download [here](https://golang.org/dl/).
Once is installed, try the following in your favorite command tool

```
go env
```
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
go build ABD4/API
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

### Test workflow

Tests are functionnal and are testing API endpoints.
For each road, we send mocked data and compare API's return with a preprared expected state.

mock.go:

```
	registerUserResponse := &server.Response{
		Status: 200,
		Data:   mock.PostUser.ToString(),
	}

    -----

    {
        Description:        "POST on /auth/register",
        URL:                "/auth/register",
        Method:             "POST",
        Handler:            handler.Register,
        ExpectedBody:       registerUserResponse.ToString(),
        ExpectedStatusCode: 200,
    },
```

main_test.go::TestUser

```
    w := httptest.NewRecorder()
    if test.Method == "POST" {
        // we post the same data we have recorded as expected
        body = strings.NewReader(mock.PostUser.ToString())
    }
    req, err = http.NewRequest(test.Method, test.URL, body)
    // we apply the handler and his middleware
    test.Handler(testApp.Ctx, w, req)
    if test.URL == "/auth/register" {
        // we compare expected data and API's return
        err := json.Unmarshal(w.Body.Bytes(), response)
        assert.NoError(err)
        err = json.Unmarshal([]byte(response.Data), user)
        assert.NoError(err)
        assert.NotNil(user.ID, test.Description)
        assert.Equal(mock.PostUser.Name, user.Name, test.Description)
        assert.Equal(mock.PostUser.Email, user.Email, test.Description)
        assert.Equal(mock.PostUser.Password, user.Password, test.Description)
        assert.Equal(mock.PostUser.Permission, user.Permission, test.Description)
        assert.Equal(mock.PostUser.Claim, user.Claim, test.Description)
    }
```


## Deployment

On a live system, the recommended way to deploy this API in a docker context. to follow...

## Built With

* [Golang](https://golang.org/) - The base language
* [gorilla/mux](https://github.com/gorilla/mux) - Rest Routing packages
* [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) - Json Web Token packages
* [boltdb/bolt](https://github.com/boltdb/bolt) - NoSql embedded database packages
* [stretchr/testify/assert](https://github.com/stretchr/testify) - testing assertion base on golang testing tools

## Authors

* **Aurélien Castellarnau** - *Initial work* -

See also the list of [contributors](https://github.com/kumatetsu/ABD4/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Based road

Response format:
```
{"status": int, "data": string|json, ?"message": string, ?"detail": string}
```

* GET /user
* GET /backup -> download the encrypted .dat database to backup it
* POST /auth/register - body: {"name": string, "email": string, "password": string, "permission": string}
    persist a new model.User into the .dat file
    return: model.User json formatted with unique ID or error if someting goes wrong
* POST /auth/login - body: {"email": string, "password":string}
    compare input and persisted users, if corresponding user is found
    a token is generate and returned into the data property of response
