//go:build !growfast

package compress

// growBuf resizes buf to n bytes.
//
// This is the baseline variant: it matches the idiom used by the package before
// the experiment, append(buf[:0], make([]byte, n)...), which the Go compiler
// optimizes into a slice extension. The appended region is zeroed.
func growBuf(buf []byte, n int) []byte {
	return append(buf[:0], make([]byte, n)...)
}
