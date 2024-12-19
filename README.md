# P2P Flooding Node in Go

A simple peer-to-peer (P2P) flooding system implemented in Go. Each node listens for messages and floods any received message to all its connected peers. When a node receives a message, it prints it out and ensures that each message is processed only once to prevent infinite loops.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Running Nodes](#running-nodes)
  - [Sending Messages](#sending-messages)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Features

- **P2P Communication**: Nodes communicate directly with each other using UDP.
- **Message Flooding**: Any message received by a node is forwarded to all its peers.
- **Duplicate Prevention**: Each message is processed only once to avoid infinite loops.
- **Simple and Lightweight**: Easy to set up and understand, suitable for learning and small-scale applications.

## Prerequisites

- **Go**: Ensure you have Go installed (version 1.16 or later). You can download it from [golang.org](https://golang.org/dl/).

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/p2p-flooding-node.git
   cd p2p-flooding-node
2. **Initialize Go Modules**:
Ensure that Go modules are enabled (Go 1.16+ has them enabled by default).
   ```bash
   go mod tidy

## Project Structure
  ```bash
  p2p-flooding-node/
  ├─ cmd/
  │  └─ node/
  │      └─ main.go            # Entry point of the application
  ├─ internal/
  │  └─ node/
  │      ├─ models.go          # Data structures and message formats
  │      ├─ node.go            # Node orchestration logic
  │      ├─ repository.go      # In-memory repository for peers and messages
  │      ├─ service.go         # Core business logic for message handling
  │      └─ transport.go       # UDP networking (listening and sending)
  ├─ go.mod
  └─ go.sum```
