package storage

import (
	"os"
	"path"
	"path/filepath"

	"github.com/kmiit/vivi/utils/config"
	"github.com/kmiit/vivi/utils/log"
)

const TAG = "Storage"

var ExistDir []string

func InitStorage() {
	// Get the current executable path
	exp, _ := os.Executable()
	pwd := filepath.Dir(exp)

	for _, p := range config.StorageConfig.WatchPath {
		// Check absolute path of monitPath exists or not
		fp := path.Join(pwd, p)
		if len(p) > 0 && p[:1] == "/" {
			fp = p
		}
		info, err := os.Stat(fp)
		switch {
		case err == nil:
			if info.IsDir() {
				log.I(TAG, "Path found: ", p)
				ExistDir = append(ExistDir, fp)
			} else {
				log.E(TAG, "Path found but isn't a directory: ", p)
			}
		case os.IsNotExist(err):
			//	path not found
			log.W(TAG, "Path not found: ", p)
			log.I(TAG, "Creating directory: ", p)
			if err := os.MkdirAll(fp, os.ModePerm); err != nil {
				log.E(TAG, "Failed to create directory: ", p)
			} else {
				log.I(TAG, "Directory created: ", p)
				ExistDir = append(ExistDir, fp)
			}
		case os.IsPermission(err):
			// permission denied
			log.E(TAG, "Have no permission to visit path:", p)

		default:
			// unknown error
			log.F(TAG, "Unknown error: %v\n", err)
		}
	}
	InitIndex()
}
