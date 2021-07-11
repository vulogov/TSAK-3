package dsl

import (
	. "github.com/glycerine/zygomys/zygo"
)

func AsAny(expr Sexp) interface{} {
	switch e := expr.(type) {
	case *SexpFloat:
		return e.Val
	case *SexpInt:
		return e.Val
	case *SexpStr:
		return e.S
	}
	return expr.SexpString(nil)
}

func ToAny(expr interface{}) Sexp {
	switch e := expr.(type) {
	case int64:
		return &SexpInt{Val: int64(e)}
	case int:
		return &SexpInt{Val: int64(e)}
	case uint64:
		return &SexpInt{Val: int64(e)}
	case float64:
		return &SexpFloat{Val: float64(e)}
	case string:
		return &SexpStr{S: string(e)}
	}
	return SexpNull
}
