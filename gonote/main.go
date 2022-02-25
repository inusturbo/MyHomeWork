package main

import "gonote/note"

func main() {
	//6.3 strings包的常见函数
	//note.PackageStrings()
	//6.4中文字符常见操作（UTF8包）
	//note.PackageUtf8()
	//6.5 时间常见操作
	//note.PackageTime()
	//6.6 文件常用操作
	//note.FileOperation()
	//6.7 文件读写
	//note.FileReadAndWrite()
	//6.8 错误
	//note.Errors()
	//6.9 日志log
	//note.Log()
	//6.11 命令行参数
	//note.CmdArgs()
	//6.12 builtin包
	//note.PackageBuiltin()
	//6.13 runtime包
	//note.PackageRuntime()
	//6.14 runtime包
	//note.PackageSync()

	//7.1 递归
	//note.Recursion()
	//7.2 闭包
	//note.Closure()
	//7.3 排序
	//note.Sort()
	//7.4 Sort包
	//note.PackageSort()
	//7.5 查找
	//note.BinarySearchTest()

	//8.1 工厂模式
	// m := factory.NewMes()
	// m.C = ""
	// m.SetPwd("123")

	//9.1 JSON常见操作
	// note.PackageJson()

	//10.1 tcp编程
	//note.TcpServer()
	//note.TcpCli()

	//11.1.1~3 LevelDB基本使用
	//note.LeveldbBasic()
	//11.1.4 LevelDB遍历
	//note.LeveldbIterate()
	// 11.1.5-6 LevelDB事务与快照
	//note.LeveldbTransactionAndSnapshot()
	// 11.2 Redis基本操作
	// note.RedisBasic()
	// 11.2.6 Redis管道
	// note.RedisPipeline()
	// 11.2.7 Redis事务
	//note.RedisTransaction()
	// 11.2.8 Redis遍历
	//note.RedisIterate()
	//11.2.9 将Redis Hash扫描至Go结构体RedisHashToStruct
	note.RedisHashToStruct()
}
