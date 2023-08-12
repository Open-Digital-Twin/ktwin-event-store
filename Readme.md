# KTWIN Event Store

KTWIN Event Store service.


## Build

```sh
docker build -t ktwin/event-store:0.1 .
```

## Load in Kind Development Environment

```sh
docker build -t dev.local/ktwin/event-store:0.1 .
kind load docker-image dev.local/ktwin/event-store:0.1
```
