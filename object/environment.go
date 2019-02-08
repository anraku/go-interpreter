// 束縛されている変数を管理するための環境を定義
package object

// 新しい環境を返す
// outerがそれを包むようにすることで環境を拡張する
func NewEncloseEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	// eの中に束縛されている変数が見つからなければ再帰的に外側の環境も検索する
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
