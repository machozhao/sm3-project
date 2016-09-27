package sm3

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
// intentionally write the same LDFLAGS differently
#cgo linux CFLAGS: -I/Users/gerryyang/github_project/goinaction/src/cgo/inc
#cgo linux LDFLAGS: -L../lib -lhello
#cgo darwin CFLAGS: -I/Users/gerryyang/github_project/goinaction/src/cgo/inc
#cgo darwin LDFLAGS: -L/usr/lib
#if 0
void hello(const char *name)
{
	printf("%s\n", name);
	return;
}
#endif
*/
import "C"

import (
"fmt"
"time"
)

func MyMain() {

	Seed(1000)

	fmt.Println(int(C.random()))
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(int(C.random()))

	fmt.Println("getpid:", int(C.getpid()))
	C.puts(C.CString("call C puts"))

}

func Seed(i int) {
	C.srandom(C.uint(i))
}