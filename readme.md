# Understat Statistics Module

The Understat Statistics Module is a Go library designed to fetch and process football data from the [Understat](https://understat.com) website. It provides convenient APIs to retrieve player, game, and team statistics for specific leagues and years. This module is perfect for developers and analysts looking to integrate Understat data into their projects.

---

## Features

- **Player Statistics**: Retrieve detailed data about players.
- **Game Statistics**: Fetch game results and related statistics.
- **Team Statistics**: Access team performance data.
- Easy-to-use APIs for interacting with Understat data.
- Comprehensive test suite for robust functionality.

---

## Prerequisites

Ensure you have the following installed before using this module:

- [Go](https://golang.org/dl/) (version 1.18 or later)

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/gkonto/understat.git
   cd understat

   ```

2. Download and resolve dependencies

```bash
   go mod tidy
```

---

## Usage

Example: Fetching Football Data

1. Import the required packages:

```go
import (
    "github.com/gkonto/understat"
    "github.com/gkonto/understat/model"
)
```

2. Create an instance of UnderstatAPI:

```go
api := understat.NewUnderstatAPI()
```

3. Fetch and display data:

- Players:

```go
players, err := api.GetPlayers(model.League("EPL"), model.Year(2023))
if err != nil {
    fmt.Println("Error fetching players:", err)
    return
}
fmt.Println("Players:", players)

```

- Games:

```go
games, err := api.GetGames(model.League("EPL"), model.Year(2023))
if err != nil {
    fmt.Println("Error fetching games:", err)
    return
}
fmt.Println("Games:", games)
```

- Teams:

```go
teams, err := api.GetTeams(model.League("EPL"), model.Year(2023))
if err != nil {
    fmt.Println("Error fetching teams:", err)
    return
}
fmt.Println("Teams:", teams)
```

---

## Running Tests

To validate the functionality of the module, you can run the included test suite:

1. Run all tests:

```bash
go test ./...
```

2. For verbose output:

```bash
go test ./... -v
```

3. Run specific test files:

```bash
go test ./internal/controller -v
```

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

- Inspired by the comprehensive data on Understat.
- Developed with Go for performance and simplicity.
