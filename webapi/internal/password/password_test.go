package password

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	fmt.Println(MakePassword("123456"))
	fmt.Println(CheckPassword("123456", "sun_sha256$36000$20f52eba0766$SGdFN6Zbxj3cG80r3NOj7Ab8wWmKvmqbUw3hWPD44ZM="))
}
