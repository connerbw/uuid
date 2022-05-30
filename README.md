# Demo UUID RPC plugin for Roadrunner 2.8

Hosted on: https://github.com/connerbw/uuid

## HOWTO

### Go

Using https://github.com/roadrunner-server/velox

Add this to your `plugins.toml` file:

```
uuid = { ref = "master", owner = "connerbw", repository = "uuid" }
```

And its dependencies:

```
rpc = { ref = "master", owner = "roadrunner-server", repository = "rpc" }
logger = { ref = "master", owner = "roadrunner-server", repository = "logger" }
```

#### Docker

```Dockerfile
FROM golang:1.18 as roadrunner
ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /go/src/roadrunner
COPY ./path/to/plugins.toml .
RUN go install github.com/roadrunner-server/velox/vx@v1.0.0-beta.1
RUN vx build -c plugins.toml -o /go/src/roadrunner/rr

### FROM php:8.0.17-cli-alpine3.15
### ... INSERT PHP STEPS HERE ...

COPY --from=roadrunner /go/src/roadrunner/rr /usr/local/bin
CMD ["/usr/local/bin/rr", "serve", "-c", "/var/www/.rr.yaml"]
```

### PHP


```php
use Spiral\Goridge\RPC\RPC;
use Spiral\RoadRunner\Environment;

// ...

$rpc = RPC::create(Environment::fromGlobals()->getRPCAddress());
$uuid = $rpc->call('uuid.Generate', 'not-used');
```


