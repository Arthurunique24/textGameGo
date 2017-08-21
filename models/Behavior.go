package models

type  Behavior interface{
 	Insert() int
	Check() int
}