package db
import (
	"github.com/dancannon/gorethink"
	"config"
)

func GetSession(cfg *config.DatabaseConnectOpts) (session *gorethink.Session, err error) {
	gorethink.Log.Out = cfg.LogFile()
	session, err = gorethink.Connect(gorethink.ConnectOpts{
		Addresses: cfg.Addresses,
		Database: cfg.Database,
		AuthKey:  cfg.AuthKey,
		DiscoverHosts: cfg.DiscoverHosts,
		Address:      cfg.Address,
		Timeout:      cfg.Timeout.Duration,
		WriteTimeout: cfg.WriteTimeout.Duration,
		ReadTimeout:  cfg.ReadTimeout.Duration,
		MaxIdle: cfg.MaxIdle,
		MaxOpen: cfg.MaxOpen,
		NodeRefreshInterval: cfg.NodeRefreshInterval.Duration,
	})
	return
}