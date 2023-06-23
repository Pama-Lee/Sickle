package entity

type ConfigStruct struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	UUID         string        `json:"uuid"`
	Source       Source        `json:"source"`
	Destinations []Destination `json:"destinations"`
}

type Source struct {
	Type   string       `json:"type"`
	Config SourceConfig `json:"config"`
}

type SourceConfig struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name         string   `json:"name"`
	Key          string   `json:"key"`
	Destinations []string `json:"destinations"`
}

type Destination struct {
	Name   string            `json:"name"`
	Type   string            `json:"type"`
	Config DestinationConfig `json:"config"`
}

type DestinationConfig struct {
	WebhookURL string     `json:"webhook_url"`
	Method     string     `json:"method"`
	Data       []DataUnit `json:"data"`
	// From 1.0.0-alpha
	Headers map[string]string `json:"headers"`
}

type DataUnit struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}
