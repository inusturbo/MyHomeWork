package note

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gonote/util"
	"io"
	"net"
	"sort"
)

//7.4 Sort包
type Person struct {
	Name string
	Age  int
}
type PersonSlice []Person

func (ps PersonSlice) Len() int {
	return len(ps)
}
func (ps PersonSlice) Less(i, j int) bool {
	return ps[i].Age > ps[j].Age
}
func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func PackageSort() {
	fmt.Println("\n7.4.1 对常见类型进行排序")
	is := []int{2, 4, 8, 10}
	v := 6
	i := sort.SearchInts(is, v)
	fmt.Printf("%v适合插入在%v的%v\n", v, is, i)

	fmt.Println("\n7.4.2 自定义排序")
	p := []Person{{"小小", 18}, {"小方", 5}, {"小块", 50}}
	sort.Slice(p, func(i, j int) bool {
		return p[i].Age < p[j].Age
	})
	fmt.Println("p=", p)

	fmt.Println("\n7.4.3 自定义查找")
	i = sort.Search(len(is), func(i int) bool {
		return is[i] >= v
	})
	fmt.Printf("%v中第一次出现不小于%v的位置是%v\n", is, v, i)

	fmt.Println("\n7.4.4 sort.Interface")
	sort.Sort(sort.Reverse(PersonSlice(p)))

	fmt.Println("p=", p)
}

//7.5 查找
//7.5.2 二分查找
func BinarySearch(s []int, key int) int {
	startIndex := 0
	endIndex := len(s) - 1
	midIndex := 0
	for startIndex <= endIndex {
		midIndex = startIndex + (endIndex-startIndex)/2
		if s[midIndex] < key {
			startIndex = midIndex + 1
		} else if s[midIndex] > key {
			endIndex = midIndex - 1
		} else {
			return midIndex
		}

	}
	return -1
}

func BinarySearchTest() {
	s := make([]int, util.RandInt(1000)+1)
	for i := 0; i < len(s); i++ {
		s[i] = util.RandInt(1000)
	}
	sort.Ints(s)
	i := BinarySearch(s, 555)
	if i == -1 {
		fmt.Println("没有找到")
	} else {
		fmt.Println("找到了，下标是", i)
	}
}

//9.1 JSON常见操作
func PackageJson() {
	type user struct {
		Name  string `json:"name"`          //别名
		Age   int    `json:"age,omitempty"` //若为0则忽略
		Email string `json:"-"`             //不管是什么都忽略
		Job   map[string]string
	}
	u1 := user{
		Name: "方块",
		Age:  3,
		Job: map[string]string{
			"早班": "保安",
			"午班": "洗碗",
			"晚班": "送外卖",
		},
	}
	data, _ := json.Marshal(u1)
	fmt.Println(string(data))
	buf := new(bytes.Buffer)
	json.Indent(buf, data, "", "\t")
	fmt.Println(buf.String())
	var u2 user
	json.Unmarshal(data, &u2)
	fmt.Println("u2=", u2)
}

//10.1 TCP编程入门
func TcpCli() {
	conn, err := net.Dial("tcp", "127.0.0.1:2022")
	if err != nil {
		fmt.Println("拨号失败")
		return
	}
	defer conn.Close()
	for {
		mes := struct {
			UserName string
			Mes      string
		}{
			UserName: "方块",
		}
		fmt.Println("请输入要发送的内容")
		fmt.Scanf("%s", &mes.Mes)
		if mes.Mes == "" {
			fmt.Println("输入为空！")
			continue
		}
		if mes.Mes == "exit" {
			return
		}
		// data, _ := json.Marshal(&mes)
		// n, err := conn.Write(data)
		// if err != nil {
		// 	fmt.Println("发送失败")
		// 	return
		// }
		// fmt.Printf("成功发送了%v个字节!\n", n)
		err = json.NewEncoder(conn).Encode(&mes)
		if err != nil {
			fmt.Println("发送失败")
			return
		}
	}
}
func TcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:2022")
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	defer listener.Close()
	for {
		fmt.Println("主线程等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听客户端失败")
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("一个客户端协程已开启")
			defer conn.Close()
			for {
				// buf := make([]byte, 4096)
				// n, err := conn.Read(buf)
				// if err == io.EOF {
				// 	fmt.Println("客户端退出")
				// 	return
				// }
				// if err != nil {
				// 	fmt.Println("读取失败")
				// 	return
				// }
				mes := struct {
					UserName string
					Mes      string
				}{}
				// json.Unmarshal(buf[:n], &mes)
				err := json.NewDecoder(conn).Decode(&mes)
				if err == io.EOF {
					fmt.Println("客户端退出")
					return
				}
				if err != nil {
					fmt.Println("读取失败")
					return
				}
				fmt.Printf("%s说%s\n", mes.UserName, mes.Mes)
			}
		}(conn)
	}
}
