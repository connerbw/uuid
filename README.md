# Demo UUID RPC plugin for Roadrunner 2.8

Hosted on: https://github.com/ls-dac-chartrand/uuid

## HOWTO

### Go

Using https://github.com/roadrunner-server/velox

Add this to your `plugins.toml` file:

```
uuid = { ref = "master", owner = "ls-dac-chartrand", repository = "uuid" }
```

And its dependencies:

```
rpc = { ref = "master", owner = "roadrunner-server", repository = "rpc" }
logger = { ref = "master", owner = "roadrunner-server", repository = "logger" }
```

### PHP


```php
$rpc = RPC::create(
    Environment::fromGlobals()->getRPCAddress()
);

$uuid = $rpc->call('uuid.Generate', 'not-used');
```
