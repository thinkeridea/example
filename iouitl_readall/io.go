package iouitl_readall

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}

func IoCopyAndJson(r io.Reader) error {
	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)

	res := Do(r)
	_, err := io.Copy(buffer, res)
	if err != nil {
		return err
	}

	m := map[string]string{}
	err = jsoniter.Unmarshal(buffer.Bytes(), &m)
	return err
}

func IouitlReadAllAndJson(r io.Reader) error {
	res := Do(r)
	data, err := ioutil.ReadAll(res)
	if err != nil {
		return err
	}

	m := map[string]string{}
	err = jsoniter.Unmarshal(data, &m)
	return err
}

func IoCopy(r io.Reader) error {
	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)

	res := Do(r)
	_, err := io.Copy(buffer, res)
	if err != nil {
		return err
	}

	return err
}

func IouitlReadAll(r io.Reader) error {
	res := Do(r)
	data, err := ioutil.ReadAll(res)
	if err != nil {
		return err
	}
	_ = data
	return err
}
