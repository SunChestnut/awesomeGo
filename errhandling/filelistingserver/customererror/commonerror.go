package commonerror

// CommonError syl 自产自研究使用
type CommonError interface {
	error
	Msg() string
}
