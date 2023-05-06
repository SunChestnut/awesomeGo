package service

import "time"

/**
*
* TODO
*
* @author  sun
* @date 2022/11/8 15:58
 */

type Request struct {
	Msg string
	Num int
}

type Result struct {
	T          any
	Msg        string
	StatusCode int
	ReplyDate  time.Time
}
