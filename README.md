# redis

A lightweight Redis-inspired in-memory database written in Go.

This project is built to understand how Redis works internally:
single-threaded execution, in-memory data structures, and TTL-based expiration.

## Features

- TCP server
- Concurrent client handling
- Single-threaded event loop
- In-memory key-value store
- TTL (key expiration support)

### Data Structures Supported
- Strings
- Hashes
- Lists

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

## Strings

### SET
```text
SET name john [ttl_seconds(optional)]
```

### GET
```text
GET name
```

### DEL
```text
DEL name
```

### MGET
```text
MGET name age city
```

---

## Hashes

### HSET
```text
HSET user name john
```

### HGET
```text
HGET user name
```

---

## Lists

### LADD (append to list)
```text
LADD mylist a
LADD mylist b
```

### LRANGE
```text
LRANGE mylist 0 10
```

### INDEX
```text
INDEX mylist 0
```

---

## TTL Support

### EXPIRE
```text
EXPIRE key seconds
```

---

### TTL
```text
TTL key
```

Returns:
- remaining seconds
- -1 → no TTL
- -2 → key not found / expired

---

### PERSIST
```text
PERSIST key
```

Removes TTL from key.

---

### TTL Behavior
- TTL applies to the **entire key**
- Not per field (hash/list items are not individually expirable)
- Expired keys are removed via:
  - lazy deletion on access
  - background cleanup goroutine

---

## Architecture

The server follows a Redis-inspired architecture:

```text
TCP Clients
    ↓
Connection Goroutines
    ↓
Command Parser
    ↓
Channel Queue
    ↓
Single Event Loop (Worker)
    ↓
In-Memory Store
```

Only the event loop goroutine mutates the datastore, which avoids locks and race conditions.
