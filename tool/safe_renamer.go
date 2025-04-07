package tool

import (
	"fmt"
	"os"
)

func SafeRename(oldName, newName string) error {
	if _, statErr := os.Stat(newName); statErr == nil {
		return os.ErrExist
	}

	if err := os.Rename(oldName, newName); err != nil {
		return fmt.Errorf("cannot rename %s to %s: %w", oldName, newName, err)
	}

	return nil
}