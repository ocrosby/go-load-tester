# ADR-0001: Use Go as the Implementation Language

## Status  
Accepted

## Context  
We need to choose a primary language for implementing a general-purpose load testing tool that is performant, scalable, and portable across environments (including Kubernetes and standalone).

## Decision  
We will implement the load testing tool in **Go (Golang)**.

## Alternatives Considered  
- **Python**: Easy to prototype, but not ideal for high-concurrency or resource-efficient workloads  
- **Node.js**: Good for I/O-heavy tasks, but more complex to distribute and harder to manage in low-level system tasks  
- **Rust**: Offers performance, but steep learning curve and slower iteration  

## Consequences  
- High concurrency and speed using goroutines  
- Easy distribution via static binaries  
- Native Kubernetes tooling and cloud compatibility  
- Simplified HTTP handling using Go’s standard library  

## Author  
Omar Crosby, 2025-06-18
