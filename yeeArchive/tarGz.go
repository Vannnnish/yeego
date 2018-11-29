/**
 * Created by angelina on 2017/5/2.
 */

package yeeArchive

import (
	"archive/tar"
	"compress/gzip"
	"github.com/yeeyuntech/yeego/yeeFile"
	"io"
	"os"
	"path"
)

// TarGz
// gzip压缩
func TarGz(srcDir, targetFile string) error {
	yeeFile.MkdirForFile(targetFile)
	fw, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer fw.Close()
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	f, err := os.Open(srcDir)
	if err != nil {
		return err
	}
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return tarGzDir(srcDir, path.Base(srcDir), tw)
	}
	return tarGzFile(srcDir, fi.Name(), tw, fi)
}

func tarGzDir(srcDirPath string, recPath string, tw *tar.Writer) error {
	dir, err := os.Open(srcDirPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		curPath := path.Join(srcDirPath, fi.Name())
		if fi.IsDir() {
			err := tarGzDir(curPath, path.Join(recPath, fi.Name()), tw)
			if err != nil {
				return err
			}
		}
		err := tarGzFile(curPath, path.Join(recPath, fi.Name()), tw, fi)
		if err != nil {
			return err
		}
	}
	return nil
}

func tarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) error {
	if fi.IsDir() {
		hdr := new(tar.Header)
		hdr.Name = recPath + "/"
		hdr.Typeflag = tar.TypeDir
		hdr.Size = 0
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()
		err := tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
	} else {
		fr, err := os.Open(srcFile)
		if err != nil {
			return err
		}
		defer fr.Close()
		hdr := new(tar.Header)
		hdr.Name = recPath
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, fr)
		if err != nil {
			return err
		}
	}
	return nil
}

// UnTarGz
// 解压
func UnTarGz(srcFile, targetDir string) error {
	err := os.Mkdir(targetDir, os.ModePerm)
	if err != nil {
		return err
	}
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if hdr.Typeflag != tar.TypeDir {
			err := os.MkdirAll(path.Join(targetDir, path.Dir(hdr.Name)), os.ModePerm)
			if err != nil {
				return err
			}
			fw, _ := os.Create(path.Join(targetDir, hdr.Name))
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
