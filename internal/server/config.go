package server

import (
	"fmt"
	"time"

	"github.com/NathanBak/maas/internal/astra"
)

// Config contains information necessary to set up a Server.
type Config struct {
	Port         int           `json:"port" envvar:"PORT,required"`
	ReadTimeout  time.Duration `json:"readTimeout" envvar:"READ_TIMEOUT,default=3s"`
	WriteTimeout time.Duration `json:"writeTimeout" envvar:"WRITE_TIMEOUT,default=3s"`

	Logger Logger `json:"-" envvar:"-"`

	CallCounter CallCounter `json:"-" envvar:"-"`

	IncludeStatusCodeInMessages bool `json:"-" envvar:"-"`

	AstraConfig astra.Config `json:"-" envvar:">,prefix=ASTRA_"`
}

// The Logger interface defines the methods required by the Server for logging.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type CallCounter interface {
	Increment(id string, inc int) error
	Lookup(id string) (int, error)
}

// CfgBuildInit initializes the Logger and CallCounter.  It should only be called by a cfgbuild.Builder.
func (cfg *Config) CfgBuildInit() error {
	if cfg.Logger == nil {
		cfg.Logger = defaultLogger{}
	}
	return nil
}

// CfgBuildValidate checks the specified values.  It should only be called by a cfgbuild.Builder.
func (cfg *Config) CfgBuildValidate() error {
	if cfg.Port < 1 || cfg.Port > 65535 {
		return fmt.Errorf("%d is not a valid port", cfg.Port)
	}

	var err error
	if cfg.CallCounter == nil {
		if cfg.AstraConfig.Token != "" {
			cfg.CallCounter, err = astra.NewClient(cfg.AstraConfig.URI, cfg.AstraConfig.Token, cfg.AstraConfig.Keyspace)
			if err != nil {
				return err
			}
		} else {
			cfg.CallCounter = &inMemoryCallCounter{cache: make(map[string]int)}
		}
	}

	return nil
}
