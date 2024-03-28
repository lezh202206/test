package t_viper

import "sync"

type DataBaseConf struct {
	Driver string
	Source string
	Debug  bool
}

var (
	dbOnce sync.Once
	dbConf map[string]DataBaseConf
)

func DBViper() map[string]DataBaseConf {
	dbOnce.Do(func() {
		Viperr("db.yaml", &dbConf)
	})
	return dbConf
}
