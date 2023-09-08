package compress

import (
	"bytes"

	"github.com/klauspost/compress/zstd"
)

type ZST struct{}

// Initialize a new ZSTD object based on the Interface interface
func NewZST() Interface {
	return &ZST{}
}

// Encode compresses the given bytes using ZSTD compression,
// returning the compressed data in a new bytes.Buffer.
func (zs *ZST) Encode(v []byte) (*bytes.Buffer, error) {
	var i bytes.Buffer
	w, _ := zstd.NewWriter(&i)
	if _, err := w.Write(v); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return &i, nil
}

// The Decode method will first decode and then overwrite the data in the input *bytes.Buffer.
func (zs *ZST) Decode(v *bytes.Buffer) {
	r, _ := zstd.NewReader(v)
	d := reader(r)
	r.Close()
	v.Reset()
	_, _ = v.Write(d)
}
