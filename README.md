# pipe-and-filter
This Go program prints a list of prime numbers until it is terminated.

The program uses Go routines and channels to implement a basic Pipe and Filter pattern to calculate prime numbers. Typicall output is shown below.

```
1596377
1596379
1596383
1596389
1596433
1596451
1596467
1596493
1596503
1596509
```

## Usage

```bash
cd ./pipe-and-filter
go build
./pipe-and-filter
```

## Credit

Source: [Go Concurrent Programming](http://www.pluralsight.com/courses/go-concurrent-programming)
