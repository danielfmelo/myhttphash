# MYHTTPHASH

This is a simple application wrote in Golang to hash a list of websites body responses. For the hash, it uses the [MD5 algorithm](https://en.wikipedia.org/wiki/MD5).

## Prerequisites

You must have installed the Golang version >= 1.14 or only the [Docker](https://www.docker.com/).

## How to use

To make easy run and build the application you can use make commands. To avoid issues with Golang version I recommend you use the `docker-` parameters.

### Running

You can run the application using the following commands and parameters:

```shell
make run parallel=5 urls="https://www.google.com https://www.globo.com"
make docker-run parallel=5 urls="https://www.google.com https://www.globo.com"
```

The urls must be given with a space between them. If the url doesn't have the scheme (http or https), it will consider `https`.

### Building

The same way to build the app, you can use either your Golang version or the docker one:

```shell 
make build
make docker-build
```

It will create an executable in the project root named myhttphash.

### Running the executable

If you want to run the executable, you can do that doing it:

```shell
./myhttphash -parallel 4 https://www.google.com https://www.globo.com
```

If the parameter `-parallel` is not informed, the value default is `10`.


## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details
