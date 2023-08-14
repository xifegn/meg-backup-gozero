package cryptx

import (
	"fmt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	fmt.Println(PasswordEncrypt("HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe", "123456"))
}
