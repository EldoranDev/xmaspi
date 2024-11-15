package mqtt

const DiscoveryPrefix = "homeassistant"

const TopicBase = "xmaspi"

const StateTopic = TopicBase + "/state"
const CommandTopic = TopicBase + "/set"
const AvailabilityTopic = TopicBase + "/availability"

const DiscoveryTopic = DiscoveryPrefix + "/device/" + TopicBase + "/config"
