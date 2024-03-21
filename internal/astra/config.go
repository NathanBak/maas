package astra

// Config contains information necessary to set up a Server.
type Config struct {
	URI      string `envvar:"URI"`
	Keyspace string `envvar:"KEYSPACE"`
	Token    string `envvar:"TOKEN"`
}
