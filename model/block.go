package model

type Block struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	MetricID int    `json:"metric_id,omitempty"`
}
