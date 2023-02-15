package logwarts

import (
	"fmt"
	"runtime"
	"strings"
)

func stack() string {
	r := "  stack"
	i := 3
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.HasPrefix(file, runtime.GOROOT()) {
			break
		}
		f := runtime.FuncForPC(pc)
		r += fmt.Sprintf("\n    %s:%d:%s()", file, line, f.Name())
		i++
	}
	return r
}
