package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newRecordMinValid() *Record {
	return &Record{
		ID:    "stuff-id",
		Kind:  "ski",
		Date:  "sometime",
		Place: "somewhere",
	}
}

func TestRequiredFields(t *testing.T) {
	rValid := newRecordMinValid()
	err := rValid.Validate()
	assert.NoError(t, err)

	rMissingID := newRecordMinValid()
	rMissingID.ID = ""
	err = rMissingID.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'ID' failed on the 'required' tag")

	rMissingKind := newRecordMinValid()
	rMissingKind.Kind = ""
	err = rMissingKind.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'Kind' failed on the 'required' tag")

	rMissingDate := newRecordMinValid()
	rMissingDate.Date = ""
	err = rMissingDate.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'Date' failed on the 'required' tag")

	rMissingPlace := newRecordMinValid()
	rMissingPlace.Place = ""
	err = rMissingPlace.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "'Place' failed on the 'required' tag")
}

func TestValidateKind(t *testing.T) {
	rInvalidKind := newRecordMinValid()
	rInvalidKind.Kind = "invalid"

	err := rInvalidKind.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation for 'Kind' failed")
}

func TestValidateDistance(t *testing.T) {
	validDistances := []string{
		"123 km",
		"1 m",
		"-1 km",
		"987834726483217468127341827987 km",
		"5.5 m",
	}
	for _, d := range validDistances {
		rValidDistance := newRecordMinValid()
		rValidDistance.Distance = d

		err := rValidDistance.Validate()
		assert.NoError(t, err)
	}

	invalidDistances := []string{
		"123 somethingelse",
		" 123 m",
		"123 m ",
		"+123 m",
		"m m",
	}
	for _, d := range invalidDistances {
		rInvalidDistance := newRecordMinValid()
		rInvalidDistance.Distance = d

		err := rInvalidDistance.Validate()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid Distance format")
	}
}
