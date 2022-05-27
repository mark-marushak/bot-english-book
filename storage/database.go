package storage

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

func storageRoot() string {
	_, b, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal("[ERR]: storageRoot ")
	}

	return filepath.Dir(b)
}

func GetDBPath() string {
	return storageRoot() + "/english-bot.db"
}

func GetStorageRoot() string {
	return storageRoot()
}

func GetBookStorage(file string) string {
	return fmt.Sprintf("%s/%s/%s",
		storageRoot(),
		"book",
		file)
}
