# Kafka TUI

A terminal user interface (TUI) application for interacting with Apache Kafka as both a consumer and producer.

## Features

- List available topics
- List consumer groups for a topic
- Listen to messages sent on a topic
- Send messages on a topic
- View configuration for the cluster, topics, etc.
- List brokers on the cluster
- List topics and partitions per broker

## Installation

### Build from Source

```bash
go build -o kafka-tui ./cmd/kafka-tui
```

### Install

```bash
go install ./cmd/kafka-tui@latest
```

## Usage

```bash
kafka-tui --brokers localhost:9092
```

Or with multiple brokers:

```bash
kafka-tui --brokers localhost:9092,localhost:9093,localhost:9094
```

### Command Line Options

- `--brokers, -b`: Kafka broker addresses (required, can be set via `KAFKA_BROKERS` env var)
- `--timeout`: Connection timeout (default: 10s, can be set via `KAFKA_TIMEOUT` env var)
- `--help, -h`: Show help
- `--version, -v`: Show version

### Environment Variables

- `KAFKA_BROKERS`: Comma-separated list of broker addresses
- `KAFKA_TIMEOUT`: Connection timeout (e.g., "10s", "30s")

## Development

### Project Structure

```
kafka-tui/
├── cmd/
│   └── kafka-tui/        # Application entry point
│       └── main.go
├── pkg/
│   └── kafka/            # Core Kafka interaction logic
│       └── client.go
└── go.mod
```

### Building

```bash
go build -o kafka-tui ./cmd/kafka-tui
```

### Running Tests

```bash
go test ./...
```

## Status

**Phase 1: Foundation** ✅ Complete
- Project structure setup
- Basic CLI setup with urfave/cli
- Kafka client initialization
- Basic bubbletea TUI skeleton

See [PROJECT_OUTLINE.md](PROJECT_OUTLINE.md) for full project details and roadmap.

## License

See [LICENSE](LICENSE) file for details.
