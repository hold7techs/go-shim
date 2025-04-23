package shim

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestFindFilePaths(t *testing.T) {
	// os.TempDir() -> /var/folders/n4/2qd9ttqd0b927rbmvy6fn7980000gn/T/
	t.Logf("temp dir: %s", os.TempDir())

	// pattern -> /var/folders/n4/2qd9ttqd0b927rbmvy6fn7980000gn/T/example1053153493
	f, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("file name: %s", f.Name())
	// defer os.Remove(f.Name()) // clean up

	// stat
	stat, err := f.Stat()
	if err != nil {
		return
	}
	t.Logf("stat: %+v", stat)

	// 写入信息 f.Write
	if _, err := f.Write([]byte("content")); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestWalkRoot(t *testing.T) {
	tempDir := os.TempDir()
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	// 目录创建
	for _, dir := range []string{"dir1", "dir2"} {
		newDir := filepath.Join("test", dir)
		err = os.MkdirAll(newDir, 0750)
		if err != nil && !os.IsExist(err) {
			t.Fatal(err)
		}
		defer os.RemoveAll(newDir)

		// 写入文件
		for i := 0; i < 5; i++ {
			err = os.WriteFile(fmt.Sprintf("%s/testfile_%d.txt", newDir, i), []byte("Hello, Gophers!"), 0660)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	// 查看当前所在目录
	var filePaths []string
	filepath.Walk(filepath.Join(tempDir, "test"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			t.Logf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// 匹配目录
		if match, err := filepath.Match("*3*", info.Name()); err != nil {
			t.Errorf("filepath match 3* got err: %s", err)
			return err
		} else if match {
			t.Logf("match file: %s", info.Name())
		}

		// 文件是目录，且文件名称是要过滤的名称，则过滤
		if info.IsDir() && info.Name() == `skip` {
			t.Logf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		// 将所有文件都收集起来
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}

		t.Logf("visited file or dir: %q\n", path)
		return nil
	})

	// 遍历得到的信息
	for _, filePath := range filePaths {
		t.Logf("file: %s", filePath)
	}
}
