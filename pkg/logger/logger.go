package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// SetFields set configure fields
func SetFields(fields map[string]interface{}) *log.Entry {
	return log.WithFields(fields)
}

// WithNamespace returns a logger with the specified nspace field.
func WithNamespace(nspace string) *log.Entry {
	return log.WithField("nspace", nspace)
}
