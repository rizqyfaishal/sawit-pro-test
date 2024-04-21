# Backend Engineering Interview Assignment (Golang)

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.20
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

    Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project

To start working, execute

```
make init
```

## Genereate Private And Public Key

Since the project usign RS256 as signing method on JWT, it need private and public key

You can generate private and public key using this following commands;


```
mkdir cert
openssl genrsa -out cert/id_rsa 4096
openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

Current testing strategy is using unit test without out-of-process dependencies (db, storage, etc). 

Those
dependencies will be replaced by mock to ensure the unit test result deterministic outcome, fast, 
and reduce flaky test.

To run test, run the following command:

```
make test
```

To run test, and get the test coverage run the following command:

```
make coverage
```

Current unit test coverage is at 85% (high)

## Notes

Since there is this update https://github.com/deepmap/oapi-codegen/releases/tag/v1.14.0 , the minimum of golang version 
will be update with minimum Go 1.20

## Mini Review Video