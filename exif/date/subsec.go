package date

type subSecondWithTimezone struct{}

func (subSecondWithTimezone) ExifDateFormat() string {
	return "2006:01:02 15:04:05.99999Z07:00"
}

func (subSecondWithTimezone) FileFormatWithoutExtension() string {
	return "2006-01-02 15.04.05.99999 -0700"
}

type SubSecondDateTimeOriginalWithTimezone struct {
	subSecondWithTimezone
}

func (SubSecondDateTimeOriginalWithTimezone) FieldName() string {
	return "SubSecDateTimeOriginal"
}

type SubSecCreateDateWithTimezone struct {
	subSecondWithTimezone
}

func (SubSecCreateDateWithTimezone) FieldName() string {
	return "SubSecCreateDate"
}

type subSecond struct{}

func (subSecond) ExifDateFormat() string {
	return "2006:01:02 15:04:05.99999"
}

func (subSecond) FileFormatWithoutExtension() string {
	return "2006-01-02 15.04.05.99999"
}

type SubSecondDateTimeOriginal struct {
	subSecond
}

func (SubSecondDateTimeOriginal) FieldName() string {
	return "SubSecDateTimeOriginal"
}
