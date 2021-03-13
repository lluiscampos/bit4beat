package filestore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/lluiscampos/bit4beat.back/model"
	"github.com/lluiscampos/bit4beat.back/store"
)

type FileStore struct {
	rootDir string
}

func NewStore(rootDir string) store.Store {
	return &FileStore{
		rootDir: rootDir,
	}
}

const (
	fileMode          = 0644
	recordFilenameFmt = "record_%s.json"
)

func (fs *FileStore) CreateRecord(record *model.Record) error {
	jsonBytes, err := json.Marshal(record)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fs.filepath(record.ID), jsonBytes, fileMode)
}

func (fs *FileStore) readRecordFromFile(filepath string) (*model.Record, error) {
	var r model.Record

	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileBytes, &r)
	return &r, err
}

func (fs *FileStore) ReadRecord(id string) (*model.Record, error) {
	return fs.readRecordFromFile(fs.filepath(id))
}

func (fs *FileStore) ListRecords() ([]model.Record, error) {
	files, err := ioutil.ReadDir(fs.rootDir)
	if err != nil {
		return nil, err
	}

	var records []model.Record
	for _, file := range files {
		r, err := fs.readRecordFromFile(path.Join(fs.rootDir, file.Name()))
		if err != nil {
			return nil, err
		}

		records = append(records, *r)
	}
	return records, nil
}

func (fs *FileStore) filepath(id string) string {
	filename := fmt.Sprintf(recordFilenameFmt, id)
	return path.Join(fs.rootDir, filename)
}
