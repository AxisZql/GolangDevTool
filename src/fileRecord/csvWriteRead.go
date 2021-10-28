package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Post 定义写入和读取的存储数据的结构体
type Post struct {
	Id      int
	Content string
	Author  string
}

func WriteAndReadCSV() {
	//定义所有写入的数据
	allPosts := []Post{
		{Id: 1, Content: "fine", Author: "cat"},
		{Id: 2, Content: "ok", Author: "dog"},
		{Id: 3, Content: "stupid", Author: "duck"},
	}
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
    defer csvFile.Close()
	//写入cvs文件部分
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line) //将一个数据写入一行
		if err != nil{
			panic(err)
		}
	}
	writer.Flush() //为了防止数据还在缓存中停留，故需要刷新

	file,err:=os.Open("posts.csv")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//读取CVS文件部分
	reader:=csv.NewReader(file)
	reader.FieldsPerRecord=-1//这个值是用户要求从每个数据中读取的字段数量，如果设置的不是-1或者0，则在读取字
	// 段少于当前设定的值的时候会抛出异常，如果是-1则没有异常，是0会把第一行的字段数目作为FieldsPerRecord的值

	record,err:=reader.ReadAll()//读取全部行数
	if err!=nil{
		panic(err)
	}
	var posts []Post//声明并定义存储读取数据的数据结构数组
	for _,item:=range record{
		id,_:=strconv.ParseInt(item[0],0,0)//将字符串转换为int类型
		post:=Post{Id: int(id),Content: item[1],Author: item[2]}
		posts=append(posts,post)

		fmt.Println(posts[0].Id)
		fmt.Println(posts[0].Content)
		fmt.Println(posts[0].Author)
	}

}
