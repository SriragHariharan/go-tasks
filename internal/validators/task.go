package validators

import (
	"errors"

	"github.com/sriraghariharan/gotasks/internal/models"
)

func VerifyNewTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}
	
	return nil
}