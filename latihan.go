package main
import "fmt"

func faktorialGenap(n int)int{
	if n <= 0 {
		return 1
	}
	if n%2 == 0{
		return n * faktorialGenap(n-1)
	}
	return faktorialGenap(n-1)
}

func main(){
	var masukan int
	fmt.Scan(&masukan)
	fmt.Print(faktorialGenap(masukan))
}

