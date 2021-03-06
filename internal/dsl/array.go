package dsl

import (
	"fmt"

	. "github.com/glycerine/zygomys/zygo"
)

func ArrayofStringsToArray(expr Sexp) (res []string) {
	switch e := expr.(type) {
	case *SexpArray:
		res = make([]string, len(e.Val))
		for n, v := range e.Val {
			switch z := v.(type) {
			case *SexpStr:
				res[n] = z.S
			default:
				res[n] = v.SexpString(nil)
			}
		}
		return
	}
	res = make([]string, 0)
	return
}

func ArrayofFloatsToArray(expr Sexp) (res []float64) {
	switch e := expr.(type) {
	case *SexpArray:
		res = make([]float64, len(e.Val))
		for n, v := range e.Val {
			switch z := v.(type) {
			case *SexpFloat:
				res[n] = float64(z.Val)
			case *SexpInt:
				res[n] = float64(z.Val)
			default:
				res = make([]float64, 0)
				return
			}
		}
		return
	}
	res = make([]float64, 0)
	return
}

func ArrayofSomethingToArray(expr Sexp) (res []interface{}) {
	switch e := expr.(type) {
	case *SexpArray:
		res = make([]interface{}, len(e.Val))
		for n, v := range e.Val {
			switch z := v.(type) {
			case *SexpFloat:
				res[n] = float64(z.Val)
			case *SexpInt:
				res[n] = float64(z.Val)
			case *SexpStr:
				res[n] = string(z.S)
			default:
				res[n] = z.SexpString(nil)
			}
		}
		return
	}
	res = make([]interface{}, 0)
	return
}

func ArrayofSomethingToLispArray(env *Zlisp, arr []interface{}) Sexp {
	res := &SexpArray{Val: make([]Sexp, 0), Env: env}
	for _, v := range arr {
		switch e := v.(type) {
		case int:
			res.Val = append(res.Val, &SexpInt{Val: int64(e)})
		case int64:
			res.Val = append(res.Val, &SexpInt{Val: int64(e)})
		case float64:
			res.Val = append(res.Val, &SexpFloat{Val: float64(e)})
		case string:
			res.Val = append(res.Val, &SexpStr{S: string(e)})
		default:
			res.Val = append(res.Val, &SexpStr{S: fmt.Sprintf("%v", e)})
		}
	}
	return res
}

func ArrayofFloatsToLispArray(env *Zlisp, arr []float64) Sexp {
	res := &SexpArray{Val: make([]Sexp, 0), Env: env, Typ: Float64RT}
	for _, v := range arr {
		res.Val = append(res.Val, &SexpFloat{Val: float64(v)})
	}
	return res
}

func ArrayofFloatsToIntLispArray(env *Zlisp, arr []float64) Sexp {
	res := &SexpArray{Val: make([]Sexp, 0), Env: env, Typ: Int64RT}
	for _, v := range arr {
		res.Val = append(res.Val, &SexpInt{Val: int64(v)})
	}
	return res
}

func ArrayofIntToFloatLispArray(env *Zlisp, arr []int) Sexp {
	res := &SexpArray{Val: make([]Sexp, 0), Env: env, Typ: Float64RT}
	for _, v := range arr {
		res.Val = append(res.Val, &SexpFloat{Val: float64(v)})
	}
	return res
}

func ArrayofFloatsToFloatLispArray(env *Zlisp, arr []float64) Sexp {
	res := &SexpArray{Val: make([]Sexp, 0), Env: env, Typ: Float64RT}
	for _, v := range arr {
		res.Val = append(res.Val, &SexpFloat{Val: float64(v)})
	}
	return res
}

func ArrayOfFloatsMulOn(arr []float64, m float64) []float64 {
	for n, v := range arr {
		arr[n] = v * m
	}
	return arr
}
