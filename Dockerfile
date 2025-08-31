# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v2.4.0@sha256:91460846c43b3de53eb77e968b17363e8747e6f3fc190575b52be60c49446e23 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
