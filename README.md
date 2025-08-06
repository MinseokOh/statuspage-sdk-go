<div align="center">
  <h1>statuspage-sdk-go</h1>
  <p>A robust Go client library for the Statuspage.io API</p>
  <p><strong>✨ Enhanced with Claude Code ✨</strong></p>

  [![Go Documentation](https://pkg.go.dev/badge/github.com/MinseokOh/statuspage-sdk-go.svg)](https://pkg.go.dev/github.com/MinseokOh/statuspage-sdk-go)
  [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
  [![Go Report Card](https://goreportcard.com/badge/github.com/MinseokOh/statuspage-sdk-go)](https://goreportcard.com/report/github.com/MinseokOh/statuspage-sdk-go)
</div>

---

## 📦 Installation

```bash
go get github.com/MinseokOh/statuspage-sdk-go
```

## 🚀 Quick Start

### Authentication

Get your API key from your Statuspage account:
1. Visit [Statuspage Dashboard](https://manage.statuspage.io/login)
2. Navigate to **User Menu** → **API info**

### Basic Usage

```go
package main

import (
    "context"
    "log"
    "os"
    
    statuspage "github.com/MinseokOh/statuspage-sdk-go"
)

func main() {
    apiKey := os.Getenv("STATUSPAGE_API_KEY")
    client := statuspage.NewClient(apiKey)
    
    ctx := context.Background()
    
    // List all pages
    pages, _, err := client.Pages.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, page := range pages {
        log.Printf("Page: %s (%s)", page.Name, page.ID)
    }
}
```
