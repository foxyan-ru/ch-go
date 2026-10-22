//go:build growfast

package compress

// growBuf resizes buf to n bytes.
//
// This is the variant under test: it reuses the existing capacity and does not
// re-zero the region, because every byte of the region is overwritten by the
// caller before it is read.
func growBuf(buf []byte, n int) []byte {
	if cap(buf) < n {
		return make([]byte, n)
	}
	return buf[:n]
}
