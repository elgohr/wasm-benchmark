```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ go get -u github.com/shurcooL/goexec
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'

```