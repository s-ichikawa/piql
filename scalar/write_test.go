package scalar_test

type TestWriter struct {
	V interface{}
}

func (w *TestWriter) Write(p []byte) (n int, err error) {
	w.V = string(p)
	return 0, nil
}
