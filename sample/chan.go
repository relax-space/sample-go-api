package sample

import "fmt"

func ChanRight() {
	param := InitParam()
	fruits := InitData()
	Right(param, fruits)
	fmt.Println("done!")
}

func ChanWrong() {
	param := InitParam()
	fruits := InitData()
	Wrong(param, fruits)
	fmt.Println("done!")
}

func Right(param *Param, fruits []*Fruit) {
	asyncChan := make(chan struct{})
	for _, fruit := range fruits {
		fruit := *fruit
		param := *param
		param.IsFresh = fruit.IsFresh
		go doIt(&param, &fruit, asyncChan)
	}
	for index := 0; index < len(fruits); index++ {
		<-asyncChan
	}
}

func Wrong(param *Param, fruits []*Fruit) {
	asyncChan := make(chan struct{})
	for _, fruit := range fruits {
		param.IsFresh = fruit.IsFresh
		go doIt(param, fruit, asyncChan)
	}
	for index := 0; index < len(fruits); index++ {
		<-asyncChan
	}
}

func doIt(param *Param, fruit *Fruit, asyncChan chan<- struct{}) {
	if param.IsFresh != fruit.IsFresh {
		fmt.Printf("param:%v,fruit:%v \n", param.IsFresh, fruit.IsFresh)
	}
	asyncChan <- struct{}{}
}

func InitParam() (param *Param) {
	param = &Param{IsFresh: false}
	return
}

func InitData() (fruits []*Fruit) {
	for index := 0; index < 50; index++ {
		fruits = append(fruits, &Fruit{Name: "apple", IsFresh: false})
	}
	for index := 0; index < 50; index++ {
		fruits = append(fruits, &Fruit{Name: "apple", IsFresh: true})
	}
	for index := 0; index < 100; index++ {
		fruits = append(fruits, &Fruit{Name: "pear", IsFresh: false})
	}
	for index := 0; index < 100; index++ {
		fruits = append(fruits, &Fruit{Name: "pear", IsFresh: true})
	}
	return
}

type Param struct {
	IsFresh bool
}
type Fruit struct {
	Name    string
	IsFresh bool
}
