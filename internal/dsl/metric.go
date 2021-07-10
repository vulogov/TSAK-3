package dsl

import (
	. "github.com/glycerine/zygomys/zygo"
)

func MetricFunctions() map[string]ZlispUserFunction {
	return map[string]ZlispUserFunction{}
}

func MetricPackageSetup(cfg *ZlispConfig, env *Zlisp) {
	myPkg := `
	(defmap Metric)
	(def metric (package "metric"
     { New := (fn [source key val] (Metric Src: source Key: key X: (now.UTCNano) Y: val)) ;
			 Marshal := (fn [m] (raw2str (msgpack m))) ;
			 Unmarshal := (fn [s] (unmsgpack (raw s))) ;
			 UpdateX := (fn [m] (hset m X: (now.UTCNano)) (:X m)) ;
			 Y := (fn [m val] (hset m Y: val) (:Y m)) ;
     }
  ))`
	_, err := env.EvalString(myPkg)
	PanicOn(err)
}
