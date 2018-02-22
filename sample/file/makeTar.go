package main

import (
	"path/filepath"
	"fmt"
	"os"
	"archive/tar"
	"strings"
	"io"
)

func makeTar(source string, target string) error {
	withDir := true
	filename := filepath.Base(source)
	fmt.Println("filename : "+filename)
	target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
	fmt.Println("target : "+target)
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		fmt.Println("source is dir : " + source)
		baseDir = filepath.Base(source)
	}

	return filepath.Walk("/tmp/tt",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(withDir)
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			fmt.Println(info)
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}
			fmt.Println(info)

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}
			return nil
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func main() {
	makeTar("/tmp/tt/2", "/tmp")
}
