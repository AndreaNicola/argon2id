# Argon2id stupid wrapper

## Setup

```bash
go get -u github.com/AndreaNicola/argon2id
```

## Generate Hash

```go
hash, hashErr := argon2id.Hash(plainText)
```

## Verify Hash

```go
res, verErr := argon2id.Verify(plainText, hash)
```