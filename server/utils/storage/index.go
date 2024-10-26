package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/db"
	"github.com/kmiit/vivi/utils/log"
	"github.com/redis/go-redis/v9"
)

func InitIndex() {
	// Initial index if file no file indexed
	allFileKeys, err := db.GetKeys(ctx, db.FILE_NAMESPACE)
	if err != nil {
		log.E(TAG, err)
	}
	if len(allFileKeys) == 0 {
		for _, dir := range ExistDir {
			log.V(TAG, "Mapping storage:", dir)
			// Map storages and files in them
			id, _ := db.GetNewId(ctx, db.STORAGE_UNIQUE_ID)
			db.Set(ctx, db.FILE_MAP_NAMESPACE+dir, id, 0)
			MapAll(dir)
		}
	}
}

// map all files and directories in the given directory
// Usually used when a new storage or folder added.
func MapAll(dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.E(TAG, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dir := filepath.Join(dir, entry.Name())
			MapAll(dir)
		} else {
			file := filepath.Join(dir, entry.Name())
			NewDescriptor(file)
		}
	}
}

// Creates a new descriptor for the given path
// p: path of file or dir.
// returns the id of file or dir, in string.
func NewDescriptor(p string) (int64, error) {
	log.V(TAG, "NewDescriptor triggered with path:", p)
	// Try to get the id of given file, return "" if already exists.
	id, err := db.GetIdByPath(ctx, p)
	if err == nil {
		log.W(TAG, "File or dir already mapped!")
		return  id, nil
	}

	// Try to get the parent id of file, create a map if not exists.
	parentPath, _ := filepath.Split(p)
	if parentPath[len(parentPath)-1:] == "/" {
		parentPath = parentPath[:len(parentPath)-1]
	}
	pID, err := db.GetIdByPath(ctx, parentPath)
 	if err == redis.Nil {
		log.W(TAG, "Parent path doesn't exist, mapping:", parentPath)
		pID, _ = NewDescriptor(parentPath)
	}

	id, _ = db.GetNewId(ctx, db.FILE_UNIQUE_ID)
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

	idS := strconv.FormatInt(id, 10)
	db.Set(ctx, db.FILE_NAMESPACE + idS, j, 0)
	db.Set(ctx, db.FILE_MAP_NAMESPACE + p, id, 0)
	return id, nil
}

// Func to remove a mapped file.
func RemoveFile(p string) {
	id, err := db.GetIdByPath(ctx, p)
	if err != nil {
		log.E(TAG, err)
	}

	idS := strconv.FormatInt(id, 10)
	db.Del(ctx, db.FILE_MAP_NAMESPACE + p)
	db.Del(ctx, db.FILE_NAMESPACE + idS)
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
