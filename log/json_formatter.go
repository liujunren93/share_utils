package log

import (
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type fieldKey string

// FieldMap allows customization of the key names for default fields.

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	Group         string
	jSONFormatter *logrus.JSONFormatter
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

const maximumCallerDepth = 25
const minimumCallerDepth = 8

func (f *JSONFormatter) getCaller() *runtime.Frame {
	maximumCallerDepth := 25
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(0, pcs)
	logrusPackage := ""
	// dynamic get the package name and the minimum caller depth
	for i := 0; i < maximumCallerDepth; i++ {
		funcName := runtime.FuncForPC(pcs[i]).Name()

		if strings.Contains(funcName, "getCaller") {
			logrusPackage = getPackageName(funcName)
			// fmt.Println(i, logrusPackage)
			break
		}
	}
	// Restrict the lookback frames to avoid runaway lookups
	pcs = make([]uintptr, maximumCallerDepth)
	depth = runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		// If the caller isn't part of this package, we're done
		if pkg != logrusPackage {
			return &f //nolint:scopelint
		}
	}

	return nil
}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Caller = f.getCaller()
	// fmt.Println(entry)
	if f.jSONFormatter == nil {
		f.jSONFormatter = &logrus.JSONFormatter{}
	}
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
