package model

type Record struct {
	// Auto-generated on insert
	ID string `json:"id,omitempty"`

	// Kind
	Kind string `json:"kind"`

	// Date, any string will do for now
	Date string `json:"date"`

	Place string `json:"place"`

	// Accepts .5 precision, suffixed with unit
	Distance string `json:"distance,omitempty"`

	Participants []string `json:"participants,omitempty`

	Reference string `json:"reference"`
}

func (r Record) Validate() error {
	// TODO: Check Kind
	// TODO: Check Distance
	// ...
	return nil
}
