![Daedalus Build Badge](https://github.com/smsunarto/daedalus/actions/workflows/build.yml/badge.svg)

# Daedalus

![Deadalus Cover](/.github/cover.png)

Deadalus is a modular toolkit for zkSNARKs development and deployment written in Go using Gnark.

```
daedalus
└── pkg
    ├── daedalus-cli   // WIP: ZK swiss-army knife: compile, trusted setup, deploy
    └── prover         // WIP: Generate proofs and send result to webhook endpoint
```

## Todo List

- [ ] Dockerize + Terraform
- [ ] Tests
  - [ ] Daedalus CLI
  - [ ] Prover

### Daedalus-CLI

- [x] Compile circuits
- [x] Perform trusted setup
- [ ] Deploy circuits to a persistence/storage layer

### Prover

- [x] Generate proof
- [ ] Poll task queue for proving request
- [ ] Call webhook endpoint supplied by task to submit proof
