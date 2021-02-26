package store

import (
	"github.com/lluiscampos/bit4beat.back/model"
)

type Store interface {
	CreateRecord(*model.Record) error
	ReadRecord(id int) (*model.Record, error)
	ListRecords() ([]model.Record, error)
}
