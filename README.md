# Pokedex CLI

A command-line interface application built with Go that lets you explore the Pokemon world, catch Pokemon, and manage your collection - all from your terminal!

## Features

- 🗺️ **Browse Location Areas**: Navigate through different Pokemon locations
- 🔍 **Explore Areas**: Discover Pokemon available in each location
- 🎣 **Catch Pokemon**: Attempt to catch Pokemon with probability-based mechanics
- 📊 **Inspect Pokemon**: View detailed information about caught Pokemon
- 📖 **Pokedex**: View your collection of caught Pokemon
- ⚡ **Smart Caching**: Automatic caching to reduce API calls and improve performance

## Prerequisites

- Go 1.25.1 or higher
- Internet connection (uses the [PokeAPI](https://pokeapi.co/))

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ahmed1bukha/pokedexcli-go.git
cd pokedexcli-go
```

2. Install dependencies:

```bash
go mod download
```

3. Build the application:

```bash
go build
```

4. Run the application:

```bash
./pokedexcli
```

Or run directly with Go:

```bash
go run .
```

## Usage

Once the application starts, you'll see a prompt:

```
Pokedex >
```

### Available Commands

| Command                  | Description                                           | Usage                        |
| ------------------------ | ----------------------------------------------------- | ---------------------------- |
| `help`                   | Display help message with all available commands      | `help`                       |
| `map`                    | Display the next 20 location areas                    | `map`                        |
| `mapb`                   | Display the previous 20 location areas (if available) | `mapb`                       |
| `explore <area-name>`    | Explore a specific location area to find Pokemon      | `explore pastoria-city-area` |
| `catch <pokemon-name>`   | Attempt to catch a Pokemon                            | `catch pikachu`              |
| `inspect <pokemon-name>` | View detailed information about a caught Pokemon      | `inspect pikachu`            |
| `podex`                  | View all Pokemon in your collection                   | `podex`                      |
| `exit`                   | Exit the Pokedex                                      | `exit`                       |

### Example Session

```
Pokedex > help
Welcome to the Pokedex!
Usage:
exit: Exit the Pokedex
help: Displays a help message
map: map the next 20 location
mapb: map the previous 20 location
explore: explore map given name
catch: chance to get pokemon
inspect: inspect poki info in your inventory
podex: get pokimons in inventory

Pokedex > map
canalave-city-area
eterna-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
- tentacool
- tentacruel
- magikarp
...

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
...magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats: ...

Pokedex > podex
Your Pokedex:
- magikarp

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## How Catching Works

The catch probability is calculated based on the Pokemon's `base_experience`:

- Higher base experience = harder to catch
- Lower base experience = easier to catch

The formula used: `catchChance = 100 / (100 + base_experience)`

## Project Structure

```
pokedexcli-go/
├── main.go                 # Main application entry point
├── command_*.go           # Individual command implementations
├── poki_cache/            # Caching package
│   ├── cache.go          # Cache implementation with TTL
│   └── repl_test.go      # Cache tests
├── poki_http/            # HTTP client package
│   ├── http_requests.go  # Generic HTTP request handler
│   └── poki_http.go      # Pokemon API specific types and functions
└── go.mod                # Go module dependencies
```

## Key Components

### Caching System

The application uses a time-based cache (`poki_cache`) that:

- Stores API responses in memory
- Automatically expires entries after a configurable interval
- Uses mutex locks for thread-safe access
- Reduces redundant API calls

### API Integration

The app integrates with the [PokeAPI](https://pokeapi.co/) to fetch:

- Location areas
- Pokemon encounter data
- Pokemon details (stats, height, weight, etc.)

## Development

### Running Tests

Run all tests:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./poki_cache
```

### Code Organization

- **Commands**: Each command has its own file (`command_*.go`)
- **Packages**:
  - `poki_cache`: In-memory caching with TTL
  - `poki_http`: HTTP client and API integration
- **Config**: Shared configuration passed to all commands

## Technologies

- **Go**: Programming language
- **PokeAPI**: Pokemon data source
- **Standard Library**: `net/http`, `encoding/json`, `time`, `sync`

## License

This project is open source and available for educational purposes.

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data
- Built as a learning project for Go development
