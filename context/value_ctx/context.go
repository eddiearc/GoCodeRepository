package value_ctx

type Context interface {
	Value(key any) any
}

func WithValue(c Context, key, value any) Context {
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
	key, value any
}

func (c *ValueCtx) Value(key any) any {
	if key == c.key {
		return c.value
	}
	return c.Context.Value(key)
}
