package sub2

import "strconv"
import "fmt"
import "strings"

func Bar() {
	query1 := "HelloYeah"
	length := strings.Count(query1, "") - 1
	fmt.Println("Query length: " + strconv.Itoa(length))
}
