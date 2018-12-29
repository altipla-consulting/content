
# content

> **DEPRECATED:** Use https://github.com/altipla-consulting/libs instead.

[![GoDoc](https://godoc.org/github.com/altipla-consulting/content?status.svg)](https://godoc.org/github.com/altipla-consulting/content)

> Models for translated content coming from multiple providers.


### Install

```shell
go get github.com/altipla-consulting/content
```

This library has no external dependencies outside the Go standard library.


### Usage

You can use the types of this package in your models structs when working with `database/sql`:

```go
type MyModel struct {
  ID          int64      `db:"id,omitempty"`
  Name        content.Translated `db:"name"`
  Description content.Translated `db:"description"`
  Description content.Provider `db:"description"`
}
```


### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using ```gofmt```.


### Running tests

Start the test database:

```shell
docker-compose up -d database
```

Install test libs:

```shell
go get github.com/stretchr/testify
go get upper.io/db.v3
```

Run the tests:

```shell
go test
```

Shutdown the database when finished testing:

```shell
docker-compose stop database
```


### License

[MIT License](LICENSE)
