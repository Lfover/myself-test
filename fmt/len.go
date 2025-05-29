package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main3() {
	p := Person{Name: "Alice", Age: 20}
	fmt.Println(p)
	bytes, _ := json.Marshal(p)
	fmt.Println("Byte count:1", len(bytes))
	fmt.Println("Character count:", len(string(bytes)))
}
func main4() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		go func() {
			fmt.Println(i, v)
		}()

	}
	time.Sleep(1 * time.Second)
}

type book struct {
}

type aa interface {
	book
}

const cl = 100

var bl = 123

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)
	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//将数组依次添加到map中
	for k, stu := range stus {
		m[stu.Name] = &stus[k]
	}
	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
func reverseWord(s []byte) string {
	slow := 0
	var q interface {
	}
	q = "q"
	fmt.Println(q)
	i, ok := q.(int)
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println(i)
	}
	fmt.Println(i)

	array := []*int{0: new(int)}
	*array[0] = 10
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if slow != 0 {
				s[slow] = ' '
				slow++
			}
			for i < len(s) && s[i] != ' ' {
				s[slow] = s[i]
				slow++
				i++
			}
		}
	}
	s = s[0:slow]
	reverse(s)
	last := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			reverse(s[last:i])
			last = i + 1
		}
	}

	return string(s)
}
func reverse(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

}
func replaceNumber(strByte []byte) string {
	// 查看有多少字符
	numCount, oldSize := 0, len(strByte)
	for i := 0; i < len(strByte); i++ {
		if (strByte[i] <= '9') && (strByte[i] >= '0') {
			numCount++
		}
	}
	// 增加长度
	for i := 0; i < numCount; i++ {
		strByte = append(strByte, []byte("     ")...)
	}
	tmpBytes := []byte("number")
	// 双指针从后遍历
	leftP, rightP := oldSize-1, len(strByte)-1
	for leftP < rightP {
		rightShift := 1
		// 如果是数字则加入number
		if (strByte[leftP] <= '9') && (strByte[leftP] >= '0') {
			for i, tmpByte := range tmpBytes {
				strByte[rightP-len(tmpBytes)+i+1] = tmpByte
			}
			rightShift = len(tmpBytes)
		} else {
			strByte[rightP] = strByte[leftP]
		}
		// 更新指针
		rightP -= rightShift
		leftP -= 1
	}
	return string(strByte)
}

func replaceNum(strByte []byte) {

	for i := 0; i < len(strByte); i++ {
		if strByte[i] <= '9' && strByte[i] >= '0' {
			inserElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strByte = append(strByte[:i], append(inserElement, strByte[i+1:]...)...)
			i = i + len(inserElement) - 1
		}
	}

	fmt.Printf(string(strByte))
}

func repalceN(s []byte) string {

	count := 0
	oldSize := len(s)
	for _, v := range s {
		if v <= '9' && v >= '0' {
			count++
		}
	}
	for i := 0; i < count; i++ {
		s = append(s, []byte("     ")...)
	}
	temp := []byte("number")
	i, j := oldSize-1, len(s)-1
	for i < j {
		rightIndx := 1
		if s[i] <= '9' && s[i] >= '0' {
			for k, v := range temp {
				s[j-len(temp)+k+1] = v
			}
			rightIndx = len(temp)
		} else {
			s[j] = s[i]
		}
		j -= rightIndx
		i -= 1
	}
	return string(s)
}

func Append(slice []byte, elems ...byte) []byte {
	// 计算新切片所需的最小容量
	newLen := len(slice) + len(elems)

	// 检查是否需要扩容
	if newLen <= cap(slice) {
		// 不需要扩容，直接在原数组上添加
		newSlice := slice[:newLen]
		copy(newSlice[len(slice):], elems)
		err := copier.Copy(newSlice, slice)
		fmt.Println(err)
		return newSlice
	}

	// 需要扩容，计算新容量
	newCap := cap(slice)
	if newCap < 1024 {
		newCap *= 2
	} else {
		newCap += newCap / 4
	}

	if newCap < newLen {
		newCap = newLen
	}

	// 分配新内存
	newSlice := make([]byte, newLen, newCap)

	// 复制原有元素
	copy(newSlice, slice)
	copier.Copy(newSlice, slice)
	// 添加新元素
	copy(newSlice[len(slice):], elems)

	return newSlice
}
