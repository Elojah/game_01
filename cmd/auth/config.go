package main

// Config is web quic server structure config.
type Config struct {
	Address string `json:"address"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
}
