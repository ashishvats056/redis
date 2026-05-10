# redis

A lightweight Redis-inspired in-memory database written in Go.

## Features

- TCP server
- Concurrent client handling
- Single-threaded event loop
- In-memory key-value store
- Redis-style command execution
- TTL for keys

## Prerequisites

- Go 1.23.1 or newer
- `ncat` or any TCP client

## Local Setup

### 1. Clone the repository

```bash
git clone <repo-url>
cd redis
```

### 2. Run the server

```bash
go run .
```

Server starts on:

```text
localhost:6379
```

### 3. Connect using ncat

Open a new terminal:

```bash
ncat localhost 6379
```

## Supported Commands

### SET

```text
SET name john
```

### GET

```text
GET name
```

### DEL

```text
DEL name
```

## Architecture

The server follows a Redis-inspired architecture:

```text
TCP Clients
    ↓
Connection Goroutines
    ↓
Command Queue (channel)
    ↓
Single Event Loop
    ↓
In-Memory Store
```

Only the event loop goroutine mutates the datastore, which avoids locks and race conditions.

## Roadmap

- [ ] MGET
- [ ] HSET / HGET
- [ ] Lists
- [ ] RESP protocol support
- [ ] redis-cli compatibility
- [ ] Persistence
- [ ] TTL support