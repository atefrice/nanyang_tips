package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"strconv"
)

var bHelpopt bool
var RoutingCnt int;
var FilePath string;

// 输入参数
func init() {
	flag.BoolVar(&bHelpopt, "help", false, "This Help Info")
	flag.IntVar(&RoutingCnt, "threadcnt", 100, "Create go routing count.(def: 100)");
	flag.StringVar(&FilePath, "path", "/mnt/beegfs", "set create file path")
}

func demoUsage() {
	//fmt.Printf("==%d==\n", flag.NArg())      0
	//fmt.Printf("==%d==\n", flag.NFlag())     1
	fmt.Printf(`Usage: ./$1 [-threadcnt=100] [-path==/mnt/beegfs]
Option:
`)

	flag.PrintDefaults()	
}








func createwritefile(number int, ch chan int) {

	fileContext := []byte(`It was right then that I started to think about Tomas Jefferson…… in the Decoration of Independence……in the part about our right to life, 
liberty and the pursuit of happiness.    I remember thinking, how did he know……to put the pursuit part in there? 
And maybe happiness is something that……you can only pursue.    And maybe we can actually……never have it.    No matter what, how did he know that?`)


	filename := FilePath + "/yb_go_create_file_"               
	filename += strconv.Itoa(number);
	filename += ".log";
	//fmt.Println(filename)
	//fmt.Println()
	fmt.Printf("in write file func [%d], filename[%s].\n", number, filename)
	
	err := ioutil.WriteFile(filename, fileContext, 0644);
	if (nil != err) {
		fmt.Printf("Write file(%s) Error!\n", filename)
	}

	// 该协程任务完成，通知主线程一下
	ch <- 0
}








func main() { 

	fmt.Println("Hello, Nangyang")
	
	flag.Parse()
	if (bHelpopt || (0 == flag.NFlag()) ) {
		demoUsage()
		return
	}


	fmt.Printf("Create %d thread routing.\n", RoutingCnt)

	chs := make([] chan int, RoutingCnt) 

	for i:=0; i < RoutingCnt; i++ {
		chs[i] = make(chan int)
		go createwritefile(i, chs[i])
	}

	// 进程等待所有的go routing 执行完成
	for i:=0; i < RoutingCnt; i++ {
		<-chs[i]
	}

}
