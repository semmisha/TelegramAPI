package http

type InputApi interface {
	Export() (text string, channel string)
}

func (T *Api) Export() (text string, channel string) {

	return T.Text, T.Channel

}
