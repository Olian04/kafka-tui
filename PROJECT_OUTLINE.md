# Kafka TUI - Project Outline

## Overview

Kafka TUI is a terminal user interface (TUI) application for interacting with Apache Kafka as both a consumer and producer. The application provides an intuitive, tree-based navigation system for exploring Kafka clusters, topics, consumer groups, and managing message consumption and production.

## Technology Stack

- **Language**: Go (Golang)
- **TUI Framework**: [bubbletea](https://github.com/charmbracelet/bubbletea) - For rendering the terminal user interface
- **CLI Framework**: [urfave/cli](https://github.com/urfave/cli) - For managing CLI configuration and flags
- **Kafka Client**: [kafka-go](https://github.com/segmentio/kafka-go) - For interacting with Kafka brokers

## Project Structure

The project follows a clear separation between core logic and application entry points:

```
kafka-tui/
├── cmd/
│   └── kafka-tui/
│       └── main.go          # Entry point and TUI rendering code
├── pkg/
│   ├── kafka/               # Core Kafka interaction logic
│   │   ├── client.go        # Kafka client wrapper
│   │   ├── consumer.go      # Consumer operations
│   │   ├── producer.go      # Producer operations
│   │   └── metadata.go      # Metadata operations (topics, brokers, etc.)
│   └── ...                  # Other core packages as needed
└── go.mod
```

### Directory Responsibilities

- **`pkg/`**: Contains reusable, testable core business logic with no UI dependencies
  - `pkg/kafka/`: Handles all Kafka protocol interactions, connection management, and data operations
  
- **`cmd/kafka-tui/`**: Contains application-specific code
  - Entry point (`main.go`)
  - TUI rendering logic using bubbletea
  - User input handling
  - Navigation state management

## Core Features

### 1. Topic Management
- List all available topics in the cluster
- View topic details (partitions, replication factor, configuration)
- Navigate topics in a tree structure

### 2. Consumer Group Operations
- List consumer groups for a topic
- View consumer group details and offsets
- Monitor consumer group lag

### 3. Message Consumption
- Listen to messages on a topic (real-time)
- View message content (key, value, headers, timestamp, offset)
- Filter messages by partition
- Navigate through message history

### 4. Message Production
- Send messages to a topic
- Specify message key, value, and headers
- Choose partition (or let Kafka decide)
- Batch message sending

### 5. Cluster Information
- View cluster configuration
- List all brokers in the cluster
- View broker details and status
- List topics and partitions per broker

### 6. Configuration Viewing
- View cluster-level configuration
- View topic-level configuration
- View broker-level configuration

## User Flow & Navigation

The application uses a **tree-based navigation model** similar to file explorers or resource browsers:

### Navigation Pattern
1. **Root View**: Overview of cluster resources (topics, brokers, consumer groups)
2. **Resource Selection**: Use arrow keys to navigate lists, Enter to drill down
3. **Detail Views**: Press Enter on a resource to see details and available actions
4. **Action Views**: Select actions (e.g., "Listen to messages", "Send message") to perform operations
5. **Back Navigation**: Use Esc or Backspace to navigate up the tree

### Example Flow
```
Cluster Overview
  ├─ Topics
  │   ├─ topic-orders (Enter)
  │   │   ├─ View Details
  │   │   ├─ Listen to Messages (Enter)
  │   │   │   └─ [Real-time message stream view]
  │   │   ├─ Send Message (Enter)
  │   │   │   └─ [Message composition form]
  │   │   └─ View Partitions
  │   └─ topic-users
  ├─ Brokers
  │   └─ [Broker list with details]
  └─ Consumer Groups
      └─ [Consumer group list]
```

### Key Bindings (Tentative)
- **Arrow Keys**: Navigate lists/options
- **Enter**: Select/Drill down
- **Esc/Backspace**: Go back/Up one level
- **Tab**: Switch between panes/views
- **q**: Quit application
- **/**: Search/Filter
- **?**: Help

## Visual Design

The visual design should be inspired by [k9s](https://github.com/derailed/k9s), which provides a clean, information-dense terminal interface.

### Design Principles
- **Multi-pane layout**: Split screen showing resource list and details simultaneously
- **Color coding**: Use colors to indicate status (active, error, warning)
- **Status bar**: Bottom bar showing current context, shortcuts, and status
- **Header**: Top bar showing cluster connection info and current view
- **Clean typography**: Clear, readable fonts with appropriate spacing
- **Real-time updates**: Live updates for message streams and cluster status

### Layout Structure
```
┌─────────────────────────────────────────────────┐
│ Kafka TUI v1.0    Cluster: localhost:9092      │ ← Header
├─────────────────────────────────────────────────┤
│                                                  │
│  Topics                    Topic Details         │
│  ┌──────────────┐         ┌──────────────────┐ │
│  │ • topic-1    │         │ Name: topic-1    │ │
│  │   topic-2    │         │ Partitions: 3    │ │
│  │   topic-3    │         │ Replication: 2   │ │
│  └──────────────┘         └──────────────────┘ │
│                                                  │
├─────────────────────────────────────────────────┤
│ [Enter]Select [Esc]Back [q]Quit [/?]Help       │ ← Status Bar
└─────────────────────────────────────────────────┘
```

## Implementation Phases

### Phase 1: Foundation
- [x] Project structure setup
- [x] Basic CLI setup with urfave/cli
- [x] Kafka client initialization
- [x] Basic bubbletea TUI skeleton

### Phase 2: Core Navigation
- [ ] Tree navigation system
- [ ] Resource list views (topics, brokers)
- [ ] Detail view rendering
- [ ] Navigation state management

### Phase 3: Kafka Operations
- [ ] List topics
- [ ] List brokers
- [ ] View topic details
- [ ] View broker details

### Phase 4: Consumer Features
- [ ] List consumer groups
- [ ] Listen to topic messages
- [ ] Message display formatting
- [ ] Real-time message streaming

### Phase 5: Producer Features
- [ ] Message composition UI
- [ ] Send message functionality
- [ ] Batch sending

### Phase 6: Configuration & Polish
- [ ] Configuration viewing
- [ ] Error handling and display
- [ ] Help system
- [ ] Visual polish and styling

## Configuration

The application should support configuration via:
- Command-line flags (via urfave/cli)
- Environment variables
- Configuration file (optional, future enhancement)

### Required Configuration
- Kafka broker addresses
- Connection timeout settings
- TLS/SASL settings (if needed)

## Error Handling

- Display errors in a user-friendly way within the TUI
- Provide context about what operation failed
- Allow users to retry failed operations
- Log errors for debugging (to file, not stdout)

## Testing Strategy

- Unit tests for `pkg/kafka/` logic
- Integration tests with a test Kafka cluster
- TUI component tests where feasible
- Manual testing for user experience validation

## Future Enhancements (Out of Scope for Initial Version)

- Message filtering and search
- Message replay from specific offsets
- Consumer group management operations
- Topic creation/deletion
- Configuration editing
- Multiple cluster connections
- Saved queries/bookmarks
- Export messages to file
