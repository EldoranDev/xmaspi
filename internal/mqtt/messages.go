package mqtt

// Following the mqtt light json schema https://www.home-assistant.io/integrations/light.mqtt/#json-schema

type State struct {
	Brightness uint8
	Color      Color
	Effect     string
	State      string
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

type Config struct {
	Dev struct {
		Model     string `json:"model"`
		Name      string `json:"name"`
		SwVersion string `json:"sw_version"`
	} `json:"dev"`

	Schema       string `json:"schema"`
	StateTopic   string `json:"state_topic"`
	CommandTopic string `json:"command_topic"`
	Platform     string `json:"platform"`
	Qos          string `json:"qos"`
	UniqueId     string `json:"unique_id"`

	Brightness          bool     `json:"brightness"`
	Effect              bool     `json:"effect"`
	EffectList          []string `json:"effect_list"`
	SupportedColorModes []string `json:"supported_color_modes"`
}
