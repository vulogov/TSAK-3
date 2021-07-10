package dsl

import (
	"fmt"

	. "github.com/glycerine/zygomys/zygo"
	floats "gonum.org/v1/gonum/floats"
	stat "gonum.org/v1/gonum/stat"
)

func NumNorm(data []float64) []float64 {
	res := make([]float64, len(data))
	xmin := floats.Min(data)
	xmax := floats.Max(data)
	diff := xmax - xmin
	if diff == 0 {
		for i := 0; i < len(data); i++ {
			res[i] = 0.0
		}
	} else {
		for i := 0; i < len(data); i++ {
			res[i] = (data[i] - xmin) / diff
		}
	}
	return res
}

func NumStand(data []float64) []float64 {
	xmean := stat.Mean(data, nil)
	xdev := stat.StdDev(data, nil)
	res := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = (data[i] - xmean) / xdev
	}
	return res
}

func NormalizeAll(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	arr := make([]float64, 0)
	if len(args) != 1 {
		return SexpNull, WrongNargs
	}
	switch e := args[0].(type) {
	case *SexpArray:
		arr = ArrayofFloatsToArray(e)
	default:
		return SexpNull, fmt.Errorf("First argument must be array")
	}
	switch name {
	case "normalize.Normalize":
		res := NumNorm(arr)
		return ArrayofFloatsToFloatLispArray(env, res), nil
	case "normalize.Standard":
		res := NumStand(arr)
		return ArrayofFloatsToFloatLispArray(env, res), nil
	}
	return SexpNull, fmt.Errorf("Requested normalization computation can not be performed: %v", name)
}

func NormalizeFunctions() map[string]ZlispUserFunction {
	return map[string]ZlispUserFunction{
		"normalizen": NormalizeAll,
		"normalizes": NormalizeAll,
	}
}

func NormalizePackageSetup(cfg *ZlispConfig, env *Zlisp) {
	myPkg := `(def normalize (package "normalize"
     { Normalize := normalizen;
			 Standard := normalizes;
     }
  ))`
	_, err := env.EvalString(myPkg)
	PanicOn(err)
}
