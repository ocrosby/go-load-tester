# ADR-0002: Support Kubernetes as a First-Class Execution Environment

## Status  
Accepted

## Context  
We need to support large-scale load testing use cases, and Kubernetes is a standard execution environment for cloud-native workloads. Many users will want to scale out horizontally and run distributed tests.

## Decision  
The load testing tool will be designed to run **natively in Kubernetes**, while still supporting standalone execution outside the cluster.

## Alternatives Considered  
- **Local/CI-only tool**: Limits scalability and parallelism  
- **Custom orchestration framework**: Reinvents the wheel and requires managing distributed infrastructure manually  
- **Serverless architecture (e.g. Lambda)**: Difficult to coordinate and persist test state across distributed runs  

## Consequences  
- K8s manifests and optional Helm chart will be included  
- Each pod can expose status and metrics endpoints  
- Horizontal scaling via Deployments and headless services  
- Coordinator/worker mode supported via CLI roles  

## Author  
Omar Crosby, 2025-06-18
