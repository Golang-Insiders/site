# Golang Inisiders Site

[![Fly Deploy Site](https://github.com/Golang-Insiders/site/actions/workflows/deploy_site.yaml/badge.svg?branch=main)](https://github.com/Golang-Insiders/site/actions/workflows/deploy_site.yaml)

## Getting started
### Requirements
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [air](https://github.com/cosmtrek/air)
- [goose](https://github.com/pressly/goose)

### Running
Copy .env.example to .env
```sh
cp .env.example .env
```

Using make
```sh
make up
```

Show available commands
```sh
make help
```
