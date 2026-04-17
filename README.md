# 👻 Ghost-Scraper-API

A high-performance, concurrent web scraping engine built in Go. This project features a custom-built **Least-Connections Load Balancer** to manage a fleet of "Ghost Workers" (browser instances) efficiently.

## 🚀 Overview

Ghost-Scraper-API is designed for stability and speed. Instead of spawning uncontrolled processes, it uses a central Load Balancer to track worker availability and distribute scraping tasks to the most idle instance.

## 🏗️ Architecture

- **HTTP API Layer**: Handles incoming JSON requests and validates input.
- **Load Balancer**: A thread-safe manager using `sync.Mutex` to coordinate worker selection.
- **Ghost Workers**: Individual units representing browser resources, tracked via memory-efficient **pointers**.
- **Scraping Engine**: The core logic that interacts with target websites using specialized Go libraries.

## 🛠️ Key Concepts Applied

- **Concurrency Control**: Prevents race conditions during high traffic using Mutex locking.
- **Memory Optimization**: Leverages Go pointers (`*`) to share worker states across the application without expensive data copying.
- **Least-Connections Algorithm**: Intelligently routes traffic to the worker with the fewest active tasks.

## 📋 Prerequisites

- [Go](https://golang.org/doc/install) (1.18 or higher)
- A working Go workspace

## 📥 Installation

```bash
git clone [https://github.com/your-username/Ghost-Scraper-API.git](https://github.com/your-username/Ghost-Scraper-API.git)
cd Ghost-Scraper-API
go mod tidy

```
## 🚦 Usage
1. **Start the server:**
   ```bash
   go run main.go
   ```
2. **Send a Scrape Request:**
   ```bash
   curl -X POST http://localhost:8080/Scrape-Data \\
   ```
## 🔐 Security & Reliability
This project implements core networking and security principles similar to those found in professional Master's level curriculum:
 - Thread Safety: Guaranteed by atomic-like operations on shared resources.
 - Resource Pooling: Limits browser overhead to prevent system crashes.
## 📜 License
Distributed under the MIT License.
