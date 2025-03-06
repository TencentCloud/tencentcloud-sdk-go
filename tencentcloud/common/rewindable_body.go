package common

import (
	"bytes"
	"io"
	"net/http"
)

func newRequestRetryer(r *http.Request) func() error {
	if r.GetBody == nil {
		rb := &rewindableBody{body: r.Body}
		r.Body = rb
		r.GetBody = func() (io.ReadCloser, error) {
			err := rb.Rewind()
			return rb, err
		}
	}

	return func() error {
		bodyCopy, err := r.GetBody()
		if err != nil {
			return err
		}

		r.Body = bodyCopy
		return nil
	}
}

type rewindableBody struct {
	body   io.ReadCloser
	buf    []byte
	r      bytes.Reader
	rewind bool
}

func (r *rewindableBody) Read(p []byte) (int, error) {
	if r.rewind {
		return r.r.Read(p)
	}

	nr, err := r.body.Read(p)
	if nr > 0 {
		r.buf = append(r.buf, p[:nr]...)
	}
	return nr, err
}

func (r *rewindableBody) Close() error {
	return r.body.Close()
}

func (r *rewindableBody) Rewind() error {
	if !r.rewind {
		// drain body
		buf := bytes.NewBuffer(r.buf)
		_, err := io.Copy(buf, r.body)
		if err != nil {
			return err
		}
		r.buf = buf.Bytes()
	}

	r.r.Reset(r.buf)
	r.rewind = true
	return nil
}
