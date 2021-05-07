# Argon2id stupid wrapper

## Setup

```bash
go get -u github.com/AndreaNicola/argon2id
```

Import

```go
import _ "github.com/AndreaNicola/argon2id"
```

## Generate Hash

```go
a2id := argon2id.NewArgon2ID()
hash, hashErr := a2id.Hash(plainText)
```

## Verify Hash

```go
res, verErr := a2id.Verify(plainText, hash)
```