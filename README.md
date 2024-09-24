# CDNCheck

**CDNCheck** is a Go library that helps you detect whether an IP address belongs to a Content Delivery Network (CDN), Web Application Firewall (WAF), or Cloud provider. The library uses custom DNS resolvers with retry logic to ensure accurate and robust detection.

## Features

- Detects if an IP belongs to a CDN, WAF, or Cloud provider.
- Supports custom DNS resolvers with retry capabilities.
- Easy-to-use API for integration in other tools.
- Built-in trusted DNS resolver list (Cloudflare, Google DNS).

## Installation

To install the package, you can use `go get`:

```bash
go get github.com/tongchengbin/cdncheck
```

## Usage
Here’s an example of how to use CDNCheck in your project:

```go
package main
import (
    "fmt"
    "net"
    "github.com/tongchengbin/cdncheck"
)

func main() {
    // Initialize the client (runs once)
    cdncheck.Init()
    
    // Test an IP address
    ip := net.ParseIP("8.8.8.8")
    matched, value, itemType, err := cdncheck.CdnClient.Check(ip)
    
    if err != nil {
        fmt.Printf("Error checking IP: %v\n", err)
    } else if matched {
        fmt.Printf("IP %v belongs to %v (%v)\n", ip, value, itemType)
    } else {
        fmt.Printf("IP %v does not match any known provider\n", ip)
    }
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.