# worker

background worker that subscribes to nats streaming and act as a buffer simulation to create sale orders from partners.

## Prerequisites

This project requires

- [Go 1.16+](https://golang.org/)
- make

## Getting Started

### Installation

Run `make` to install dependencies.

```bash
make install
```

### Usage

Setup `.env` file according to template provided in `.env.sample`, then run locally using

```bash
make serve
```

