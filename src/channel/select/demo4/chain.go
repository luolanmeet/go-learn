package main

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 1
		}
		close(out)
	}()
	return out
}

func main() {

	in := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	// 相当于每个数都自增了3
	out := chain(chain(chain(in)))

	for v := range out {
		println(v)
	}
}
