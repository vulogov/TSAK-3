package dsl

import (
	"fmt"
	"reflect"

	"github.com/elliotchance/orderedmap"
	. "github.com/glycerine/zygomys/zygo"
)

type SexpOrderedMap struct {
	Sexp
	Map *orderedmap.OrderedMap
	Typ *RegisteredType
	Env *Zlisp
}

func (om *SexpOrderedMap) SexpString(ps *PrintState) string {
	return fmt.Sprintf("(OrderedMap: %v)", om.Map)
}

func (om *SexpOrderedMap) Type() *RegisteredType {
	return om.Typ
}

func IsOrderedMap(expr Sexp) bool {
	if expr == SexpNull {
		return false
	}
	switch expr.(type) {
	case *SexpOrderedMap:
		return true
	}
	return false
}

func orderedmapnew(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	return &SexpOrderedMap{Map: orderedmap.NewOrderedMap(), Env: env, Typ: GoStructRegistry.Userdef["OrderedMap"]}, nil
}

func orderedmapset(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 3 {
		return SexpNull, WrongNargs
	}

	switch e := args[0].(type) {
	case *SexpOrderedMap:
		e.Map.Set(AsAny(args[1]), AsAny(args[2]))
		res, succ := e.Map.Get(AsAny(args[1]))
		if !succ {
			return SexpNull, fmt.Errorf("Failure to get key in (OrderedMap)")
		}
		return ToAny(res), nil
	default:
		return SexpNull, fmt.Errorf("First parameter must be (OrderedMap): %v", reflect.TypeOf(e).String())
	}
	return SexpNull, fmt.Errorf("General failure in (OrderedMap)")
}

func orderedmapget(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 2 {
		return SexpNull, WrongNargs
	}
	switch e := args[0].(type) {
	case *SexpOrderedMap:
		res, succ := e.Map.Get(AsAny(args[1]))
		if !succ {
			return SexpNull, fmt.Errorf("Failure to get key in (OrderedMap)")
		}
		return ToAny(res), nil
	default:
		return SexpNull, fmt.Errorf("First parameter must be (OrderedMap): %v", reflect.TypeOf(e).String())
	}
	return SexpNull, fmt.Errorf("General failure in (OrderedMap)")
}

func orderedmapkeys(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 1 {
		return SexpNull, WrongNargs
	}
	switch e := args[0].(type) {
	case *SexpOrderedMap:
		res := e.Map.Keys()
		return ArrayofSomethingToLispArray(env, res), nil
	}
	return SexpNull, fmt.Errorf("First parameter must be (OrderedMap): %v", reflect.TypeOf(args[0]).String())
}

func orderedmapvalues(env *Zlisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 1 {
		return SexpNull, WrongNargs
	}
	switch e := args[0].(type) {
	case *SexpOrderedMap:
		res := make([]interface{}, 0)
		for el := e.Map.Front(); el != nil; el = el.Next() {
			res = append(res, el.Value)
		}
		return ArrayofSomethingToLispArray(env, res), nil
	}
	return SexpNull, fmt.Errorf("First parameter must be (OrderedMap): %v", reflect.TypeOf(args[0]).String())
}

func OrdermapFunctions() map[string]ZlispUserFunction {
	return map[string]ZlispUserFunction{
		"orderedmapnew":    orderedmapnew,
		"orderedmapset":    orderedmapset,
		"orderedmapget":    orderedmapget,
		"orderedmapkeys":   orderedmapkeys,
		"orderedmapvalues": orderedmapvalues,
	}
}

func OrdermapPackageSetup(cfg *ZlispConfig, env *Zlisp) {
	myPkg := `(def orderedmap (package "orderedmap"
     { New := orderedmapnew ;
			 Set := orderedmapset ;
			 Get := orderedmapget ;
			 Keys := orderedmapkeys ;
			 Values := orderedmapvalues ;
     }
  ))`
	_, err := env.EvalString(myPkg)
	PanicOn(err)
}

func OrderedMapSetup() {
	udsR := NewRecordDefn()
	udsR.SetName("OrderedMap")
	rtR := NewRegisteredType(func(env *Zlisp, h *SexpHash) (interface{}, error) {
		return &SexpOrderedMap{Map: orderedmap.NewOrderedMap()}, nil
	})
	rtR.DisplayAs = "OrderedMap"
	rtR.RegisteredName = "OrderedMap"
	rtR.IsUser = true
	rtR.GenDefMap = true
	rtR.ReflectName = "OrderedMap"
	rtR.UserStructDefn = udsR
	GoStructRegistry.RegisterUserdef(rtR, false, "OrderedMap", "SexpOrderedMap", "dsl.SexpOrderedMap")
	// GoStructRegistry.RegisterUserdef(&RegisteredType{GenDefMap: true, Factory: func(env *Zlisp, h *SexpHash) (interface{}, error) {
	// 	return &SexpOrderedMap{Map: orderedmap.NewOrderedMap(), Typ: GoStructRegistry.Userdef["OrderedMap"], Env: env}, nil
	// }}, true, "OrderedMap")
}
