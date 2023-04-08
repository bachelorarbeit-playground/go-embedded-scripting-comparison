package model

type RawWindDataPayload struct {
	ParkId       string  `json:"park_id"`
	TurbineId    string  `json:"turbine_id"`
	Date         string  `json:"date"`
	Interval     int     `json:"interval"`
	Timezone     string  `json:"timezone"`
	Value        float32 `json:"value"`
	Availability int     `json:"availability"`
	Region       string  `json:"region"`
}

type AnomalyDetectionPayload struct {
	ParkId       string  `json:"park_id"`
	Timestamp    string  `json:"timestamp"`
	Value        float32 `json:"value"`
	Availability float32 `json:"availability"`
	Region       string  `json:"region"`
}

type Fibonacci struct {
	N int `json:"n"`
}
