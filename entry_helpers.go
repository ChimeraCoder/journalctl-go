package journalctl

func (e Entry) Message() string {
	switch x := e.MESSAGE.(type) {
	case string:
		return x
	case []byte:
		return string(x)
	default:
		return ""
	}
}
