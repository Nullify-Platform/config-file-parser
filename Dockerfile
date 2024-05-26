# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.59.0@sha256:8ad7dc3d98d77dec753f07408c7683ab854752a3eb8dc6a5e5f0728f9a89ae2c as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
