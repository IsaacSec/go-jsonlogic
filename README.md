# JsonLogic Go Implementation

A Go implementation of [JsonLogic](https://jsonlogic.com), a standard for writing complex rules as portable JSON.

## Overview

This library allows you to write business logic in a JSON format that's:
- Platform independent
- Easy to read and write
- Safe to evaluate
- Portable between front-end and back-end

## Installation

```bash
go get github.com/IsaacSec/go-jsonlogic
```

## Usage

Basic usage example:

```go
import "github.com/IsaacSec/go-jsonlogic/jsonlogic"

rules := map[string]interface{}{
    "and": []interface{}{
        map[string]interface{}{"<": []interface{}{1, 2}},
        map[string]interface{}{">": []interface{}{1, 0}},
    },
}

data := map[string]interface{}{}

result, err := jsonlogic.Apply(rules, data)

if err != nil {
    // Handle error
}
// result will be true because (1 < 2) && (1 > 0)
```

## Supported Operations

### Access Operations
- `var`: Access data values using dot notation
- `missing`: (WIP) Check for missing keys
- `missing_some`: (WIP) Check for partial missing keys

### Logic Operations
- `and`: Returns true if all values are truthy
- `or`: Returns true if any value is truthy
- `!`: Returns the logical complement of the value
- `if`: Conditional branching

### Numeric Operations
- `>`, `>=`, `<`, `<=`: Numeric comparisons
- `+`, `-`, `*`, `/`: (WIP) Basic arithmetic
- `max`, `min`: (WIP) Maximum and minimum values
- `%`: (WIP) Modulo operation

### Array Operations
- `map`: (WIP) Apply logic to each item in an array
- `reduce`: (WIP) Reduce array to single value
- `filter`: (WIP) Select items from array
- `all`, `some`: (WIP) Universal/existential quantifiers
- `merge`:(WIP)  Combine arrays

### String Operations
- `cat`: (WIP) Concatenate strings
- `substr`: (WIP) Get substring
- `in`: (WIP) Check substring/array contains

## Data-Driven Examples

Here are some practical examples showing how to use JsonLogic rules:

```go
// Check if a value is between 0 and 100
rules := map[string]interface{}{
    "and": []interface{}{
        map[string]interface{}{
            ">=": []interface{}{
                map[string]interface{}{"var": "value"},
                0,
            },
        },
        map[string]interface{}{
            "<=": []interface{}{
                map[string]interface{}{"var": "value"},
                100,
            },
        },
    },
}

data := map[string]interface{}{
    "value": 50,
}

// result will be true because "value" >= 0 AND "value" <= 100 
result, err := jsonlogic.Apply(rules, data)
```

## Custom Operations (WIP)

You can extend JsonLogic with custom operations:

```go

```

## Error Handling (WIP)

The library provides detailed error messages for various scenarios:
- Invalid JSON syntax
- Unknown operations
- Type mismatches
- Missing required arguments

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Clone the repository
2. Install dependencies
3. Run tests: `go test -v ./...`

## Acknowledgments

- Based on the [JsonLogic](https://jsonlogic.com) specification
- Inspired by various JsonLogic implementations in other languages
