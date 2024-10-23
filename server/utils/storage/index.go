package storage

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/db"
	"github.com/kmiit/vivi/utils/log"
)

var ctx = context.Background()

func InitIndex() {
	// Initial index if file no file indexed
	allFileKeys, err := db.GetKeys(ctx, db.FILE_NAMESPACE)
	if err != nil {
		log.E(TAG, err)
	}
	if len(allFileKeys) == 0 {
		for _, dir := range ExistDir {
			id, _ := db.GetNewId(ctx, db.STORAGE_UNIQUE_ID)
			MapAll(dir, id)
		}
	}
}

// map all files and directories in the given directory
// Usually used when a new storage or folder added.
func MapAll(dir string, pID int64) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.E(TAG, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dir := filepath.Join(dir, entry.Name())
			MapAll(dir, NewDescriptor(dir, pID))
		} else {
			file := filepath.Join(dir, entry.Name())
			NewDescriptor(file, pID)
		}
	}
}

// Creates a new descriptor for the given path
// p: path of file or dir.
// pID: parent folder id, top is storage id.
// returns the id of file or dir.
func NewDescriptor(p string, pID int64) int64 {
	id, _ := db.GetNewId(ctx, db.FILE_UNIQUE_ID)
	idS := strconv.FormatInt(id, 10)
	var d types.Descriptor
	des := types.FDescriptor{Path: p}
	if info, _ := os.Stat(p); info.IsDir() {
		newDirDescriptor(&des, pID, id)
		d = &des
	} else {
		newFileDescriptor(&des, pID, id)
		d = &des
	}
	j, _ := json.Marshal(d)

	db.Set(ctx, db.FILE_NAMESPACE+idS, j, 0)
	db.Set(ctx, db.FILE_MAP_NAMESPACE+p, idS, 0)
	return id
}

func newDirDescriptor(d *types.FDescriptor, pID int64, id int64) {
	_, dir := filepath.Split(d.Path)
	d.Public.ID = id
	d.Public.IsDir = true
	d.Public.Name = dir
	d.Public.Parent = pID
}

func newFileDescriptor(f *types.FDescriptor, pID int64, id int64) {
	parent, file := filepath.Split(f.Path)

	// Split the file name and extension
	ext := filepath.Ext(file)
	name := file[:len(file)-len(ext)]

	f.Public.Ext = ext
	f.Public.FullName = file
	f.Public.ID = id
	f.Public.IsDir = false
	f.Public.Name = name
	f.Public.Parent = pID
	f.Public.Related = findRelated(parent, name)
}

// Find Related files such as ass file in the directory
// p: parent path
// f: file name, usually FileDescriptor.Private.Name
func findRelated(p string, f string) (related []string) {
	var files []string
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
	}
	for _, file := range files {
		_, fn := filepath.Split(file)
		if fn[:len(fn)-len(filepath.Ext(fn))] == f {
			related = append(related, file)
		}
	}

	return related
}
