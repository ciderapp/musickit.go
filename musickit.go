package musickit

import "log"

type MusicKit struct {
	Config *Config
}

type Config struct {
	DeveloperToken string
	SourceType     int
	Verbose        bool
}

func (m *MusicKit) Configure(config *Config) {
	m.Config = config
	if m.Config.Verbose {
		log.Println("Configured musickit.go")
	}
}
