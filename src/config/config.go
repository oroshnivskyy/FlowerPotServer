package config

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"io"
	"os"
	"time"
	"strings"
	"github.com/client9/reopen"
)

type WebSocketConfig struct {
	WriteWait       Duration
	// Time allowed to read the next pong message from the peer.
	PongWait        Duration
	// Send pings to peer with this period. Must be less than pongWait.
	PingPeriod      Duration
	// Maximum message size allowed from peer.
	MaxMessageSize  int64
	ReadBufferSize  int
	WriteBufferSize int
}
type DatabaseConnectOpts struct {
	Addresses           []string
	Database            string
	AuthKey             string
	DiscoverHosts       bool
	Address             string
	Timeout             Duration
	WriteTimeout        Duration
	ReadTimeout         Duration
	MaxIdle             int
	MaxOpen             int
	NodeRefreshInterval Duration
	Log                 bool
	logFile            io.Writer
}
func (d DatabaseConnectOpts) LogFile() 	io.Writer{
	return d.logFile
}

type Logger struct {
	LogFilePath  string
	Debug        bool
	LoggingLevel string
	LogFile      *reopen.FileWriter
}
type HttpServer struct {
	ListenHost string
	ListenPort string
}
type Config struct {
	Logger
	DatabaseConnectOpts
	HttpServer
	WebSocketConfig
}

func GetConfig(configFilePath string) (config *Config, err error) {
	var (
		file *os.File
		buf  []byte
	)
	file, err = os.Open(configFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	buf, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	config = new(Config)
	if err = toml.Unmarshal(buf, config); err != nil {
		return
	}
	config.Logger.LogFile, err = reopen.NewFileWriter(config.Logger.LogFilePath)

	if config.DatabaseConnectOpts.Log {
		config.DatabaseConnectOpts.logFile = config.Logger.LogFile
	}else {
		config.DatabaseConnectOpts.logFile = ioutil.Discard
	}
	return
}

type Duration struct {
	Int      int
	Duration time.Duration
}

func (d *Duration) UnmarshalTOML(data []byte) (err error) {
	d.Duration, err = time.ParseDuration(strings.Trim(string(data), "\""))
	d.Int = int(d.Duration.Seconds())
	return
}
