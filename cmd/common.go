package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/axgle/mahonia"
	"github.com/go-ini/ini"
)

func readFileData(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if nil != err {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if nil != err {
		return nil, err
	}
	return data, err
}

func readIniFile(filename string) (*ini.File, error) {
	fileData, err := readFileData(filename)
	if nil != err {
		return nil, err
	}

	// transcode to utf-8
	dec := mahonia.NewDecoder("gbk")
	fileData = []byte(dec.ConvertString(string(fileData)))
	if nil == fileData ||
		len(fileData) == 0 {
		return nil, errors.New("Transcode to utf-8 failed")
	}

	return ini.InsensitiveLoad(fileData)
}
