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

const (
	FILE_MAP_NAMESPACE = "file_map:"
	FILE_NAMESPACE    = "files:"
	FILE_UNIQUE_INDEX = "files_unique_id"
	STORAGE_UNIQUE_ID = "storage_unique_id"
)

var ctx = context.Background()

func InitIndex() {
	// Initial index if file no file indexed
	allFileKeys, err := db.GetKeys(ctx, FILE_NAMESPACE)
	if err != nil {
		log.E(TAG, err)
	}
	if len(allFileKeys) == 0 {
		for _, dir := range ExistDir {
			var (
				id, _ = db.GetNewId(ctx, STORAGE_UNIQUE_ID)
				pid   = "S" + strconv.FormatInt(id, 10)
			)
			mapAll(dir, pid)
		}
	}
}

// map all files and directories in the given directory
func mapAll(dir string, pChain string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.E(TAG, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dir := filepath.Join(dir, entry.Name())
			mapAll(dir, NewDescriptor(dir, pChain, true))
		} else {
			file := filepath.Join(dir, entry.Name())
			NewDescriptor(file, pChain, false)
		}
	}
}

// Creates a new descriptor for the given path
// returns the id Chain of the descriptor
func NewDescriptor(p string, pChain string, isDir bool) string {
	id, _ := db.GetNewId(ctx, FILE_UNIQUE_INDEX)
	idS := strconv.FormatInt(id, 10)
	var d types.Descriptor
	des := types.FDescriptor{Path: p}
	if isDir {
		newDirDescriptor(&des, pChain, id)
		d = &des
	} else {
		newFileDescriptor(&des, pChain, id)
		d = &des
	}
	j, _ := json.Marshal(d)

	fChain := pChain + ":" + idS
	db.RDB.Set(ctx, FILE_NAMESPACE+fChain, j, 0)
	db.RDB.Set(ctx, FILE_MAP_NAMESPACE+idS, fChain, 0)
	return fChain
}

func newDirDescriptor(d *types.FDescriptor, pChain string, id int64) {
	_, dir := filepath.Split(d.Path)
	d.Outer.ID = id
	d.Outer.IsDir = true
	d.Outer.Name = dir
	d.Outer.Parent = pChain
}

func newFileDescriptor(f *types.FDescriptor, pChain string, id int64) {
	parent, file := filepath.Split(f.Path)

	// Split the file name and extension
	ext := filepath.Ext(file)
	name := file[:len(file)-len(ext)]

	f.Outer.Ext = ext
	f.Outer.FullName = file
	f.Outer.ID = id
	f.Outer.IsDir = false
	f.Outer.Name = name
	f.Outer.Parent = pChain
	f.Outer.Related = findRelated(parent, name)
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
