package filestore

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lluiscampos/bit4beat.back/model"
)

func TestCreateRecord(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmpdir)

	store := NewStore(tmpdir)
	assert.NotNil(t, store)

	myRecord := &model.Record{ID: "42"}
	err = store.CreateRecord(myRecord)
	assert.NoError(t, err)

	content, err := ioutil.ReadFile(path.Join(tmpdir, "record_42.json"))
	assert.NoError(t, err)
	assert.Contains(t, string(content), `{"id":"42"`)

	myOtherRecord := &model.Record{ID: "unique-id-foo-bar"}
	err = store.CreateRecord(myOtherRecord)
	assert.NoError(t, err)

	content, err = ioutil.ReadFile(path.Join(tmpdir, "record_unique-id-foo-bar.json"))
	assert.NoError(t, err)
	assert.Contains(t, string(content), `{"id":"unique-id-foo-bar"`)
}

func TestReadRecord(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmpdir)

	store := NewStore(tmpdir)
	assert.NotNil(t, store)

	err = ioutil.WriteFile(path.Join(tmpdir, "record_42.json"), []byte(`{"id":"42"}`), 0644)
	require.NoError(t, err)

	myRecord, err := store.ReadRecord("42")
	assert.NoError(t, err)
	assert.Equal(t, "42", myRecord.ID)
}

func TestListRecords(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tmpdir)

	store := NewStore(tmpdir)
	assert.NotNil(t, store)

	err = ioutil.WriteFile(path.Join(tmpdir, "record_42.json"), []byte(`{"id":"42"}`), 0644)
	require.NoError(t, err)

	err = ioutil.WriteFile(path.Join(tmpdir, "record_bad.json"), []byte(`{"id":"47"}`), 0644)
	require.NoError(t, err)

	myRecords, err := store.ListRecords()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(myRecords))
}
