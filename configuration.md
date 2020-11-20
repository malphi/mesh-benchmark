### Istio configuration

```
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: httpbin
  namespace: default
spec:
  hosts:
  - httpbin
  http:
  - route:
    - destination:
        host: httpbin
      weight: 100
```
istio envoy

```
Containers:
  goclient:
    ...
  istio-proxy:
    Container ID:  docker://e90df3b903cd6fedcc6df4e295b460b3e27f5a42651422de024378f1461b3ea5
    Image:         docker.io/istio/proxyv2:1.7.0
    ...
```
do a test make sure it works

```
curl -i http://localhost:31952/byte/100
HTTP/1.1 200 OK
date: Thu, 19 Nov 2020 01:59:43 GMT
content-length: 100
content-type: application/octet-stream
x-envoy-upstream-service-time: 35
server: istio-envoy
x-envoy-decorator-operation: goclient.default.svc.cluster.local:8888/*

envoy log:

2020-11-19T01:59:43.709921Z	debug	envoy router	[C204][S13954033629459345604] cluster 'inbound|8000|http|httpbin.default.svc.cluster.local' match for URL '/bytes/100'
2020-11-19T01:59:43.709955Z	debug	envoy router	[C204][S13954033629459345604] router decoding headers:
':authority', 'httpbin:8000'
':path', '/bytes/100'
':method', 'GET'
':scheme', 'http'
'user-agent', 'Go-http-client/1.1'
```

### appmesh configuration

```
k get mesh
NAME                     ARN                                                                       AGE
mymesh                   arn:aws:appmesh:ap-northeast-1:054818973631:mesh/mymesh                   5d
virtualrouter.appmesh.k8s.aws/httpbin-router   arn:aws:appmesh:ap-northeast-1:054818973631:mesh/mymesh/virtualRouter/httpbin-router_default   3h25m

NAME                                   ARN                                                                                    AGE
virtualnode.appmesh.k8s.aws/goclient   arn:aws:appmesh:ap-northeast-1:054818973631:mesh/mymesh/virtualNode/goclient_default   3h21m
virtualnode.appmesh.k8s.aws/httpbin    arn:aws:appmesh:ap-northeast-1:054818973631:mesh/mymesh/virtualNode/httpbin_default    3h21m

NAME                                                               ARN                                                                                                        AGE
virtualservice.appmesh.k8s.aws/httpbin.default.svc.cluster.local   arn:aws:appmesh:ap-northeast-1:054818973631:mesh/mymesh/virtualService/httpbin.default.svc.cluster.local   3h25m

Containers:
  goclient:
    ...
  envoy:
    Container ID:   docker://edc06a842f0764c67c253c8dc6d9ca11fc0e5d46a2358fd6a20e3691c1223831
    Image:          840364872350.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-envoy:v1.15.1.0-prod
    Image ID:       docker-pullable://840364872350.dkr.ecr.us-west-2.amazonaws.com/aws-appmesh-envoy@sha256:183159918171543f9b1c7234d0c156b632aeca571d3d9a9cd1895e1f21ab94ae
    Port:           9901/TCP
    ...
```
