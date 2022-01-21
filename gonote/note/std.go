package note

import (
	"fmt"
	"gonote/util"
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
