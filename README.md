# grpc-http-rest-microservice
Source code for test task. Develop Go gRPC microservice with HTTP/REST endpoint, middleware, Kubernetes deployment, etc.

Env configuration

| Name                                          | Required | Default                 | Destination                        |
|-----------------------------------------------|----------|-------------------------|------------------------------------|
| GRPC_PORT                                     | yes      | 9090                    | gRPC server port                   |
| HTTP_PORT                                     | yes      | 8080                    | HTTP server Port                   |
| REGENERATE                                    | yes      | 5 min                   | Time regenerate UUID               |


Build
--
Build service by Makefile
> make
