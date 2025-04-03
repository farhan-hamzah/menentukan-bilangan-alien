package main
import (
		"fmt"
		"math"
)

func bilAlien(n int)int{
	var basis, baris, i, hasil int
	var j, k int
	hasil = 0
	baris = n
	for baris > 0{
		baris = baris/10
		basis++
	}
	basis = basis-1
	for basis >= 0{
		i = n%10
		j = int(math.Pow(7, float64(basis)))
		k = j*i
		hasil = hasil+k
		n = n/10
		basis-=1
		k = 1
	}
	return hasil
}
func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(bilAlien(num))
}
