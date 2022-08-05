package commonerror

// CustomerError syl 自产自研究使用
type CustomerError string

func (c CustomerError) Error() string {
	return c.Msg()
}

func (c CustomerError) Msg() string {
	return string(c)
}
