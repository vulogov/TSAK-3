package dsl

import (
	. "github.com/glycerine/zygomys/zygo"
)

func SampleFunctions() map[string]ZlispUserFunction {
	return map[string]ZlispUserFunction{}
}

func SamplePackageSetup(cfg *ZlispConfig, env *Zlisp) {
	myPkg := `
	(defmap Sample)
	(def sample (package "sample"
     { New := (fn [source key] (Sample Src: source Key: key Stamp: (now.UTCNano) Sample: (orderedmap.New))) ;
			 X := (fn [h] (def s []) (range k v h (set s (append s k))) s) ;
			 Y := (fn [h] (def s []) (range k v h (set s (append s v))) s) ;
			 Add := (fn [m]
				 (cond (not (== (type? m) "Metric")) false
				   (begin
						 true
					 )
			   )
			 )
     }
  ))`
	_, err := env.EvalString(myPkg)
	PanicOn(err)
}
