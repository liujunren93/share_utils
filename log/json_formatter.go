package log

import (
	"github.com/sirupsen/logrus"
)

type fieldKey string

// FieldMap allows customization of the key names for default fields.

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	Group         string
	jSONFormatter *logrus.JSONFormatter
}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data, err := f.jSONFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	if f.Group == "" {
		return data, nil
	}
	tmp := []byte("[" + f.Group + "]:")
	tmp = append(tmp, data...)
	return tmp, err
}
