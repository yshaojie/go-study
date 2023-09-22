package main

//
//import (
//	"context"
//	"fmt"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//func main() {
//
//	//创建监听退出chan
//	si := make(chan os.Signal)
//	//监听指定信号 ctrl+c kill
//	signal.Notify(si, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
//		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
//	ctx, cancel := context.WithCancel(context.Background())
//	go func() {
//		for s := range si {
//			switch s {
//			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
//				fmt.Println("Program Exit...", s)
//				cancel()
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				println("abd")
//				go func() {
//					time.Sleep(10 * time.Second)
//					println("结束了")
//				}()
//				println("aaaaaaaaaaaaaaa")
//			case syscall.SIGUSR1:
//				fmt.Println("usr1 signal", s)
//			case syscall.SIGUSR2:
//				fmt.Println("usr2 signal", s)
//			default:
//				fmt.Println("other signal", s)
//			}
//		}
//	}()
//
//	gen(ctx)
//	//go func() {
//	//	time.AfterFunc(10*time.Second, func() {
//	//		cancelFunc()
//	//		println("ddddddddddddd")
//	//	})
//	//}()
//	for range gen(ctx) {
//		println(time.Now().Format("YYYY-MM-dd HH:mm:ss"))
//	}
//
//}
//
//func gen(ctx context.Context) chan struct{} {
//	c := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				println("close....")
//				close(c)
//				return
//			default:
//				c <- struct{}{}
//			}
//			time.Sleep(5 * time.Second)
//		}
//
//	}()
//	return c
//}
