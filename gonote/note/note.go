package note

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gonote/util"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

//6.3 strings包的常见函数
func PackageStrings() {
	fmt.Println(strings.Contains("Hello", "o"))
	fmt.Println(strings.Index("Hello", "ll"))
	fmt.Println(strings.Replace("hello", "l", "o", -1)) //最后的数字是替换次数，-1表示全部替换
	fmt.Println(strings.Repeat("Hello", 5))
	fmt.Println(strings.Fields("mia mia\n mia\tmia"))
	fmt.Println(strings.Split("he-llo-wor-ld", "-"))
	fmt.Println(strings.Trim("#*\nwww.www.www&%#", "#*%&\n"))
}

//6.4中文字符常见操作（UTF8包）
func PackageUtf8() {
	str := "hello, 世界"
	fmt.Println(utf8.RuneCountInString("hello, 世界"))
	fmt.Println(utf8.ValidString(str[:len(str)-1]))
	fmt.Println(utf8.ValidString(str))
}

//6.5 时间常见操作
func PackageTime() {
	fmt.Println("\n6.5.1 时段")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println()
	d1, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println("d1=", d1)
	t1, err := time.Parse("2006年1月2日，15点4分", "2022年1月1日，18点18分")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t1))
	t2, err := time.Parse("2006年1月2日，15点4分", "2023年1月1日，18点18分")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t2))
	var intChan chan int = make(chan int)
	select {
	case <-intChan:
		fmt.Println("收到了用户发送的验证码")
	case <-time.After(1 * time.Second):
		fmt.Println("验证码已过期")
	}

	fmt.Println("\n6.5.2 时区")
	l1, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	println(l1.String())

	fmt.Println("\n6.5.3 时刻")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006年1月2日，15点4分"))
	t3, err := time.ParseInLocation("2006年1月2日，15点4分", "2100年12月23日，15点14分", l1)
	if err != nil {
		panic(err)
	}
	fmt.Println(t3)
	fmt.Println(t3.Add(d1))

	fmt.Println("\n6.5.4 周期计时器")
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
TickerFor:
	for {
		select {
		case <-intChan:
			fmt.Println()
			break TickerFor
		case <-time.NewTicker(100 * time.Millisecond).C:
			fmt.Print(".")
		}
	}
	fmt.Println("\n6.5.7 单次计时器")
	intChan = make(chan int)
	select {
	case <-intChan:
		fmt.Println("收到了用户发送的验证码")
	case <-time.NewTimer(time.Second).C:
		fmt.Println("验证码已过期")
	}
}

//6.6 文件常用操作
func FileOperation() {
	//util.MkdirWithFilePath("d1/d2/fil2")
	fmt.Println("6.6.5 文件夹操作")
	DirEntrys, err := os.ReadDir("/Users/mashanpeng")
	if err != nil {
		panic(err)
	}
	for _, v := range DirEntrys {
		fmt.Println(v.Name())
	}
	fmt.Println("6.6.6 文件操作")
	file, err := os.OpenFile("f1", os.O_RDWR|os.O_CREATE, 0665)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("6.6.7 无缓冲区读写（适合小文件）")
	data, err := os.ReadFile("f1")
	if err != nil {
		panic(err)
	}
	fmt.Println("f1中数据为", string(data))
	err = os.WriteFile("f2", data, 0775)
	if err != nil {
		panic(err)
	}
}

//6.7 文件读写
func FileReadAndWrite() {
	f5, err := os.OpenFile("f5", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f5.Close()
	writer := bufio.NewWriter(f5)
	for i := 1; i <= 4; i++ {
		fileName := fmt.Sprintf("f%v", i)
		data, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		data = append(data, '\n')
		writer.Write(data) //写入缓冲区
	}
	writer.Flush()
}

//6.8 错误
func Errors() {
	defer func() {
		err := recover()
		fmt.Println("捕捉到了错误", err)
	}()
	err1 := errors.New("可爱的错误")
	fmt.Println("err1=", err1)
	err2 := fmt.Errorf("%s的错误", "温柔")
	fmt.Println("err2=", err2)
	panic(err1)
}

//6.9 日志log
func Log() {
	err := errors.New("可爱的错误")
	util.INFO.Println(err)
	util.WARN.Panic(err)
	util.ERR.Fatalln(err)
}

//6.10 单元测试
func IsNotNegative(n int) bool {
	return n > -1
}

//6.11 命令行参数
func CmdArgs() {
	fmt.Printf("接收到了%v个参数\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v个参数是%v\n", i, v)
	}
	fmt.Println()
	vPtr := flag.Bool("v", false, "GoNote版本号")
	var userName string
	flag.StringVar(&userName, "u", "", "用户名")
	flag.Func("f", "", func(s string) error {
		fmt.Println("s=", s)
		return nil
	})
	flag.Parse()
	if *vPtr {
		fmt.Println("GoNote版本是v0.0.0")
	}
	fmt.Println("当前用户：", userName)
	for i, v := range flag.Args() {
		fmt.Printf("第%v个无flag参数是%v\n", i, v)
	}
}

//6.12 builtin包
func PackageBuiltin() {
	c1 := complex(12.34, 45.67)
	fmt.Println("c1=", c1)
	r1 := real(c1)
	i1 := imag(c1)
	fmt.Println("r1=", r1)
	fmt.Println("i1=", i1)
}

//6.13 runtime包
func PackageRuntime() {
	if runtime.NumCPU() > 7 {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}
	//runtime.Goexit()

}

//6.14 sync包（同步）
func PackageSync() {
	fmt.Println("\n6.14.1 Mutex互斥锁 | 6.14.2 WaitGroup")
	var c int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	primeNum := func(n int) {
		defer wg.Done()
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return
			}
		}
		mutex.Lock()
		c++
		mutex.Unlock()
	}

	for i := 2; i < 100001; i++ {
		wg.Add(1)
		go primeNum(i)
	}
	wg.Wait()
	fmt.Printf("\n共找到%v个素数\n", c)

	fmt.Println("\n6.14.3 Cond")
	cond := sync.NewCond(&mutex)
	for i := 0; i < 10; i++ {
		go func(n int) {
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("协程%v被唤醒了\n", n)
			cond.L.Unlock()
		}(i)
	}
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
			cond.Signal()
		}
		if i == 10 {
			fmt.Println()
			cond.Broadcast()
		}
	}
	fmt.Println("\n6.14.4 Once")
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				fmt.Println("只有一次机会")
			})
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("\n6.14.5 Map")
	var m sync.Map
	m.Store(1, 100)
	m.Store(2, 300)
	m.Store(2, 300)
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("m[%v]=%v\n", key, value.(int))
		return true
	})
}
