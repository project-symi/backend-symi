package controllers

type Context interface {
	Param(string) string
	Query(string) string
	BindJSON(interface{}) error
	ShouldBind(interface{}) error
	ShouldBindUri(interface{}) error
	ShouldBindJSON(interface{}) error
	Bind(interface{}) error
	GetHeader(key string) string
	Status(int)
	JSON(int, interface{})
	AbortWithStatus(int)
}
