package value_ctx

type Context interface {
	Value(key interface{}) interface{}
}

func WithValue(c Context, key, value interface{}) Context {
	if c == nil {
		panic("context not nil")
	}
	if key == nil {
		panic("nil key")
	}
	return &ValueCtx{
		Context: c,
		key:     key,
		value:   value,
	}
}

type ValueCtx struct {
	Context
	key, value interface{}
}

func (c *ValueCtx) Value(key interface{}) interface{} {
	if key == c.key {
		return c.value
	}
	return c.Context.Value(key)
}
