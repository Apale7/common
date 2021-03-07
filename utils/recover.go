package utils

import (
	"fmt"
	"runtime"
)

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		if err != nil {
			switch err.(type) {
			case runtime.Error: // 运行时错误
				fmt.Println("recover from panic: runtime error: ", err)
			default: // 非运行时错误
				fmt.Println("recover from panic: error:", err)
			}
		}
	}()
	entry()
}
