```shell
go get github.com/better-maksim/counter
```

```shell
counter := NewCASCounter()
counter.Add(1)
value := counter.Read()

fmt.println(value)
```