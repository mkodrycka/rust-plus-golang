package main
import "fmt"

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./lib/libhello.a -ldl
#include "./lib/hello.h"
*/
import "C"

func main() {
	var x int32 = 64
	sum := C.hello(C.int(x), C.int(x))
	fmt.Print(sum);
	C.whisper(C.CString("this is code from the static library"))
}
