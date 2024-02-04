package db

// These configurations are required to start the conductor.
type ConductorConfig struct {
	Historian historian `json:"historian"`
	Security  security  `json:"security"`
	Atlas     atlas     `json:"atlas"`
}
type historian struct {
	RedisClusterEndpoint string  `json:"redis_cluster_endpoint"`
	RedisClusterPort     float64 `json:"redis_cluster_port"`
}
type security struct{}
type atlas struct{}

// This struct contains the values required by the conductor to be setup and ready.
type Conductor struct {
}
