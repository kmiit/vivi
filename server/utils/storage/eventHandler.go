package storage

import (
	"github.com/fsnotify/fsnotify"
	"github.com/kmiit/vivi/utils/log"
)

// File add event handler, including CREATE and where RENAME to(CREATE).
func addEvent(event fsnotify.Event) {
	log.V(TAG, "Create event triggered:", event.Op, event.Name)
	NewDescriptor(event.Name)
}

// File remove event handler, including REMOVE and where RENAME from.  
func removeEvent(event fsnotify.Event) {
	log.V(TAG, "Remove event triggered:", event.Op, event.Name)
	RemoveFile(event.Name)
}
