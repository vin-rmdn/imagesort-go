package exif

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/barasher/go-exiftool"
)

type DateParser interface {
	ExifDateFormat() string
	FileFormatWithoutExtension() string
	FieldName() string
}

type Exif struct {
	tool                      *exiftool.Exiftool
	recognizedDates           []DateParser
	recognizedImageExtensions []string
}

func New(recognizedDates []DateParser, recognizedImageExtensions []string) (Exif, error) {
	tool, err := exiftool.NewExiftool()
	if err != nil {
		return Exif{}, fmt.Errorf("error in creating exif: %w", err)
	}

	return Exif{
		tool:                      tool,
		recognizedDates:           recognizedDates,
		recognizedImageExtensions: recognizedImageExtensions,
	}, nil
}

func (e Exif) RenameImagesWithCreationDate(ctx context.Context, files []string) error {
	for _, file := range files {
		if !slices.Contains(e.recognizedImageExtensions, strings.ToLower(filepath.Ext(file))) {
			slog.Info("format not recognized, skipping", slog.String("filename", file))
			continue
		}

		slog.Debug("processing image", slog.String("filename", file))

		imagesMetadata := e.tool.ExtractMetadata(file)
		imageMetadata := imagesMetadata[0]
		if imageMetadata.Err != nil {
			slog.Error("metadata error", slog.String("error", imageMetadata.Err.Error()))

			return imageMetadata.Err
		}

		creationDate, outputDateFormat, err := e.creationTimeFromExif(imageMetadata.Fields)
		if err != nil {
			slog.Error(
				"creation time extraction error",
				slog.String("error", err.Error()),
				slog.String("filename", file),
			)

			return fmt.Errorf("creation time extraction error: %w", err)
		}

		fileExtension := filepath.Ext(file)
		headerExtension, ok := imageMetadata.Fields["FileType"]
		if ok {
			fileExtension = fmt.Sprintf(".%s", headerExtension.(string))
			fileExtension = strings.ToLower(fileExtension)
		}

		newFile := fmt.Sprintf(
			"%s%c%s%s",
			filepath.Dir(file),
			os.PathSeparator,
			creationDate.Format(outputDateFormat),
			fileExtension,
		)

		if _, statErr := os.Stat(newFile); statErr == nil {
			slog.Warn("file exists, skipping", slog.String("new_filename", newFile))

			continue
		}

		slog.Debug("new name", slog.String("new_filename", newFile))

		if err := os.Rename(file, newFile); err != nil {
			slog.Error(
				"rename failed",
				slog.String("old_name", file),
				slog.String("new_name", newFile),
				slog.String("error", err.Error()),
			)

			break
		}
	}

	return nil
}

func (e Exif) creationTimeFromExif(exifData map[string]any) (time.Time, string, error) {
	for _, recognizedDate := range e.recognizedDates {
		slog.Debug(
			"fields",
			slog.String("field_name", recognizedDate.FieldName()),
			slog.String("ExifDateFormat", recognizedDate.ExifDateFormat()),
		)

		if val, ok := exifData[recognizedDate.FieldName()]; ok {
			slog.Debug("exif date format found", slog.String("field_name", recognizedDate.FieldName()))
			date, err := time.Parse(recognizedDate.ExifDateFormat(), val.(string))
			if err != nil {
				slog.Warn("can't parse exif date format", slog.String("field_name", recognizedDate.FieldName()))
				// return time.Time{}, fmt.Errorf("cannot parse time: %w", err)
				continue
			}

			return date, recognizedDate.FileFormatWithoutExtension(), nil
		}
	}

	return time.Time{}, "", errors.New("empty datetime")
}
