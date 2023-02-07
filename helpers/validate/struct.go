package validation

type Validator struct {
	Entity interface{}
	DTO    interface{}
	ArrDTO interface{}
	Error  []interface{}
	// DB     *gorm.DB
}

type Model[T any] struct {
	Data *T
}
