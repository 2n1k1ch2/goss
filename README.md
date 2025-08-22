# StackScope

<p align="center">
<strong>Catch, cluster, and alert on leaking goroutines in production Go services.</strong>
</p>

<p align="center">
  <a href="#overview">Overview</a> •
  <a href="#how-it-works">How it Works</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#contributing">Contributing</a>
</p>

## Overview

StackScope is a production-ready agent that continuously monitors your Go application for goroutine leaks. Unlike testing utilities, it runs in production, clustering leaks by their origin and alerting you before they become critical.

**Why StackScope?**
- **Clustering:** Groups thousands of goroutines into unique "leak signatures" based on their call stack.
- **Alerting:** Alerts on leak growth over time, correlating spikes with deployments.
- **Production-First:** Designed for zero overhead in live environments.
- **OSS Core:** The core agent and library are open source.

## How It Works

1.  **Collect:** Periodically takes snapshots of all goroutine stack traces.
2.  **Cluster:** Normalizes and hashes stack traces to group them by origin, ignoring variable line numbers.
3.  **Analyze:** Tracks the count of each cluster over time, identifying growing trends.
4.  **Export:** Exposes Prometheus metrics (`stackscope_goroutines_total{signature="<hash>"}`) for alerting and dashboards.
5.  **Alert:** (Cloud Service) Sends alerts when a leak signature's growth rate exceeds a threshold.
## Usage
Integrate the agent into your main.go:

```go
package main

import (
"context"
"github.com/2n1k1ch2/goss"
)

func main() {
// Start the StackScope agent
ctx := context.Background()
agent.Start(ctx, agent.DefaultConfig())

    // ... your application code ...
}
```
Key Prometheus metrics will be available at /metrics:

stackscope_goroutines_total{signature="abc123"}

## Contributing
We welcome contributions! Please feel free to open issues, suggest features, and submit pull requests.