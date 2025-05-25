package iteration

//import "fmt"

func repeat(letter string) string {
	res := ""
	for i := 0; i < 5; i++ {
		res += letter
	}
	return res
}
