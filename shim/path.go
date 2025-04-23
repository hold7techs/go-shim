package shim

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

// FindFilePaths 寻找文件dir的目录下有多少md文件
func FindFilePaths(root string, pattern string) ([]string, error) {
	// 检测root是否为常规文件
	stat, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if stat.Mode().IsRegular() {
		return []string{root}, nil
	}

	// change root
	if err := os.Chdir(root); err != nil {
		return nil, err
	}

	// 通过walk 遍历目录，将所有子目录都找到
	var paths []string
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		// 将匹配的信息捞出来
		if match, err := filepath.Match(pattern, info.Name()); err != nil {
			return err
		} else if match {
			paths = append(paths, path)
		}

		return nil
	})

	return paths, nil
}

// GetRootPath 根据给定的相对路径返回项目根目录
func GetRootPath(realpath string) string {
	// 获取当前文件的绝对路径
	currentPath, _ := os.Getwd()
	start := strings.Index(currentPath, realpath)
	if start == -1 {
		return currentPath
	} else if start == 0 {
		return "/"
	}

	return currentPath[:start-1]
}
