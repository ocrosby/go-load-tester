ADR-0003: Use YAML for Load Plan Configuration

Status

Accepted

Context

We need a way for users to define a flexible, readable set of API scenarios and load parameters. This configuration must be user-friendly, versionable, and integratable with CI/CD workflows.

Decision

We will use YAML files as the primary format for load plan configuration.

Alternatives Considered
	•	JSON: Machine-friendly but less human-friendly and lacks comments
	•	TOML: Cleaner than JSON, but less familiar and less widely adopted
	•	HCL: Designed for infrastructure, but not a strong fit for API load scenarios

Consequences
	•	Human-readable, commentable configuration
	•	Simple to parse using gopkg.in/yaml.v3
	•	Easy to extend with defaults and scenario weighting
	•	YAML is well-supported in CI environments and DevOps tooling

Author

Omar Crosby, 2025-06-18
