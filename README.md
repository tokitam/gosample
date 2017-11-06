# gosample

ex...
```
$ export GOPATH=$HOME/your_working_dir/
$ cd $GOPATH
$ mkdir -p src/github.com/tokitam/
$ cd src/github.com/tokitam/
$ git clone git@github.com:tokitam/gosample.git
$ go install github.com/tokitam/gosample/helloworld/
$ bin/helloworld
Hello, World!
```


hello sample
```
$ go run helloworld/helloworld.go
Hello, World!
```

search prime number
```
$ go run prime_numbers/prime_numbers.go
2
3
5
.
.
.

83
89
97
```
