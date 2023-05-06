package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

// 如何使用 GoLang 执行 SHELL 命令

func createCmdAndRun() {
	// 创建 *exec.Cmd
	cmd := exec.Command("/bin/bash", "-c", "echo hello")
	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func createCmdAndCombined() {
	cmd := exec.Command("/bin/bash", "-c", "sleep 5;ls")
	// 执行命令并捕获子进程的输出
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalln(err)
	} else {
		log.Println(string(output))
	}
}

type Result struct {
	output []byte
	err    error
}

func createCmdAndKillA() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	context.TODO()
	// 🍑channel 在这里起到的作用：会阻塞协程被终止，直到 channel 拿到数据为止
	resultChan := make(chan *Result, 100)
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 3;echo hello")
		output, err := cmd.CombinedOutput()
		resultChan <- &Result{
			output: output,
			err:    err,
		}
	}()

	time.Sleep(time.Second)

	cancelFunc()

	res := <-resultChan
	fmt.Printf("res = %v, err = %v\n", string(res.output), res.err)
}

var res = &Result{}

func createCmdAndKillB() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	// 🍏此处的 res 无法取到值
	//res := &Result{}
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 3;echo hello")
		output, err := cmd.CombinedOutput()

		fmt.Println("a")
		res.output = output
		res.err = err
		fmt.Println("b")
	}()

	time.Sleep(time.Second)
	cancelFunc()
	//fmt.Printf("res = %v, err = %v\n", string(res.output), res.err)
}

func main() {
	createCmdAndCombined()
	createCmdAndKillA()
	createCmdAndKillB()

	time.Sleep(5 * time.Second)
	fmt.Printf("res = %v, err = %v\n", string(res.output), res.err)
}
