package storage

import (
	"fmt"

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

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.I(TAG, "Event: ", event)
				fmt.Print(event.String())
				switch {
				case event.Has(fsnotify.Write):
					log.I(TAG, "Modified file: ", event.Name)
				case event.Has(fsnotify.Create):
					log.I(TAG, "Created file: ", event.Name)
				case event.Has(fsnotify.Remove):
					log.I(TAG, "Removed file: ", event.Name)
				case event.Has(fsnotify.Rename):
					log.I(TAG, "Renamed file: ", event.Name)
				case event.Has(fsnotify.Chmod):
					log.I(TAG, "Changed permission file: ", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.E(TAG, "Error: ", err)
			}
		}
	}()
	<-make(chan struct{})
}
