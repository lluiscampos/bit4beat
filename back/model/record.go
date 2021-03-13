package model

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	distanceReg = regexp.MustCompile(`^[-0-9]+ (km|m)$`)
)

type Record struct {
	// Auto-generated on insert
	ID string `json:"id" validate:"required"`

	// One of ski, hike, bike
	Kind string `json:"kind" validate:"required,eq=ski|eq=hike|eq=bike"`

	// Any string will do for now
	Date string `json:"date" validate:"required"`

	// Any string
	Place string `json:"place" validate:"required"`

	// Accepts .5 precision, suffixed with unit
	Distance string `json:"distance,omitempty"`

	Participants []string `json:"participants,omitempty`

	Reference string `json:"reference,omitempty"`
}

func (r Record) Validate() error {
	v := validator.New()
	err := v.Struct(r)
	if err != nil {
		return err
	}

	// Custom validation for Distance
	if r.Distance != "" {
		match := distanceReg.MatchString(r.Distance)
		if !match {
			return errors.New(fmt.Sprintf("invalid Distance format in %q", r.Distance))
		}
	}

	return nil
}
