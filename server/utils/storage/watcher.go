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
			log.I(TAG, "Watching:", p)
		}
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				switch {
				case event.Has(fsnotify.Create):
					addEvent(event)
				case event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename):
					removeEvent(event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.E(TAG, "Error:", err)
			}
		}
	}()
	<-make(chan struct{})
}
