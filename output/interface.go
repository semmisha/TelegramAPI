package output

type Output interface {
	Get(string)
	Write([]byte) (int, error)
}
