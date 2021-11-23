package main
import "fmt"
func main()  {
	XX:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break XX
			}
			fmt.Printf("%d - %d\n",i,j)
		}
	}

	fmt.Println("over")
}
