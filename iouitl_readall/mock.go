package iouitl_readall

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

var response []byte

func init() {
	m := make(map[string]string, 150)

	for i := 0; i < 150; i++ {
		m["X"+strconv.Itoa(i)] = strings.Repeat("A", i/2)
	}

	var err error
	response, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
}

func Do(r io.Reader) io.Reader {
	return bytes.NewReader(response)
}
