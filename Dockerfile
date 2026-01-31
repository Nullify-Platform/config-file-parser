# nosec The image is used solely for linting purposes and does not run as a service
FROM golangci/golangci-lint:v2.8.0@sha256:bebcfa63db7df53e417845ed61e4540519cf74fcba22793cdd174b3415a9e4e2 as golangci-lint
FROM hadolint/hadolint:v2.14.0@sha256:27086352fd5e1907ea2b934eb1023f217c5ae087992eb59fde121dce9c9ff21e as hadolint
