package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vin-rmdn/imagesort-go/config"
	image "github.com/vin-rmdn/imagesort-go/media"
)

func RenameCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "rename",
		Short:   "Renames supported medias according to ascending index",
		Long:    "Renames supported medias according to ascending index",
		Example: "./imagesort_go rename /path/to/destination/folder",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			renamer := image.Renamer{
				RecognizedImageExtensions: config.RecognizedImageExtensions,
			}

			folderPath := args[0]

			dirEntries, err := os.ReadDir(folderPath)
			if err != nil {
				return fmt.Errorf("error on reading dir %s: %w", folderPath, err)
			}

			index := 1

			for _, dirEntry := range dirEntries {
				if dirEntry.IsDir() {
					continue
				}

				_, err := dirEntry.Info()
				if err != nil {
					return fmt.Errorf("error in getting file info: %w", err)
				}

				file := dirEntry.Name()
				oldFile := fmt.Sprintf("%s%c%s", folderPath, os.PathSeparator, file)
				for {
					newFile := fmt.Sprintf("%s%c%d%s", folderPath, os.PathSeparator, index, filepath.Ext(file))
					if err := renamer.RenameMediaByIndex(oldFile, newFile); err != nil {
						slog.Warn("rename failed, increasing index", slog.String("newFile", newFile))

						index++

						continue
					}

					break
				}

				index++
			}

			return nil
		},
	}
}
