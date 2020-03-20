package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"miao_sticker_server/index/logger"
)

func readAllByFile(f *os.File) ([]byte, error) {
	if f != nil {
		return ioutil.ReadAll(f)
	}
	return []byte{}, errors.New("f is nil")
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteWithBytes(path string, info []byte) error {
	if err := ioutil.WriteFile(path+"__", info, 0666); err != nil {
		return err
	}
	if err := os.Rename(path+"__", path); err != nil {
		return err
	}
	return nil
}

func GetFile(path string) (*os.File, error) {
	var f *os.File
	var err error
	if exits, e := pathExists(path); !exits {
		if e == nil {
			if len(path) > 2 && path[len(path)-2:] != "__" {
				logger.Info("File Not Exits! Create File.")
			}
			f, err = os.Create(path)
		} else {
			logger.Error("CheckLogin With Error: %v", e)
			return nil, e
		}
	} else {
		logger.Info("File Exits! Open File.")
		f, err = os.Open(path)
	}
	return f, err
}

func WriteWithMultipartFile(path string, file multipart.File) (err error) {
	var outFile *os.File
	outFile, err = GetFile(path + "__")
	if err != nil {
		logger.Error("Create File Error: %v", err)
		return
	}
	_, err = io.Copy(outFile, file)
	if err = outFile.Close(); err != nil {
		logger.Error("File Close Error: %v", err)
	}
	if err != nil {
		logger.Error("Copy File Error: %v", err)
	}
	if err = os.Rename(path+"__", path); err != nil {
		logger.Error("Rename File Error: %v", err)
		return
	}
	return
}

func GetFileInfo(path string) ([]byte, error) {
	var file *os.File
	var err error
	var info []byte
	file, err = GetFile(path)
	if err != nil {
		logger.Error("GetFile Error: %v", err)
		return []byte{}, err
	}
	if info, err = readAllByFile(file); err != nil {
		logger.Error("ReadAllByFile Error: %v", err)
		return []byte{}, err
	}
	if err = file.Close(); err != nil {
		logger.Error("File Close Error: %v", err)
		return []byte{}, err
	}
	return info, nil
}
