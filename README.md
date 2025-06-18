Go Load Tester

A scalable, configurable, and Kubernetes-ready load testing tool for RESTful APIs written in Go.

Features
	•	High-performance HTTP load testing using Go’s native concurrency
	•	YAML-based configurable scenarios and routes
	•	Supports both standalone and Kubernetes deployment modes
	•	Live status updates via /status HTTP endpoint
	•	Prometheus-compatible /metrics endpoint
	•	Weighted scenario execution
	•	Role-based CLI support (standalone, worker, coordinator)

Use Cases
	•	Load testing RESTful APIs locally or in CI pipelines
	•	Distributed performance testing at scale in Kubernetes
	•	Continuous monitoring or regression testing under load

Project Structure

go-load-tester/
├── cmd/
│   └── loadtester/             # Main CLI entry point
├── internal/
│   ├── attack/                 # Core load generation logic
│   ├── config/                 # CLI + YAML config loaders
│   ├── stats/                  # Result aggregation and reporting
│   └── coordinator/           # (Planned) distributed coordination logic
├── deploy/k8s/                # Kubernetes manifests (coming soon)
├── go.mod
└── README.md

Sample loadplan.yaml

base_url: https://api.example.com
defaults:
  method: GET
  headers:
    Authorization: Bearer dummy-token
    Content-Type: application/json
  timeout: 5s
  concurrency: 10
  requests: 100

scenarios:
  - name: Get users
    path: /v1/users
    method: GET
    weight: 2

  - name: Create user
    path: /v1/users
    method: POST
    body: |
      {
        "name": "John",
        "email": "john@example.com"
      }
    weight: 1

Usage

Standalone Mode

loadtester --plan loadplan.yaml

Kubernetes Worker

loadtester --plan loadplan.yaml --role worker --coordinator http://coordinator-svc:8080

Kubernetes Coordinator

loadtester --role coordinator

Status and Metrics
	•	GET /status returns live test progress as JSON
	•	GET /metrics returns Prometheus-formatted metrics

Planned Enhancements
	•	Coordinator/worker sync for distributed load tests
	•	HTMX-based dashboard UI
	•	gRPC or REST-based result aggregation
	•	Helm chart for Kubernetes deployment

License

MIT License
