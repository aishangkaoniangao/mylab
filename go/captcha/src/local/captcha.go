package local

import "fmt"
import "math/rand"
import "time"

var _ = fmt.Sprintf("")

func GetText(slen int) string {
	character := "abcdefghjkmnpqrstuvwxyABCDEFGHJKLMNPQRSTUVWXYZ023456789";
	maxlen := int32(len(character))

	var result string
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i < slen; i++ {
		idx := rand.Int31n(maxlen)
		result += (string)(character[idx])
	}
	return result
}
