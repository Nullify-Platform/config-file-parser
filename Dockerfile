# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.58.2@sha256:6d951ed621e2e4c4c7de1b284341b5eb62911b67f64d9ef83d78966fdff5b022 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
