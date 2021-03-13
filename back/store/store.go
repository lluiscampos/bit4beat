package store

import (
	"github.com/lluiscampos/bit4beat.back/model"
)

type Store interface {
	CreateRecord(*model.Record) error
	ReadRecord(id string) (*model.Record, error)
	ListRecords() ([]model.Record, error)
}
