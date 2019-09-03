package concurrent

import (
	f "fmt"
	"time"
)

//  在入口类通过使用go和不使用go语句 使用并发
func Say(str string){
	for i:=0;i<5 ;i++  {
		f.Println(i,":",str)
		time.Sleep(100*time.Millisecond)
	}
}

func Concurrent(){

}