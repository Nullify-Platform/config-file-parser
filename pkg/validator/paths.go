package validator

import (
	"github.com/gobwas/glob"
	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func ValidatePaths(config *models.Configuration) []ValidationError {
	errors := []ValidationError{}
	if config.IgnorePaths == nil {
		return errors
	}

	for _, pattern := range config.IgnorePaths {
		_, err := glob.Compile(pattern)
		// log.Printf(">>>>>>>>> pattern: %s, gl: %+v", pattern, gl)
		if err != nil {
			errors = append(errors, ValidationError{
				Field:   "ignore_paths",
				Message: "Invalid paths",
				Line:    config.LocationInfo["ignore_paths"].Line,
				Column:  config.LocationInfo["ignore_paths"].Column,
			})
		}
	}
	return errors
}
