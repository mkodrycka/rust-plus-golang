package main
import "fmt"

// NOTE: There should be NO space between the comments and the `import "C"` line.

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

func main() {
	var x int32 = 64
	sum := C.hello(C.int(x), C.int(x))
	fmt.Print(sum);

	C.whisper(C.CString("this is code from the dynamic library"))
}
