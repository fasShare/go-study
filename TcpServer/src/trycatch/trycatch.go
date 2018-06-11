package trycatch

func try(panic_source func(), panic_handle func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			panic_handle(err)
		}
	}()
	defer panic_source()
}
