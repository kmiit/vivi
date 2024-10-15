package storage

import (
	"github.com/fsnotify/fsnotify"
	"github.com/kmiit/vivi/utils/log"
)

func WatchStorage() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.F(TAG, err)
	}
	defer watcher.Close()

	for _, p := range ExistDir {
		err = watcher.Add(p)
		if err != nil {
			log.F(TAG, err)
		} else {
			log.I(TAG, "Watching: ", p)
		}
	}
}
