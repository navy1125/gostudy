package main
import "fmt"
var cmap map[int]chan bool
func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
			fmt.Println(len(c))
		case <-quit:
			fmt.Println("quit")
			return
		default:
			break
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(len(c),<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
