# gRPC-web simple demo

This is a simplest echo request/response gRPC example. 

Due to the JavaScript code( HTTP v1) wouldn't post data to gRPC server by HTTP version 2, You have to use proxy to delegated the request of client, there are two proxy software one is `Envoy` and other is `Nginx`. we are using the second one for two consideration.
	1. Nginx is more famous than Envoy.
	2. Nginx's configuration is simpler than Envoy and has lot of documents and samples. 

![grpc-web](./asset/grpc-web.png)

## Server side

The gRPC server's duty is echo back the message which send from client and plus the prefix `Hello` that as the reply. 

To activate the server, you can run

```shell

 $ go run cmd/server/main.go

```

> NOTE:
> By default server listen on the tcp port 8090, but you can specified 
> the `-port` parameter flag as needed. 

## Proxy

The proxy is mandatory of the infrastructure of gRPC-web. we are running the proxy in the docker container

```shell

 $  cd docker
 $  docker-compose up

```

> NOTE:
> we are mapping the `conf.d/grpc_proxy.conf` as the additional configuration file of the nginx container. you can modify as your needs. Keep in mind , restart the docker by press `Ctrl + c` to kill the docker-compose process, and rerun it.


## Client side

If you want modified the code in `cmd/client/echo_client.js`, you have to use `webpack` to rebuild the minimal js version. 

Use webpack in the js folder. 

```shell
  $ npx webpack echo_client.js
```

We put the client html files of directory mapping to the volume of the nginx docker container. You need copy the `index.html` and `dist` to the `docker/html` which is the nginx document root from which serve to.


## The next...

We'll use react to write a new `index.html` page with input form and submit button. You are post the message that you inputed.