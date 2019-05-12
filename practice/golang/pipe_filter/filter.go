package main
/**
pipe filter architecture
 */

type Request interface {}

type Response interface {}

type Filter interface {
	Process (request Request) (response Response, err error)
}
