package configure

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
)

const localConfigDir = "./configs/"

//Cfg app level config variable
var Cfg *Config

//Config app configuration structure
type Config struct {
	Env string `config:"ENVIRONMENT"`
	DB  *DB    `config:"DB"`
}

//DB database configuration structure
type DB struct {
	DBName string `config:"DBName"`
	DBUser string `config:"DBUser"`
	DBPass string `config:"DBPass"`
}

//LoadConfigs apply configuration in certain sources order
func LoadConfigs() error {
	Cfg = new(Config)
	Cfg.DB = &DB{}
	loader := confita.NewLoader(
		file.NewBackend(localConfigDir + "local.toml"),
	)
	err := loader.Load(context.Background(), Cfg) //load local config
	if err != nil {
		println("local configs load error") //if read local file error - hope for env vars data
	}
	loader = confita.NewLoader(
		env.NewBackend(),
	)
	err = loader.Load(context.Background(), Cfg) // load env config
	if err != nil {
		return errors.Wrap(err, "configs load error")
	}
	return nil
}
