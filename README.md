
# content

[![GoDoc](https://godoc.org/github.com/altipla-consulting/content?status.svg)](https://godoc.org/github.com/altipla-consulting/content)

> Models for translated content coming from multiple providers.


## Install

```shell
go get github.com/altipla-consulting/content
```

This library has no external dependencies outside the Go standard library.


## Usage

You can use the types of this package in your models structs when working with `database/sql`:

```
type MyModel struct {
  ID          int64      `db:"id,omitempty"`
  Name        content.Translated `db:"name"`
  Description content.Translated `db:"description"`
  Description content.Provider `db:"description"`
}
```


# Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using ```gofmt```.


# License

[MIT License](LICENSE)
