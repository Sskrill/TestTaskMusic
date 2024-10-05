package customLogger

import (
	"fmt"
	"log/slog"
	"os"
)

type CustomLogger struct {
	logger *slog.Logger
}

func NewCSLogger() CustomLogger {
	handler := slog.NewTextHandler(os.Stdout, nil)
	return CustomLogger{logger: slog.New(handler)}
}
func (csL CustomLogger) PrintInfo(info string, infoArgs ...interface{}) {
	csL.logger.Info(fmt.Sprintf(info, infoArgs))
}
func (csL CustomLogger) PrintError(error string, errorArgs ...interface{}) {
	csL.logger.Error(fmt.Sprintf(error, errorArgs))
}
func (csL CustomLogger) PrintDebug(debug string, debugArgs ...interface{}) {
	csL.logger.Debug(fmt.Sprintf(debug, debugArgs))
}
