# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v1.63.4@sha256:7f4c8ee8a63d56caa41c099cf658f68b192b615e0f30e94b8864e81a3ceafb53 as golangci-lint
FROM hadolint/hadolint:2.12.0@sha256:30a8fd2e785ab6176eed53f74769e04f125afb2f74a6c52aef7d463583b6d45e as hadolint
