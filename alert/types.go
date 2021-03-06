package alert

import "time"

// Valid alert states.
const (
	AlertFiring   = "firing"
	AlertResolved = "resolved"
)

// Alert represents an alert received from alertmanager.
type Alert struct {
	allfail      int32
	current      int32
	Responders   []string          `json:"responders"`
	Status       string            `json:"status"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
}

// Data is the data received from alertmanager.
type Data struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	ExternalURL       string            `json:"externalURL"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	Alerts            []Alert           `json:"alerts"`
}
