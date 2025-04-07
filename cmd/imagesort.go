package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/vin-rmdn/imagesort-go/config"
	"github.com/vin-rmdn/imagesort-go/exif"
	"github.com/vin-rmdn/imagesort-go/exif/date"
)


func ImageSortCommand() (*cobra.Command, error) {
	exifTool, err := exif.New([]exif.DateParser{
		date.SubSecondDateTimeOriginalWithTimezone{}, // iPhones, Fujifilm
		date.SubSecondDateTimeOriginal{}, // GoPro
		date.DateTimeOriginalWithTimezone{}, // old handycam MTS file
		date.DateTimeOriginal{},
		date.CreationDateWithTimezone{}, // Preferred for videos
		date.CreationDate{},
		date.SubSecCreateDateWithTimezone{},
		date.CreateDate{},
		date.FileModifyDateWithTimezone{},
	}, config.RecognizedImageExtensions)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize exif tool: %w", err)
	}

	return &cobra.Command{
		Use:     "sort-media",
		Aliases: []string{"sortimage", "sort-image"},
		Short:   "Sorts images and renames according to EXIF date",
		Long:    "Sorts images and renames according to EXIF date (long explanation)",
		Example: "./imagesort_go sort-image /path/to/destination/folder",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("expected 1st argument, found %s instead", args[0])
			}

			folderPath := args[0]

			dirEntries, err := os.ReadDir(folderPath)
			if err != nil {
				return fmt.Errorf("error on reading dir %s: %w", folderPath, err)
			}

			for _, dirEntry := range dirEntries {
				if dirEntry.IsDir() {
					continue
				}

				fileInfo, err := dirEntry.Info()
				if err != nil {
					return fmt.Errorf("error in getting file info: %w", err)
				}

				slog.Debug("file information", slog.Any("info", fileInfo))

				exifTool.RenameImagesWithCreationDate(
					context.Background(),
					[]string{fmt.Sprintf("%s/%s", folderPath, dirEntry.Name())},
				)
			}

			return nil
		},
		SilenceErrors:              false,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 5,
	}, nil
}
