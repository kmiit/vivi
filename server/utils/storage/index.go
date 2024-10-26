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
	// Do reMap if not the first time to run
	if _, err := db.Get(ctx, db.IS_FIRST_RUN); err != redis.Nil {
		log.V(TAG, "Syncing changed files while service offline...")
		reMap()
	} else {
		log.I(TAG, "First run detected!")
	}

	// always do MapAll, mainly handles new file added
	for _, dir := range ExistDir {
		_, err := db.GetIdByPath(ctx, db.FILE_MAP_NAMESPACE + dir)
		if err == redis.Nil {
			// Map storages newly added
			id, _ := db.GetNewId(ctx, db.STORAGE_UNIQUE_ID)
			db.Set(ctx, db.FILE_MAP_NAMESPACE + dir, id, 0)
			MapAll(dir)
		} else {
			MapAll(dir)
		}
	}
	db.Set(ctx, db.IS_FIRST_RUN, false, 0)
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
			NewDescriptor(dir)
			MapAll(dir)
		} else {
			file := filepath.Join(dir, entry.Name())
			NewDescriptor(file)
		}
	}
}

// Creates a new descriptor for the given path
// p: path of file or dir.
// returns the id of file or dir.
func NewDescriptor(p string) (int64, error) {
	log.V(TAG, "NewDescriptor triggered with path:", p)
	// Try to get the id of given file, return the id if already exists.
	id, err := db.GetIdByPath(ctx, p)
	if err == nil {
		log.V(TAG, "File or dir already mapped:", p)
		return id, nil
	}

	// Try to get the parent id of file, create a map if not exists.
	parentPath, _ := filepath.Split(p)
	if parentPath[len(parentPath)-1:] == "/" {
		parentPath = parentPath[:len(parentPath)-1]
	}
	pID, err := db.GetIdByPath(ctx, parentPath)
 	if err == redis.Nil {
		log.W(TAG, "Path doesn't exist, mapping:", parentPath)
		pID, _ = NewDescriptor(parentPath)
	}

	// Map new file or dir
	id, _ = db.GetNewId(ctx, db.FILE_UNIQUE_ID)
	des := types.FDescriptor{Path: p}
	if info, _ := os.Stat(p); info.IsDir() {
		newDirDescriptor(&des, pID, id)
	} else {
		newFileDescriptor(&des, pID, id)
	}
	j, _ := json.Marshal(des)

	// Save to database
	idS := strconv.FormatInt(id, 10)
	db.Set(ctx, db.FILE_NAMESPACE + idS, j, 0)
	db.Set(ctx, db.FILE_MAP_NAMESPACE + p, id, 0)
	return id, nil
}

// Func to remove a mapped file.
// p: path of removed file
func RemoveFile(p string) {
	log.V(TAG, "Removing:", p)
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
	d.Public.FullName = dir
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
	f.Public.Related = findRelated(parent, f)
}

// Find Related files such as ass file in the directory
// p: parent path
// f: file descriptor
func findRelated(p string, f *types.FDescriptor) (related []string) {
	var files []string
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil
	}
	for _, file := range files {
		_, fn := filepath.Split(file)
		if fn[:len(fn)-len(filepath.Ext(fn))] == f.Public.Name && fn != f.Public.FullName {
			related = append(related, fn)
		}
	}

	return related
}

// Redo mapping files
// handles: IsDir changed, file removed
// New file added already handled by MapAll()
func reMap() {
	keys, err := db.GetKeys(ctx, db.FILE_MAP_NAMESPACE)
	if err != nil {
		log.F(TAG, err)
	}
	for _, key := range keys {
		filePath := key[len(db.FILE_MAP_NAMESPACE):]
		info, err := os.Stat(filePath)
		switch {
		case err == nil:
			file, _ := db.GetPublic(ctx, key)
			// Handle file changes, file id always>0, storage id always<0.
			if info.IsDir() != file.IsDir && file.ID > 0 {
				log.V(TAG, "File type changed:", filePath)
				RemoveFile(filePath)
				NewDescriptor(filePath)
			}
		case os.IsNotExist(err):
			// File removed
			log.V(TAG, "File or dir removed:", filePath)
			RemoveFile(filePath)
		default:
			// unknown error
			log.F(TAG, "Unknown error:", err)
		}
	}
}
