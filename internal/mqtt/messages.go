package mqtt

import "github.com/EldoranDev/xmaspi/v3/internal/led"

// Representation of the mqtt light json schema https://www.home-assistant.io/integrations/light.mqtt/#json-schema

type State struct {
	Brightness *uint8     `json:"brightness"`
	Color      *led.Color `json:"color"`
	Effect     *string    `json:"effect"`
	State      string     `json:"state"`
}

type Config struct {
	Device Device `json:"dev"`
	Origin Origin `json:"o"`

	Components map[string]Component `json:"cmps"`
}

// Representation of the mqtt device discovery schema https://www.home-assistant.io/integrations/mqtt/#discovery-messages

type Device struct {
	Model        string   `json:"mdl"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"mf"`
	SwVersion    string   `json:"sw"`
	Identifiers  []string `json:"identifiers"`
}

type Origin struct {
	Name            string `json:"name"`
	SoftwareVersion string `json:"sw"`
	Url             string `json:"url"`
}

type Component struct {
	Schema       string       `json:"schema"`
	StateTopic   string       `json:"state_topic"`
	CommandTopic string       `json:"command_topic"`
	Availability Availability `json:"availability"`
	Platform     string       `json:"p"`
	Qos          string       `json:"qos"`
	UniqueId     string       `json:"unique_id"`

	Name string `json:"name"`

	Brightness          bool     `json:"brightness"`
	Effect              bool     `json:"effect"`
	EffectList          []string `json:"effect_list"`
	SupportedColorModes []string `json:"supported_color_modes"`
}

type Availability struct {
	Topic string `json:"topic"`
}
