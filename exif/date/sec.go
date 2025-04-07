package date

type second struct{}

func (second) ExifDateFormat() string {
	return "2006:01:02 15:04:05"
}

func (second) FileFormatWithoutExtension() string {
	return "2006-01-02 15.04.05"
}

type DateTimeOriginal struct {
	second
}

func (DateTimeOriginal) FieldName() string {
	return "DateTimeOriginal"
}

type CreateDate struct {
	second
}

func (CreateDate) FieldName() string {
	return "CreateDate"
}

type CreationDate struct {
	second
}

func (CreationDate) FieldName() string {
	return "CreationDate"
}

type secondWithTimezone struct{}

func (secondWithTimezone) ExifDateFormat() string {
	return "2006:01:02 15:04:05Z07:00"
}

func (secondWithTimezone) FileFormatWithoutExtension() string {
	return "2006-01-02 15.04.05 -0700"
}

type CreationDateWithTimezone struct {
	secondWithTimezone
}

func (CreationDateWithTimezone) FieldName() string {
	return "CreationDate"
}

type FileModifyDateWithTimezone struct {
	secondWithTimezone
}

func (FileModifyDateWithTimezone) FieldName() string {
	return "FileModifyDate"
}

type DateTimeOriginalWithTimezone struct {
	secondWithTimezone
}

func (DateTimeOriginalWithTimezone) FieldName() string {
	return "DateTimeOriginal"
}

