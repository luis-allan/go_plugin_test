### Golang plugin test
- shared : shared(releasing with *.so) package
- plugin : plugin source package(produce *.so)
- main.go : test plugin

### Build
```bash
$ cd plugin
$ go build -buildmode=plugin hello.go
$ cd ..
$ go build
$ ./plugin-test
Hello test1...
Hello test3...
```