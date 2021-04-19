package gpee

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
)

func respContent(body io.ReadCloser) (string, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return respMessage(bodyBytes)
}

func respMessage(bodyBytes []byte) (string, error) {
	contentBytes, err := unzip(bodyBytes)
	if err != nil {
		return "", err
	}

	return to1251(contentBytes)
}

func to1251(content []byte) (string, error) {
	dec := charmap.Windows1251.NewDecoder()
	out, err := dec.Bytes(content)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func unzip(content []byte) ([]byte, error) {
	str := base64.StdEncoding.EncodeToString(content)
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
