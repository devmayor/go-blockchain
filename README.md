# Go Blockchain

[![CI](https://github.com/devmayor/go-blockchain/actions/workflows/ci.yml/badge.svg)](https://github.com/devmayor/go-blockchain/actions/workflows/ci.yml)

An educational blockchain node written in Go while working through Lukas
Lukac's *Build a Blockchain from Scratch in Go* tutorial. This repository is a
guided implementation, not an original blockchain protocol and not a
production blockchain.

The code demonstrates how common blockchain components fit together:

- an append-only block database and account state;
- proof-of-work mining and block validation;
- signed transactions, nonces, and encrypted keystores;
- a transaction pool and mining loop;
- peer discovery and block synchronization; and
- a CLI and HTTP API for running a local learning network.

## Project status

This is a preserved learning project based on the tutorial's `c12_crypto`
checkpoint. It is useful for reading, experimentation, and local development.
It has not been security-audited and must not be used to hold funds, manage real
keys, or run an internet-facing network.

The HTTP transaction endpoint accepts a keystore password so it can sign a
transaction on the local node. That design is intentionally simple for the
tutorial and is unsafe for a real service. Keep the node bound to localhost and
use test-only accounts.

## Requirements

- Go 1.22 or newer
- macOS, Linux, or Windows with a C toolchain available for the legacy
  go-ethereum dependency

## Run locally

```bash
git clone https://github.com/devmayor/go-blockchain.git
cd go-blockchain
go build -o bin/tbb ./cmd/tbb

./bin/tbb wallet new-account --datadir="$(pwd)/.data"
./bin/tbb run \
  --datadir="$(pwd)/.data" \
  --ip=127.0.0.1 \
  --port=8080 \
  --bootstrap-ip=127.0.0.1 \
  --bootstrap-port=8080
```

In another terminal:

```bash
curl http://127.0.0.1:8080/node/status
curl http://127.0.0.1:8080/balances/list
```

Run the fast validation suite:

```bash
go test -short ./...
```

The full node suite includes probabilistic proof-of-work and integration tests.
It may take several minutes and uses local network ports:

```bash
go test -tags=integration -p=1 -timeout=30m ./...
```

## Repository layout

```text
cmd/tbb/   CLI entry point
database/  blocks, transactions, state, and disk persistence
fs/        filesystem helpers
node/      HTTP API, peer sync, mempool, and mining
wallet/    account creation, signing, and verification
```

## Attribution and licenses

This project is derived from
[web3coach/the-blockchain-bar](https://github.com/web3coach/the-blockchain-bar)
by Lukas Lukac and follows the associated tutorial. Upstream attribution and
copyright remain with their respective owners.

As specified by upstream, code outside `cmd/` is licensed under GNU LGPL-3.0
and code inside `cmd/` is licensed under GNU GPL-3.0. See `COPYING`,
`COPYING.LESSER`, and `NOTICE`.
