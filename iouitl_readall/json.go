package iouitl_readall

import (
	"bytes"
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

func Json(r map[string]string) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	_ = data
	return nil
}

func JsonPool(r map[string]string) error {
	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)

	e := json.NewEncoder(buffer)
	err := e.Encode(r)
	if err != nil {
		return err
	}

	return nil
}

func JsonIter(r map[string]string) error {
	data, err := jsoniter.Marshal(r)
	if err != nil {
		return err
	}

	_ = data
	return nil
}

func JsonIterPool(r map[string]string) error {
	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)

	e := jsoniter.NewEncoder(buffer)
	err := e.Encode(r)
	if err != nil {
		return err
	}

	return nil
}
