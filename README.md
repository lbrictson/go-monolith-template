# go-monolith-template
An opinionated starter template for a Go web application, designed for the local developer experience.

This template is a great starting place to develop open source tools and applications.  By default it uses a sqlite3 database
implementation, however postgres is fully supported in the configuration file.


## Features

- [x] Docker Compose for local development
- [x] Makefile for common tasks
- [x] Ent for database schema and automatic migrations
- [x] Mailhog for email testing
- [ ] Multifactor authentication (TOTP)
- [ ] Authentication middleware
- [ ] Templ for html templating
- [x] Slog for structured logging
- [x] Hot reload with Air
- [x] GitHub Actions testing
- [x] Environment variables for configuration

## Required Tools

- <a href="https://www.docker.com/" target="_blank">Docker</a>
- <a href="https://docs.docker.com/compose/" target="_blank">Docker Compose</a>
- <a href="https://www.gnu.org/software/make/" target="_blank">Make</a>
- <a href="https://golang.org/" target="_blank">Go</a>
- <a href="https://templ.guide/quick-start/installation" target="_blank">Templ</a>
- <a href="https://github.com/cosmtrek/air" target="_blank">Air</a>

## Quick Start

The fastest way to get started is to work with hot reloading so your changes are (almost) instantly reflected in the running application.

```bash
make dependencies # Only needed once
make hot
```

There is always a default user seeded as part of the startup process with the username `admin@example.com` and the password `Password1234!`.
This user will be auto recreated if the database contains no users; it will not be created if the database contains at least one 
existing user.

Access the application at http://localhost:8080

### Adding New Database Models

```bash
go run -mod=mod entgo.io/ent/cmd/ent new ModelInCamelCase
```

Note:  Database migrations are automatic and run at the start of the application.  If you need to modify the migration process
or disable automatic runs you can remove the call in `main.go` and run the migrations manually.

### Modifying Web Templates

Templates are located in the `templates` directory and use the `templ` package for rendering.  The `templ` package is a wrapper around the `html/template` package that provides a few extra features.

Feel free to modify as needed, read the docs here: <a href="https://templ.guide/" target="_blank">Templ Guide</a>

## Repo Structure

`cmd/server` - Entry point for the application, runs configuration and launches the web server.

`ent` - Database schema and migrations, generated code.
`ent/schema` - Database schema definitions (You create and modify the schema here).

`local` - Local storage of the sqlite3 database and other local files.

`pkg` - Application code.
`pkg/logging` - Global logger and logging helpers.
`pkg/config` - Configuration loading and validation.

`templates` - Web templates.

`web` - Static web assets like javascript, css, and images, this directory gets bundled into the final binary

`.air.toml` - Air configuration file for hot reloading.


## Make Commands

Use `make help` to find commands for local development.  Generally you just need `make run` to start and `make test`


