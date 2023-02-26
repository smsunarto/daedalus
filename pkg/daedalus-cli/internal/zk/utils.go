package zk

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Gnark objects have a serializer that we can access
// using the io.WriterTo and io.ReaderFrom interfaces.
// See: https://docs.gnark.consensys.net/HowTo/serialize
type GnarkSerializableObj interface {
	io.WriterTo
	io.ReaderFrom
}

func WriteGnarkObjBinary(obj GnarkSerializableObj, filepath string) error {
	buf := new(bytes.Buffer)
	if _, err := obj.WriteTo(buf); err != nil {
		return err
	}

	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	w := bufio.NewWriter(fo)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func LoadGnarkObjBinary(buf *bytes.Buffer, filepath string) error {
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}

	return nil
}
