package logger
import (
	"config"
	"github.com/op/go-logging"
	"os"
)
func GetLogger(cfg *config.Logger) (*logging.Logger, error) {
	var (
		log = logging.MustGetLogger("main")
		err error
		logLevel logging.Level
		logFormat = "%{time:2006-01-02 15:04:05.999} %{module} - [%{level}] %{shortfile} :: %{message}"
		logFormatter = logging.MustStringFormatter(logFormat)
		coloredLogFormatter = logging.MustStringFormatter("%{color}" + logFormat + "%{color:reset}")
	)
	logLevel, err = logging.LogLevel(cfg.LoggingLevel)
	if err != nil {
		return nil, err
	}

	var logBackends []logging.Backend
	b1 := logging.NewBackendFormatter(logging.NewLogBackend(cfg.LogFile, "", 0), logFormatter)
	logBackends = append(logBackends, b1)
	b2 := logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), coloredLogFormatter)
	logBackends = append(logBackends, b2)

	logging.SetBackend(logBackends...).SetLevel(logLevel, "")

	return log, nil
}