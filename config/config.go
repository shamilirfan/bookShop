package config

// Struct define for configaration
type Configaration struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
}

// Configaration type variable define
var config *Configaration

// Configaration loading function
func loadConfig() {
	ProjectConfig()
}
