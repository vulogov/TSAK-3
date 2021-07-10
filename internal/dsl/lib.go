package dsl

import (
	"github.com/glycerine/zygomys/zygo"
	. "github.com/glycerine/zygomys/zygo"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/pipe"
)

var Env *Zlisp
var Cfg *ZlispConfig

func GetTheAnswer(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	log.Debug("Someone is looking for an answer. Well it is 42")
	return &SexpInt{Val: int64(42)}, nil
}

func LispListToLispArray(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 1 {
		return SexpNull, WrongNargs
	}
	res, err := ListToArray(args[0])
	if err != nil {
		log.Errorf("Error converting list to array: %v", err)
		return SexpNull, err
	}
	return &SexpArray{Val: res, Env: env}, nil
}

func TsakBuiltinFunctions() map[string]zygo.ZlispUserFunction {
	log.Debug("Registering TSAK built-in functions")
	return MergeFuncMap(
		AllBuiltinFunctions(),
		AllTsakCoreFunctions(),
		LogFunctions(),
		PerceptronModuleFunctions(),
		pipe.PipeFunctions(),
		SignalFunctions(),
		SleepFunctions(),
		NowFunctions(),
		FakeFunctions(),
		MathFunctions(),
		RandFunctions(),
		NormalizeFunctions(),
		FloatFunctions(),
		PredictorFunctions(),
		MetricFunctions(),
	)
}

func AllTsakCoreFunctions() map[string]ZlispUserFunction {
	log.Debug("Registering TSAK core functions")
	return map[string]ZlispUserFunction{
		"answer":  GetTheAnswer,
		"toarray": LispListToLispArray,
	}
}

func AllEnvInitBeforeCreationOfEnv() {
	log.Debug("DSL initialization before environment creation")
	PerceptronSetup()
	RandomSetup()
	PredictorSetup()
}
