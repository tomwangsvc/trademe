# build service

```
go fmt trademe/... && go build -o ./bin/trademe ./cmd
```

# run service
```
./bin/trademe test1
```
test1 result will show as default but subcommand can also support `test2` or `test3` or `test4` or `test5`

# run service
```
 go test -timeout 5000ms -cover trademe/...
 ```