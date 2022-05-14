package storage

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetDBPath() string {
	_, b, _, ok := runtime.Caller(0)
	basePath := filepath.Dir(b)

	if !ok {
		log.Fatal("[ERR]: GetDBPath ")
	}
	return basePath + "/english-bot.db"
}
