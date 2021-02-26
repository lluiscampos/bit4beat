package filestore

import (
	"github.com/lluiscampos/bit4beat.back/model"
	"github.com/lluiscampos/bit4beat.back/store"
)

type FileStore struct {
	rootDir string
}

func NewStore() store.Store {
	return &FileStore{rootDir: "/dummy/path"}
}

func (fs *FileStore) GetRecords() ([]model.Record, int, error) {
	return nil, 0, nil
}

func (fs *FileStore) GetRecord(id int) (*model.Record, error) {
	r := &model.Record{ID: id}
	return r, nil
}
