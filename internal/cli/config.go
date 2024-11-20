package cli

import (
	"github.com/EldoranDev/xmaspi/v3/internal/config"
	"github.com/spf13/viper"
)

func GetConfig() *config.Config {
	cfg := config.Config{}

	viper.BindEnv("data_pin", "DATA_PIN")

	viper.BindEnv("led.file", "LED_FILE")
	viper.BindEnv("led.max_brightness", "LED_MAX_BRIGHTNESS")

	viper.BindEnv("mqtt.broker", "MQTT_BROKER")
	viper.BindEnv("mqtt.user", "MQTT_USER")
	viper.BindEnv("mqtt.password", "MQTT_PASSWORD")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
