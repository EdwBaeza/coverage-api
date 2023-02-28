Coverage API

Run tests

```console
$ docker-compose run --rm app go test -v ./...
```

Run server and build for develop

```console
$ docker-compose up --build
```

Build app

```console
$ docker-compose build
```

Run app

```console
$ docker-compose up
```

Run mongo console
```console
$ docker-compose run --rm mongodb mongosh mongodb://mongodb:27017/
```

 