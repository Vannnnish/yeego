/**
 * Created by angelina on 2017/5/2.
 */

package archive

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Zip
// Zip打包
func Zip(srcDir, targetFile string) error {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}
	fZip, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	w := zip.NewWriter(fZip)
	defer w.Close()
	for _, file := range files {
		fw, err := w.Create(file.Name())
		if err != nil {
			return err
		}
		fileContent, err := ioutil.ReadFile(path.Join(srcDir, file.Name()))
		if err != nil {
			return err
		}
		_, err = fw.Write(fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

// Unzip
// unzip解打包
func Unzip(srcFile, targetDir string) error {
	reader, err := zip.OpenReader(srcFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}
	for _, file := range reader.File {
		if err := unzip(file, targetDir); err != nil {
			return err
		}
	}
	return nil
}

// unzip
// 单独使用函数是因为循环体内进行了defer,防止内存泄漏
func unzip(file *zip.File, targetDir string) error {
	path := filepath.Join(targetDir, file.Name)
	if file.FileInfo().IsDir() {
		os.MkdirAll(path, file.Mode())
		return nil
	}
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()
	targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, fileReader); err != nil {
		return err
	}
	return nil
}
