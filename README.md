# abagile-go

A comprehensive Go utility library that provides a curated set of helper functions, primarily wrapping the excellent [samber/lo](https://github.com/samber/lo) library while adding custom extensions for concurrency, event buses, and database operations.

## Overview

`abagile-go` is designed to be a "standard library plus" for abagile projects. It provides a clean, unified API surface for common tasks:

- **Type Manipulation:** Pointer helpers, zero-value checks, and type conversions.
- **Collections:** Functional programming helpers (Map, Filter, etc.) via `lo`.
- **Concurrency:** Worker pools and concurrent job execution.
- **Event Bus:** Simple, type-safe in-process event pub/sub.
- **Database:** Helpers for `pgx` and SQL placeholder conversion.
- **Math & Errors:** Generic math utilities and "Must" style error handling.

## Installation

```bash
go get github.com/abagile/go
```

## Usage Examples

### Functional Helpers (via lo)

```go
import "github.com/abagile/go"

names := []string{"alice", "bob", "charlie"}

// Map to uppercase
upperNames := abagile.Map(names, func(s string, _ int) string {
    return strings.ToUpper(s)
})

// Filter
shortNames := abagile.Filter(names, func(s string, _ int) bool {
    return len(s) <= 5
})
```

### Type Manipulation

```go
// Safely handle pointers
val := "hello"
ptr := abagile.ToPtr(val)
original := abagile.FromPtr(ptr)

// Ternary helper
result := abagile.Ternary(len(names) > 0, "has users", "empty")
```

### Concurrency

```go
jobs := []int{1, 2, 3, 4, 5}
workerCount := 2

// Process jobs concurrently and collect results
results := abagile.RunConcurrent(jobs, workerCount, func(i int) int {
    return i * 10
})
```

### Event Bus

```go
bus := abagile.NewEventBus()

type MyEvent struct { Message string }

// Subscribe
abagile.Subscribe(bus, func(ctx context.Context, e MyEvent) {
    fmt.Println("Received:", e.Message)
})

// Publish
abagile.Publish(context.Background(), bus, MyEvent{Message: "Hello World"})
```

## Attribution

This library heavily utilizes [samber/lo](https://github.com/samber/lo). We wrap these functions to provide a stable, curated API for our internal projects while maintaining the ability to extend them.
