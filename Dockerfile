# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.58.0@sha256:42692fa92a34b74d9e85687b5efea48f0c5eaa13f93d08edd72fb629d9385933 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
