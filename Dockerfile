# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.57.2@sha256:8f3a60a00a83bb7d599d2e028ac0c3573dc2b9ec0842590f1c2e59781c821da7 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
