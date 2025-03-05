package common

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"testing"
)

func randBytes(length int) (p []byte) {
	for i := 0; i < length; i++ {
		p = append(p, byte(rand.Int()%256))
	}
	return p
}

func TestRewindableBody_Read(t *testing.T) {
	body := randBytes(12345)
	rb := &rewindableBody{body: ioutil.NopCloser(bytes.NewReader(body))}

	readBody, err := ioutil.ReadAll(rb)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, readBody) != 0 {
		t.FailNow()
	}
}

func TestRewindableBody_Close(t *testing.T) {
	body := randBytes(12345)
	rb := &rewindableBody{body: ioutil.NopCloser(bytes.NewReader(body))}

	err := rb.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRewindableBody_Rewind(t *testing.T) {
	body := randBytes(12345)
	rb := &rewindableBody{body: ioutil.NopCloser(bytes.NewReader(body))}

	readBody, err := ioutil.ReadAll(rb)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, readBody) != 0 {
		t.FailNow()
	}

	err = rb.Rewind()
	if err != nil {
		t.Fatal(err)
	}

	readBody, err = ioutil.ReadAll(rb)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, readBody) != 0 {
		t.FailNow()
	}
}

func TestRewindableBody_RewindNoRead(t *testing.T) {
	body := randBytes(12345)
	rb := &rewindableBody{body: ioutil.NopCloser(bytes.NewReader(body))}

	err := rb.Rewind()
	if err != nil {
		t.Fatal(err)
	}

	readBody, err := ioutil.ReadAll(rb)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, readBody) != 0 {
		t.FailNow()
	}
}

func TestRewindableBody_RewindPartialRead(t *testing.T) {
	body := randBytes(12345)
	rb := &rewindableBody{body: ioutil.NopCloser(bytes.NewReader(body))}

	tmpBuffer := make([]byte, 1024)
	n, err := rb.Read(tmpBuffer)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(tmpBuffer) {
		t.FailNow()
	}

	err = rb.Rewind()
	if err != nil {
		t.Fatal(err)
	}

	readBody, err := ioutil.ReadAll(rb)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, readBody) != 0 {
		t.FailNow()
	}
}
