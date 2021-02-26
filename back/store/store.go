package store

import (
	"github.com/lluiscampos/bit4beat.back/model"
)

type Store interface {
	GetRecords() ([]model.Record, int, error)
	GetRecord(id int) (*model.Record, error)
}
