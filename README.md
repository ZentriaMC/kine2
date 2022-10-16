# kine2 is discontinued

kine2 was planned to be needed for lightweight flannel deployments, but it didn't get very far.  
Besides, flannel now supports etcdv3 API (tried with v0.19.2), this project won't be needed anymore.

# kine2

kine2 is an etcd v2.3 API shim running on top of SQLite.

See https://etcd.io/docs/v2.3/api/

See `etcdserver/etcdhttp/client.go`

## Endpoints implemented

_None yet_

## Endpoints to implement

### GET /version

Response:
```json
{ "etcdserver":"2.3.8"
, "etcdcluster":"2.3.0"
}
```

### PUT /v2/keys/${path}

Payload:
```
```

Response:
```json
{ "action": "set"
, "node":
  { "createdIndex": 2
  , "key": "/${path}"
  , "modifiedIndex": 2
  , "value": "Hello world"
  }
, "prevNode?":
  { ... (see above) }
}
```
