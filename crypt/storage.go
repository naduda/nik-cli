package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"nik-cli/lms/scheduler/model"
	"os"
	"sync"
)

type Storage struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	Data        []model.ConfigLms
}

func NewStorage(encodingKey, filepath string) *Storage {
	return &Storage{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Storage) Load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.Data = []model.ConfigLms{}
		return nil
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()
	r, err := DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readData(r)
}

func (v *Storage) ReadDataFromFile(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()
	return v.readData(f)
}

func (v *Storage) readData(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.Data)
}

func (v *Storage) Save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()
	w, err := EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Storage) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.Data)
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	_, err := fmt.Fprint(hasher, key)
	if err != nil {
		return nil, err
	}
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
