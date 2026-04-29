package validators

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/sriraghariharan/gotasks/internal/models"
)

func NewUserValidator(user *models.User) error {
	// Sanitize inputs
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if len(user.Username) > 50 {
		return errors.New("username must be at most 50 characters")
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}