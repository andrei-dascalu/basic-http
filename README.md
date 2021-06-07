# Basic HTTP #


Idiomatic Go http server

* servers only / endpoint
* uses template in `handler/template.html` to provide a HTML response
* can handler GET param `favoriteTree`
* if this parameter is not specified or is empty, will ask for your favorite tree in response

Execute tests

```bash
go test -v ./...
```

Start server on port `8000`

```bash
go run main.go
```
