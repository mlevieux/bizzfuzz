# BizzFuzz - A fizz-buzz sequence generation server

## Usage

Make sure all dependencies are installed using:

```
go get
```

at the root of the project.


Once all dependencies are installed, you can build the project using:

```
go build
```

which will create a binary at the root of the project. Then you only need to run the given binary using

```
./bizzfuzz
```

Additionally, you can use the -a flag to specify an address to listen to. By default, the address will be localhost:8080.

The server is running after that and you can call it directly.

## Endpoints

The server has two endpoints:

/fizzbuzz

which takes 5 parameters: int1, int2, limit, str1, and str2, such that it returns the fizzbuzz sequence corresponding to the
given parameters, as specified in the project specification.

/statistics

which can take a 'top' parameter. If top is not given, it will be automatically assigned the value 1.
This endpoint returns the most 'top' called parameter sets on /fizzbuzz endpoint, along with the number of times they've been called.