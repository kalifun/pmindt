package entity

import "encoding/json"

func UnmarshalPcmind(data []byte) (Pcmind, error) {
	var r Pcmind
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Pcmind) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Pcmind struct {
	Type string  `json:"type"`
	Data []Datum `json:"data"`
}

type Datum struct {
	ID             string       `json:"id"`
	Data           Data         `json:"data"`
	Children       []DatumChild `json:"children"`
	Width          int64        `json:"width"`
	Height         int64        `json:"height"`
	Layout         string       `json:"layout"`
	RightNodeCount int64        `json:"rightNodeCount"`
	IsRoot         bool         `json:"isRoot"`
	Type           string       `json:"type"`
	Points         [][]int64    `json:"points"`
}

type DatumChild struct {
	ID       string       `json:"id"`
	Data     Data         `json:"data"`
	Children []DatumChild `json:"children"`
	Width    int64        `json:"width"`
	Height   int64        `json:"height"`
}

type Data struct {
	Topic Topic `json:"topic"`
}

type Topic struct {
	Children []TopicChild `json:"children"`
}

type TopicChild struct {
	Text string `json:"text"`
}
