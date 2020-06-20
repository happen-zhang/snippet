# benchmark

对比 grpc-gateway、http、go-kit http 的性能测试。

## grpc-gateway

### proxy模式

```
./grpc-gateway/bin/hello --inprocess=false --http=":10221"

wrk -t8 -c100 -d1m "http://127.0.0.1:10221/say?word=world&repeatCount=10"
```

### in-process模式

```
./grpc-gateway/bin/hello --inprocess=true --http=":10221"

wrk -t8 -c100 -d1m "http://127.0.0.1:10221/say?word=world&repeatCount=10"
```

## http

```
./http/bin/http --http=":12221"

wrk -t8 -c100 -d1m "http://127.0.0.1:12221/say?word=world&repeatCount=10"
```

## go-kit

```
./go-kit/bin/hello --http="11221"

./grpc-gateway/bin/client -c 100 -n 100000 --word=world --repeat=10
```

## 运行环境

测试在单机上执行：

- go version：go1.14 darwin/amd64
- 操作系统：MacOSX 10.15.2 (19C57)
- 内存：16G
- 处理器：2.7 GHz 四核Intel Core i7

## 数据对比

| - | 平均(ms) | 中位(ms) | 最大(ms) | P90(ms) | P99(ms) | QPS |
| --- | --- | --- | --- | --- | --- | --- |
| go http | 0.96 | 0.88 | 84.67 | 1.11 | 2.53 | 104314.83 |
| grpc-gateway-proxy | 2.98 | 2.76 | 123.28 | 4.27 | 7.26 | 33853.49 |
| grpc-gateway-inprocess | 1.58 |0.89 | 158.46 | 3.26 | 8.50 | 80651.55 |
| go-kit http | 1.03 | 0.87 | 113.41 | 1.22 | 2.89 | 106041.51 |

上面的表格中可见，go http ≈ go-kit http > grpc-gateway-inprocess > grpc-gateway-proxy。

具体数值在不同机中有差异，视具体情况而定，但是总体趋势是有参考意义的。

## 参考

- [Feature Request: Add support for In-Process transport](https://github.com/grpc/grpc-go/issues/906)
- [Performance implications of using the grpc-gateway for a REST API](https://github.com/grpc-ecosystem/grpc-gateway/issues/1458)

