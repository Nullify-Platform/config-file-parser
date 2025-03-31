# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v2.0.2@sha256:d55581f7797e7a0877a7c3aaa399b01bdc57d2874d6412601a046cc4062cb62e as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
