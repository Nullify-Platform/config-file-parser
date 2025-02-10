package validator

import (
	"log"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/require"
)

const configStr = `
severity_threshold: pasdf
`

func TestInvalidSeverityThreshold(t *testing.T) {
	str := configStr
	parsed, parseErr := parser.ParseConfiguration([]byte(str))
	log.Printf("%+v", parsed)
	require.Nil(t, parseErr)

	validationResult := ValidateConfig(parsed)
	log.Println(validationResult)
	require.Equal(t, 1, len(validationResult.Errors))
	require.Equal(t, "severityThreshold", validationResult.Errors[0].Field)
	require.Equal(t, "Invalid severity threshold", validationResult.Errors[0].Message)
	require.Equal(t, 2, validationResult.Errors[0].Line)
	require.Equal(t, 21, validationResult.Errors[0].Column)
}

func TestInvalidPriorityThreshold(t *testing.T) {
	str := `priority_threshold: hello`
	parsed, parseErr := parser.ParseConfiguration([]byte(str))
	log.Printf("%+v", parsed)
	require.Nil(t, parseErr)

	validationResult := ValidateConfig(parsed)
	log.Println(validationResult)
	require.Equal(t, 1, len(validationResult.Errors))
	require.Equal(t, "priorityThreshold", validationResult.Errors[0].Field)
	require.Equal(t, "Invalid priority threshold", validationResult.Errors[0].Message)
	require.Equal(t, 1, validationResult.Errors[0].Line)
	require.Equal(t, 21, validationResult.Errors[0].Column)
}

func TestInvalidNotifications(t *testing.T) {
	str := `
notifications:
  all-events-webhook:
    targets:
      email:
        address: "invalid-email"
`
	parsed, parseErr := parser.ParseConfiguration([]byte(str))
	log.Printf("%+v", parsed)
	require.Nil(t, parseErr)

	validationResult := ValidateConfig(parsed)
	log.Printf("%+v", validationResult)
	require.Equal(t, 1, len(validationResult.Errors))
	require.Equal(t, "notifications.all-events-webhook.targets.email.address", validationResult.Errors[0].Field)
	require.Equal(t, "Invalid notifications", validationResult.Errors[0].Message)
	require.Equal(t, 6, validationResult.Errors[0].Line)
	require.Equal(t, 18, validationResult.Errors[0].Column)
}

func TestInvalidPaths(t *testing.T) {
	str := `
severity_threshold: high
ignore_paths:
  - "invalid-path//[a-"
  - "valid/path"
  - "valid/path/*"
  - "**[!"
`
	parsed, parseErr := parser.ParseConfiguration([]byte(str))
	log.Printf(">>>>>>>>> parsed: %+v <<<", parsed.LocationInfo["ignore_paths"])
	require.Nil(t, parseErr)

	validationResult := ValidateConfig(parsed)
	log.Printf("%+v", validationResult)
	require.Equal(t, 2, len(validationResult.Errors))
	require.Equal(t, "ignore_paths", validationResult.Errors[0].Field)
	require.Equal(t, "Invalid paths", validationResult.Errors[0].Message)
	require.Equal(t, 4, validationResult.Errors[0].Line)
	require.Equal(t, 3, validationResult.Errors[0].Column)
}

func TestInvalidAutoFix(t *testing.T) {
	str := `code:
  auto_fix:
    enabled: true
    max_pull_requests_open: -1
    max_pull_request_creation_rate:
      count: 5
      days: 7

dependencies:
  auto_fix:
    enabled: true
    max_pull_requests_open: 3
    max_pull_request_creation_rate:
      count: -2
      days: 7
`
	parsed, parseErr := parser.ParseConfiguration([]byte(str))
	log.Printf("%+v", parsed)
	require.Nil(t, parseErr)

	validationResult := ValidateConfig(parsed)
	log.Printf("%+v", validationResult)
	require.Equal(t, 2, len(validationResult.Errors))

	require.Equal(t, "code.auto_fix", validationResult.Errors[0].Field)
	require.Equal(t, "Invalid auto fix", validationResult.Errors[0].Message)
	require.Equal(t, 3, validationResult.Errors[0].Line)
	require.Equal(t, 5, validationResult.Errors[0].Column)

	require.Equal(t, "dependencies.auto_fix", validationResult.Errors[1].Field)
	require.Equal(t, "Invalid auto fix", validationResult.Errors[1].Message)
	require.Equal(t, 11, validationResult.Errors[1].Line)
	require.Equal(t, 5, validationResult.Errors[1].Column)
}
