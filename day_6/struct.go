package main

type BuildQueryExcutor func(val string) int

type Database struct{
	HostName   string
	UserName   string
	Password   string
	BuildQuery BuildQueryExcutor
}

// embedings