package validator

import (
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func TestValidateEmail(t *testing.T) {
	// Test case 1: Valid email addresses
	config := &models.Configuration{
		Notifications: map[string]models.Notification{
			"notification1": {
				Targets: models.NotificationTargets{
					Email: &models.NotificationTargetEmail{
						Addresses: []string{"user1@example.com", "user2@example.com"},
					},
				},
			},
		},
	}

	if !ValidateEmail(config) {
		t.Errorf("Expected email addresses to be valid, but got invalid.")
	}

	// Test case 2: Invalid email address
	config = &models.Configuration{
		Notifications: map[string]models.Notification{
			"notification1": {
				Targets: models.NotificationTargets{
					Email: &models.NotificationTargetEmail{
						Addresses: []string{"user1@example.com", "invalid-email"},
					},
				},
			},
		},
	}

	if ValidateEmail(config) {
		t.Errorf("Expected email addresses to be invalid, but got valid.")
	}

	// Test case 3: No email addresses to validate
	config = &models.Configuration{
		Notifications: map[string]models.Notification{
			"notification1": {
				Targets: models.NotificationTargets{
					Email: nil,
				},
			},
		},
	}

	if !ValidateEmail(config) {
		t.Errorf("Expected email addresses to be valid when none are provided, but got invalid.")
	}
}
