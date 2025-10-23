package config

// Get loadConfig function
func GetConfig() *Configaration {
	if config == nil {
		loadConfig()
	}
	return config
}
