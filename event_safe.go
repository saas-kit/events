// +build !go1.5

package events

func cloneValue(v interface{}) interface{} {
	return v
}

func bytesToString(b []byte) string {
	return string(b)
}

func noescape(a []interface{}) []interface{} {
	return a
}
