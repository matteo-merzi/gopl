package tempflag

import (
	"flag"
	"fmt"
	"testing"
)

func TestCelsiusFlag(t *testing.T) {
	temp := CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}
