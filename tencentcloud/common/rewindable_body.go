package common

import (
	"bytes"
	"io"
	"net/http"
)

// newRequestRetryer creates a function to reset the request body for retries.
// It handles cases where the request body needs to be read multiple times.
func newRequestRetryer(r *http.Request) func() error {
	// If the request doesn't have a GetBody function, we need to add one.
	if r.GetBody == nil {
		// Wrap the original body in a rewindableBody to allow reading it multiple times.
		rb := &rewindableBody{body: r.Body}
		r.Body = rb
		// Set GetBody to a function that rewinds the rewindableBody.
		r.GetBody = func() (io.ReadCloser, error) {
			err := rb.Rewind() // Prepare the body for reading again.
			return rb, err
		}
	}

	// Return a closure that resets the request's Body to a new reader.
	return func() error {
		bodyCopy, err := r.GetBody() // Get a new ReadCloser for the body.
		if err != nil {
			return err
		}

		r.Body = bodyCopy // Set the request's Body to the new reader.
		return nil
	}
}

// rewindableBody allows an io.ReadCloser to be read multiple times.
// It buffers the body content to enable rewinding.
type rewindableBody struct {
	body   io.ReadCloser // The original request body.
	buf    []byte        // Buffer to store the body content.
	r      bytes.Reader  // Reader to read from the buffer.
	rewind bool          // Flag indicating if the body has been rewound.
}

// Read implements the io.Reader interface.
// If rewind is true, it reads from the buffer; otherwise, it reads from the original body
// and buffers the data.
func (r *rewindableBody) Read(p []byte) (int, error) {
	if r.rewind {
		return r.r.Read(p) // Read from the buffer.
	}

	nr, err := r.body.Read(p) // Read from the original body.
	if nr > 0 {
		r.buf = append(r.buf, p[:nr]...) // Buffer the read data.
	}
	return nr, err
}

// Close implements the io.Closer interface.
// It closes the original request body.
func (r *rewindableBody) Close() error {
	return r.body.Close()
}

// Rewind prepares the body for rereading.
// If the body has not been rewound, it drains and buffers the original body.
// It then resets the buffer reader and sets the rewind flag.
func (r *rewindableBody) Rewind() error {
	if !r.rewind {
		// Drain and buffer the original body.
		buf := bytes.NewBuffer(r.buf)
		_, err := io.Copy(buf, r.body)
		if err != nil {
			return err
		}
		r.buf = buf.Bytes() // Store the buffered data.
	}

	r.r.Reset(r.buf) // Reset the buffer reader.
	r.rewind = true  // Set the rewind flag.
	return nil
}
