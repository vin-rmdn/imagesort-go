package image

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/vin-rmdn/imagesort-go/tool"
)

type Renamer struct {
	RecognizedImageExtensions []string
}

func (r Renamer) RenameMediaByIndex(oldFile, newFile string) error {
	if !slices.Contains(r.RecognizedImageExtensions, strings.ToLower(filepath.Ext(oldFile))) {
		slog.Info("format not recognized, skipping", slog.String("filename", oldFile))
		
		return nil
	}

	if err := tool.SafeRename(oldFile, newFile); err != nil {
		return fmt.Errorf("failed to safely rename: %w", err)
	}

	return nil
}
