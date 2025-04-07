package exif_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vin-rmdn/imagesort-go/exif"
	"github.com/vin-rmdn/imagesort-go/exif/date"
)

func TestRenameImagesWithCreationDate(t *testing.T) {
	tool, err := exif.New([]exif.DateParser{
		date.SubSecondDateTimeOriginal{},
	}, []string{".jpg"})
	require.NoError(t, err)

	t.Cleanup(func() {
		if _, err := os.Stat("testdata/2024-09-13 12.21.29.937.jpeg"); err == nil {
			os.Rename("testdata/2024-09-13 12.21.29.937.jpeg", "testdata/sample.jpg")	
		}
	})

	t.Run("jpg", func(t *testing.T) {
		if _, err := os.Stat("testdata/2024-09-13 12.21.29.937.jpeg"); err == nil {
			os.Rename("testdata/2024-09-13 12.21.29.937.jpeg", "testdata/sample.jpg")	
		}

		err := tool.RenameImagesWithCreationDate(context.Background(), []string{"testdata/sample.jpg"})
		require.NoError(t, err)
	})
}
