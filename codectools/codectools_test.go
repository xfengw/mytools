package codectools

import (
	"fmt"
	"testing"
)

func TestCodeConvertByEncode(t *testing.T) {
	str := GenerateRandomString(32)
	fmt.Println(str)
}
