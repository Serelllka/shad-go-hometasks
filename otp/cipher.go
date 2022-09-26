package otp

import (
	"io"
)

type encryptedReader struct {
	reader, cipher io.Reader
}

type encryptedWriter struct {
	writer io.Writer
	cipher io.Reader
}

func (er *encryptedReader) Read(p []byte) (n int, err error) {
	inp := make([]byte, 1)
	cipher := make([]byte, 1)
	var k int
	for i := 0; i < len(p); i++ {
		for k == 0 {
			k, err = er.cipher.Read(cipher)
			if err != nil {
				return i, err
			}
		}

		k = 0
		for k == 0 {
			k, err = er.reader.Read(inp)

			if err != nil {
				if err == io.EOF {
					p[i] = inp[0] ^ cipher[0]
					return i + 1, err
				} else {
					return i, err
				}
			}
		}

		p[i] = inp[0] ^ cipher[0]
	}
	return len(p), nil
}

func (er *encryptedWriter) Write(p []byte) (n int, err error) {
	cipher := make([]byte, 1)
	toWrite := make([]byte, 1)
	for i := 0; i < len(p); i++ {
		_, err = er.cipher.Read(cipher)
		if err == io.EOF {
			return i, err
		} else if err != nil {
			return 0, err
		}
		toWrite[0] = cipher[0] ^ p[i]
		_, err = er.writer.Write(toWrite)
		if err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	return &encryptedReader{
		reader: r,
		cipher: prng,
	}
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	return &encryptedWriter{
		writer: w,
		cipher: prng,
	}
}
