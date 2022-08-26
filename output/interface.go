package output

type Output interface {
	Write([]byte) (int, error)
}
