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

## Get ScyllaDB credentials from secrets

```sh
kubectl get secrets/scylla-auth-token -n scylla --template={{.data}}
```

## Expose ScyllaDB Port

```sh
kubectl port-forward --address 0.0.0.0 -n scylla svc/scylla-client 9042:9042
```
