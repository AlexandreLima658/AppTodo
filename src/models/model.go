package models

type Todo struct {
	ID   uint64 `json:"id,omitempty"`
	Item string `json:"item,omitempty"`
}
