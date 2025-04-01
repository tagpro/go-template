# A go server template

## Create your new project

Install gonew

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

Create your new project

```sh
gonew github.com/tagpro/go-template <example.com/your.go-project.name>
```

## Roadmap

- [ ] Add docker
- [ ] Add healthcheck
- [ ] Add logging
- [ ] Add Tracing
- [ ] Add Monitoring
- [ ] Add OpenAPI

## Description

This is a template for a go application that serves go APIs using chi router.

The aim is to reduce the amount of boilerplate code for application I create.
Also, this also serves as a good source of example to be used in other go applications.

## Libraries

- [go-chi](https://github.com/go-chi/chi)
- [sqlc](https://github.com/sqlc-dev/sqlc) along with [goose](https://github.com/pressly/goose)
- [templ](https://github.com/a-h/templ)

## Database setup

1.  Use Goose to generate migration files for your database. An initial migration file without any DDL queries is provided. Add your initial database DDL statements there.
2.  Run `make db/up` to apply the migration to the latest version (the initial migration in this case).
3.  Run `make db/gen` to generate the SQLC code based on your database schema.
4.  After generating the SQLC code, run `go mod tidy` to install any new dependencies introduced by SQLC.
