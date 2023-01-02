package optional

type Optional[T any] struct {
	empty bool
	value T
}

func None[T any]() Optional[T] {
	return Optional[T]{
		empty: true,
	}
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		empty: false,
		value: value,
	}
}

func From[T any](value T) Optional[T] {
	return Some[T](value)
}

func FromNillable[T any](value interface{}) Optional[T] {
	if value == nil {
		return Optional[T]{
			empty: true,
		}
	}

	return Optional[T]{
		empty: false,
		value: value.(T),
	}
}

func FromError[T any](value T, err error) Optional[T] {
	if err != nil {
		return None[T]()
	}

	return Some[T](value)
}

func (optional Optional[T]) IsEmpty() bool {
	return optional.empty
}

func (optional Optional[T]) IsPresent() bool {
	return !optional.IsEmpty()
}

func (optional Optional[T]) Get() T {
	return optional.value
}

func (optional Optional[T]) GetTry() (value T, err error) {
	if optional.IsEmpty() {
		err = ERROR_EMPTY
		return
	}

	value = optional.value
	return
}

func (optional Optional[T]) GetDefault(value T) T {
	if optional.IsEmpty() {
		return value
	}

	return optional.value
}

func (optional Optional[T]) GetElse(fn func() T) T {
	if optional.IsEmpty() {
		return fn()
	}

	return optional.value
}

func (optional Optional[T]) IfPresent(fn func(T)) {
	if optional.IsPresent() {
		fn(optional.Get())
	}
}

func (optional Optional[T]) IfPresentElse(fn func(T), fn_ func()) {
	if optional.IsPresent() {
		fn(optional.Get())
	} else {
		fn_()
	}
}
