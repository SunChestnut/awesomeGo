package service

import (
	"errors"
	"fmt"
	"time"
)

/**
*
* 服务端中的方法被客户端的调用方式为：Service.Method
*
* @author  sun
* @date 2022/11/8 15:34
 */

type WorkerService struct{}

// HealthCheck ==> 健康检查
func (WorkerService) HealthCheck(msg string, result *Result) error {
	fmt.Printf("🍎WorkerService: Calling HealthCheck Method\n")
	if len(msg) == 0 {
		return errors.New("empty message")
	}
	result.Msg = msg
	result.StatusCode = 200
	result.ReplyDate = time.Now()
	return nil
}

// ClimbStairs  ==> 共 n 级台阶，每次爬 1 个或者 2 个台阶，计算有多少种方法能爬到楼顶
func (WorkerService) ClimbStairs(param Request, result *Result) error {
	memory := []int{1, 1}
	if param.Num <= 1 {
		result.T = memory[param.Num]

	}
	sum := 0
	for i := 2; i <= param.Num; i++ {
		sum = memory[i-1] + memory[i-2]
		memory = append(memory, sum)
	}

	// todo 可否用生成器模式改造下列代码？
	result.T = memory[param.Num]
	result.Msg = "success"
	result.StatusCode = 200
	result.ReplyDate = time.Now()
	return nil
}
