package bootstrap

import "flag"

type CommandFlags struct {
	Conf       string
	Env        string
	ConfigHost string
	ConfigType string
	Daemon     bool
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Conf:       "",
		Env:        "",
		ConfigHost: "",
		ConfigType: "",
		Daemon:     false,
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.Conf, "conf", "../../config", "config path, eg: -conf config.yaml")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev")
	flag.StringVar(&f.ConfigHost, "chost", "127.0.0.1:8500", "config server host, eg: -chost 127.0.0.1:8500")
	flag.StringVar(&f.ConfigType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.BoolVar(&f.Daemon, "d", false, "run app as a daemon with -d=true.")

	if f.Daemon {
		BeDaemon("-d")
	}

	flag.Parse()
}
