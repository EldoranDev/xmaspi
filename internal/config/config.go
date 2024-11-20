package config

type Config struct {
	DataPin string `mapstructure:"data_pin"`

	Led  LedConfig  `mapstructure:"led"`
	Mqtt MqttConfig `mapstructure:"mqtt"`
}

type MqttConfig struct {
	Broker   string `mapstructure:"broker"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type LedConfig struct {
	MaxBrightness string `mapstructure:"max_brightness"`
	LedFile       string `mapstructure:"file"`
}
