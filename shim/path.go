package go_shim

import (
	"log"
	"os"
	"path/filepath"
)

// MustGetFilePath 如果改路径不存在，则自动创建，若存在，则直接返回对应的路径名
func MustGetFilePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("MustGetFilePath filepath.Abs() got err: %s", err)
	}

	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(absPath, os.ModePerm)
		if err != nil {
			log.Fatalf("MustGetFilePath os.MkdirAll() got err: %s", err)
		}
	}

	return absPath
}
