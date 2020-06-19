# benchmark

对比 grpc、grpc-gateway、http、go-kit 的性能测试。

## grpc

```
./grpc-gateway/bin/hello

./grpc-gateway/bin/client -c 1 -n 10000
```

## grpc-gateway

### proxy模式

```
./grpc-gateway/bin/hello --inprocess=false --http=":8899"

ab -c 1 -n 10000 http://127.0.0.1:8899/say
```

### in-process模式

```
./grpc-gateway/bin/hello --inprocess=true --http=":8899"

ab -c 1 -n 10000 http://127.0.0.1:8899/say
```

## http

```
./http/bin/http --http=":7788"

ab -c 1 -n 10000 http://127.0.0.1:7788/say
```

## go-kit

```
```

