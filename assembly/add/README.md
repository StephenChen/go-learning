```shell
go build -gcflags "-N -l" .
go tool objdump -s "main\." add
```