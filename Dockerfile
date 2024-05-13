# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.58.1@sha256:5bef7ef61a4e2529b39d4e39de3564d82c38291636cdb9b79a656cedb09ab175 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
