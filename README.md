# envoy-ingress-controller

基于envoy实现的Kubernetes Ingress Controller。

## 功能特性
***
* HTTP流量转发，支持HTTP2.0, Grpc
* 自定义路由规则，支持根据host、path、headers进行路由和重定向
* 多组service灰度
* HTTP、TCP 探针
* 7层LB日志