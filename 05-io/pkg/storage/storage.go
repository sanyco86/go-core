package storage

import (
	"encoding/json"
	"go-core/pkg/crawler"
	"io"
	"os"
)

//Write для записи в файл
func Write(docs []crawler.Document, fileName string) (bool, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	defer f.Close()

	data, err := json.Marshal(docs)
	if err != nil {
		return false, err
	}
	err = os.WriteFile(f.Name(), data, 0666)
	if err != nil {
		return false, err
	}
	return true, err
}

//Read для чтения из файла
func Read(fileName string) ([]crawler.Document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	var data []byte
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {
			data = append(data, buf[:n]...)
		}
	}
	var res []crawler.Document
	json.Unmarshal([]byte(data), &res)
	return res, nil
}

//Empty для проверки наличия
func Empty(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return true
	}
	fileInfo, err := os.Lstat(fileName)
	if fileInfo.Size() == 0 {
		return true
	}
	if err != nil {
		return true
	}
	return false
}
