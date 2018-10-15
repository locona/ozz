# ozz

## Install

1. Setup middlewares.

```
make compose
```

2. Install hydra commands.

```
cd $GOPATH/src/github.com/ory/hydra
HYDRA_LATEST=$(git describe --abbrev=0 --tags)
git checkout $HYDRA_LATEST

dep ensure -vendor-only
go install \
    -ldflags "-X github.com/ory/hydra/cmd.Version=$HYDRA_LATEST -X github.com/ory/hydra/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/hydra/cmd.GitHash=`git rev-parse HEAD`" \
    github.com/ory/hydra
```

3. Install oathkeeper commands.

```
cd $GOPATH/src/github.com/ory/oathkeeper
OATHKEEPER_LATEST=$(git describe --abbrev=0 --tags)
git checkout $OATHKEEPER_LATEST
dep ensure -vendor-only
go install \
    -ldflags "-X github.com/ory/oathkeeper/cmd.Version=$OATHKEEPER_LATEST -X github.com/ory/oathkeeper/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/oathkeeper/cmd.GitHash=`git rev-parse HEAD`" \
    github.com/ory/oathkeeper

```

4. hydra clients 作成

```
make hydra.clients.import
```

5. oathkeeper 作成

```
make oathkeeper.rules.import
```

## Getting Started

1. Run client .

```
make client
```

2. Run server.

```
make server
```

