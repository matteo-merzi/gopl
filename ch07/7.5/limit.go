package limit

import "io"

type limitedReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitedReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitedReader(r io.Reader, limit int) io.Reader {
	return &limitedReader{r: r, limit: limit}
}
