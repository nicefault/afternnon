package afternoon

import "fmt"

func GoodAfternoon(name string) string {
	message := fmt.Sprintf("Hey %v, Good afternoon!", name)
	return message
}
