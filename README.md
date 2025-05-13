# mary
mary is gRPC server for logging toddler activity.

## Features
- Log activity (currently only supports food intake).
- Track API call into cache. Later, worker picks up the history and load it to DB.
- Determine whether user is registered without calling to DB.

## Usage
### Start server
```bash
./mary serve
```
### Create membership registry
Create bloom filter in cache based on existing users. It will check whether hashed username exists on cache without calling DB.
```bash
./mary create-membership-registry
```

## Minimum third party requirement
- PostgreSQL 15.6
- Valkey 8.0.1
- Dbmate 2.24.2
- sqlc 1.28.0
- Buf 1.50.0
- Mockery 2.53.3
- grpcui dev build
