package storage

import "io"

// Write для записи
func Write(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

// Read для чтения
func Read(r io.Reader) ([]byte, error) {
	buf := make([]byte, 1024)
	var b []byte
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if n > 0 {
			b = append(b, buf[:n]...)
		}
	}
	return b, nil
}
