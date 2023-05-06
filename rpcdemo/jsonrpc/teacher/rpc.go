package teacher

import "errors"

// DemoService  ==> 服务端中的方法被客户端的调用方式为：Service.Method
type DemoService struct{}

type Args struct {
	A, B int
}

// Div  ==> 被 RPC Client 端调用的方法格式是有要求的，第一个参数，第二个结果值，然后返回 error
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
