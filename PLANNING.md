Load Tester Project Planning Document

Step 1: Start from Your Goal

🎯 Goal

To build a high-performance, horizontally scalable load testing tool for RESTful APIs, suitable for both Kubernetes-native environments and local execution.

👤 Target Users
	•	DevOps engineers testing service durability under load
	•	QA professionals validating API performance regressions
	•	Backend engineers profiling new services before production

💡 Value Proposition
	•	Scalable across Kubernetes for distributed load testing
	•	Fully configurable via human-readable YAML files
	•	Clear observability with real-time status and Prometheus metrics
	•	Operable in both cloud-native and developer-local environments
	•	Open source, with readable, idiomatic Go code

Step 2: Write User Stories
	1.	The user should be able to define test scenarios in a YAML file
	2.	The user should be able to run tests locally using a single command
	3.	The user should be able to deploy the tool to a Kubernetes cluster
	4.	The user should be able to run it in standalone or distributed mode
	5.	The user should be able to view real-time status via HTTP
	6.	The user should be able to scrape Prometheus metrics from each pod
	7.	The user should be able to receive test results via CLI or JSON
	8.	The user should be able to weight scenarios proportionally
	9.	The user should be able to define headers, methods, timeouts
	10.	The user should be able to see success/failure and latency metrics
	11.	The user should be able to run the coordinator independently
	12.	The user should be able to aggregate results in distributed mode
	13.	The user should be able to see output via stdout or a /results endpoint
	14.	The user should be able to set configuration via CLI flags or ENV
	15.	The user should be able to test multiple routes under a single plan
	16.	The user should be able to run it in CI pipelines for regression checks
	17.	The user should be able to extend the tool with plugins/hooks later

Step 3: Define Your Data Models

LoadPlan
	•	base_url: string
	•	defaults: struct
	•	method, headers, timeout, concurrency, requests
	•	scenarios: []Scenario

Scenario
	•	name: string
	•	path: string
	•	method: string (optional)
	•	headers: map[string]string (optional)
	•	body: string (optional)
	•	weight: int

Result
	•	status_code: int
	•	duration: time.Duration
	•	error: error (optional)

Metrics (in-memory or Prometheus)
	•	success_count: int
	•	failure_count: int
	•	avg_latency: float
	•	per-scenario metrics (map[string]ScenarioMetrics)

Step 4: Minimum Viable Product (MVP)
	•	Single binary
	•	CLI flag for --plan loadplan.yaml
	•	Execute requests concurrently using goroutines
	•	Parse YAML config with defaults/scenarios
	•	Simple result output to stdout
	•	/status endpoint with in-memory results
	•	Optional /metrics endpoint with Prometheus format
	•	Role support: --role standalone (default), worker, coordinator

Step 5: Simple Wireframe

                +---------------------+
                |   Coordinator Pod   |
                |   (optional mode)   |
                +----------+----------+
                           |
          +-----------------------------+
          |                             |
+---------v---------+       +-----------v--------+
| Load Tester Pod 1 | ...   | Load Tester Pod N  |
+-------------------+       +--------------------+
         |                             |
         +-----------> Target API <----+

Web endpoints:
	•	/status — JSON health check/status output
	•	/metrics — Prometheus format

Step 6: Future-Proofing
	•	Supports K8s but doesn’t depend on it (runs locally too)
	•	CLI and API can support plugins or future features
	•	Prometheus-ready from day one
	•	Possible future integration with Jaeger/OpenTelemetry
	•	Helm chart planned for quick cluster installs

Step 7: Drill Into Specific Components
	•	YAML Config Loader: validate + default
	•	Scenario Scheduler: weight-based distribution
	•	Runner Engine: goroutines with result channel
	•	Metrics Collector: internal + Prometheus exporter
	•	Status Server: exposes /status and /metrics
	•	Coordinator: (later) assigns load to registered workers
	•	Worker: executes load plans from coordinator

Step 8: Stack Selection
	•	Language: Go (Golang)
	•	CLI: Native flag package
	•	Config: YAML via gopkg.in/yaml.v3
	•	HTTP Client: Go’s net/http
	•	Metrics: Prometheus Go client
	•	Kubernetes: native deployments + headless service
	•	Distribution: Docker + Helm (later)

Step 9: Development Process

Skeleton Setup
	•	Create project layout and go.mod
	•	Define cmd/loadtester/main.go
	•	Set up internal packages: config, attack, stats

Data Model
	•	Define LoadPlan, Scenario, and Result

Backend Logic
	•	Load YAML config
