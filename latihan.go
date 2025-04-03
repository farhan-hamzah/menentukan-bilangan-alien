package main
import "fmt"

func mtk(n int)int{
	var hasil, i, o int
	hasil = 0
	for n >= 1{
		i = n
		o = 1
		for i >= 1{
			o*= i
			i-= 1
		}
		hasil+= o
		n-= 1
	}
	return hasil
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(mtk(num))
}
