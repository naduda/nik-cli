package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
)

type NikUtils struct {
}

func (u NikUtils) Unzip(content []byte) ([]byte, error) {
	str := base64.StdEncoding.EncodeToString(content)
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func (u NikUtils) To1251(content []byte) (string, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(content)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
