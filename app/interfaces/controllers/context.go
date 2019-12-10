package controllers

type Context interface {
	Param(string) string
	Query(string) string
	Bind(interface{}) error
	BindJSON(interface{}) error
	GetHeader(key string) string
	Status(int)
	JSON(int, interface{})
	AbortWithStatusJSON(int, interface{})
}
