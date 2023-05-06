package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

// ListDirectoryA returns the content of dir.
func ListDirectoryA(dir string) ([]string, error) {
	// TODO 2023/4/20 16:24 sun: 遍历整个目录层级，返回各个层级中的目录名

	return nil, nil
}

// ListDirectoryB returns a channel over which directory entries will be published.
// When the list of entries is exhausted, the channel will be closed.
func ListDirectoryB(dir string) chan string {
	// TODO 2023/4/20 16:27 sun: 创建一个 Channel，通过 goroutine 将遍历到的目录存入 channel 中
	return nil
}

func main() {
	err := filepath.WalkDir("errhandling/", func(path string, d fs.DirEntry, err error) error {
		fmt.Printf("[path=%s], [name=%s], [isDirectory=%v]\n", path, d.Name(), d.IsDir())
		return nil
	})
	if err != nil {
		log.Fatalf("impossible to walk directories: %s", err)
	}
}
