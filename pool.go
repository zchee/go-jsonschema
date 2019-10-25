package jsonschema

import "sync"

// Pooler represents a poolable interface.
type Pooler interface {
	Reset()
}

// models
var (
	ArrayPool   *sync.Pool
	BooleanPool *sync.Pool
	IntegerPool *sync.Pool
	NullPool    *sync.Pool
	NumberPool  *sync.Pool
	ObjectPool  *sync.Pool
	StringPool  *sync.Pool
)

func init() {
	ArrayPool = &sync.Pool{
		New: func() interface{} {
			return &Array{}
		},
	}
	BooleanPool = &sync.Pool{
		New: func() interface{} {
			return &Boolean{}
		},
	}
	IntegerPool = &sync.Pool{
		New: func() interface{} {
			return &Integer{}
		},
	}
	NullPool = &sync.Pool{
		New: func() interface{} {
			return &Null{}
		},
	}
	NumberPool = &sync.Pool{
		New: func() interface{} {
			return &Number{}
		},
	}
	ObjectPool = &sync.Pool{
		New: func() interface{} {
			return &Object{}
		},
	}
	StringPool = &sync.Pool{
		New: func() interface{} {
			return &String{}
		},
	}
}

// types
var (
	TypesPool       *sync.Pool
	StringArrayPool *sync.Pool
	ItemsPool       *sync.Pool
	ConstPool       *sync.Pool
	EnumPool        *sync.Pool
	SchemaPool      *sync.Pool
)

func init() {
	ConstPool = &sync.Pool{
		New: func() interface{} {
			return &Const{}
		},
	}
	EnumPool = &sync.Pool{
		New: func() interface{} {
			return &Enum{}
		},
	}
	ItemsPool = &sync.Pool{
		New: func() interface{} {
			return &Items{}
		},
	}
	StringArrayPool = &sync.Pool{
		New: func() interface{} {
			return &StringArray{}
		},
	}
	TypesPool = &sync.Pool{
		New: func() interface{} {
			return &Types{}
		},
	}
}

// schema
var (
	Draft7Pool *sync.Pool
)

func init() {
	Draft7Pool = &sync.Pool{
		New: func() interface{} {
			return &Draft7{}
		},
	}
}
