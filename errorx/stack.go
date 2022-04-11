package errorx

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/gookit/goutil/mathutil"
)

// stack represents a stack of program counters.
type stack []uintptr

// Format stack trace
func (s *stack) Format(fs fmt.State, verb rune) {
	switch verb {
	case 'v', 's':
		_, _ = s.WriteTo(fs)
	}
}

// StackLen for error
func (s *stack) StackLen() int {
	return len(*s)
}

// WriteTo for error
func (s *stack) WriteTo(w io.Writer) (int64, error) {
	if len(*s) == 0 {
		return 0, nil
	}

	nn, _ := w.Write([]byte("\nSTACK:\n"))
	for _, pc := range *s {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		f := runtime.FuncForPC(pc - 1)
		if f == nil {
			continue
		}

		file, line := f.FileLine(pc - 1)
		pos := f.Name() + "()\n  " + file + ":" + mathutil.String(line) + "\n"

		n, _ := w.Write([]byte(pos))
		nn += n
	}

	return int64(nn), nil
}

// String format to string
func (s *stack) String() string {
	var sb *strings.Builder
	_, _ = s.WriteTo(sb)
	return sb.String()
}

// StackTrace frame list
func (s *stack) StackTrace() *runtime.Frames {
	return runtime.CallersFrames(*s)
}

// Location for error report
func (s *stack) Location() (file string, line int) {
	if len(*s) > 0 {
		// For historical reasons if pc is interpreted as a uintptr
		// its value represents the program counter + 1.
		pc := (*s)[0] - 1
		f := runtime.FuncForPC(pc)

		if f != nil {
			return f.FileLine(pc)
		}
	}

	return "", 0
}

/*************************************************************
 * helper func for callers stacks
 *************************************************************/

// ErrStackOpt struct
type ErrStackOpt struct {
	SkipDepth  int
	TraceDepth int
}

// default option
var stdOpt = newErrOpt()

func newErrOpt() *ErrStackOpt {
	return &ErrStackOpt{
		SkipDepth:  3,
		TraceDepth: 20,
	}
}

// Config the stdOpt setting
func Config(fns ...func(opt *ErrStackOpt)) {
	for _, fn := range fns {
		fn(stdOpt)
	}
}

// SkipDepth setting
func SkipDepth(skipDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.SkipDepth = skipDepth
	}
}

// TraceDepth setting
func TraceDepth(traceDepth int) func(opt *ErrStackOpt) {
	return func(opt *ErrStackOpt) {
		opt.TraceDepth = traceDepth
	}
}

func callersStack(skip, depth int) *stack {
	pcs := make([]uintptr, depth)
	num := runtime.Callers(skip, pcs[:])

	var st stack = pcs[0:num]
	return &st
}