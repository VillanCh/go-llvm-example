package main

import (
	"fmt"

	"tinygo.org/x/go-llvm"
)

func main() {
	// 确保 LLVM 已经正确初始化
	llvm.InitializeNativeTarget()
	llvm.InitializeNativeAsmPrinter()

	// 创建一个新的上下文
	context := llvm.NewContext()
	defer context.Dispose()

	// 创建一个模块
	module := context.NewModule("example")
	defer module.Dispose()

	// 创建一个简单的函数
	// 函数类型: int32 (int32)
	i32Type := context.Int32Type()
	funcType := llvm.FunctionType(i32Type, []llvm.Type{i32Type}, false)
	// 添加函数到模块
	function := llvm.AddFunction(module, "double", funcType)

	// 创建基本块
	block := llvm.AddBasicBlock(function, "entry")

	// 创建 IRBuilder
	builder := context.NewBuilder()
	defer builder.Dispose()
	builder.SetInsertPointAtEnd(block)

	// 获取参数
	param := function.Param(0)

	// 创建乘法指令：param * 2
	result := builder.CreateMul(param, llvm.ConstInt(i32Type, 2, false), "double")

	// 创建返回指令
	builder.CreateRet(result)

	// 验证模块
	if err := llvm.VerifyModule(module, llvm.PrintMessageAction); err != nil {
		fmt.Println("Error verifying module:", err)
		return
	}

	// 打印生成的 IR
	fmt.Println(module.String())
}
