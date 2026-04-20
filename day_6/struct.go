package main

type BuildQueryExcutor func(val string) int

type Database struct{
	HostName   string
	UserName   string
	Password   string
	BuildQuery BuildQueryExcutor
}

// embedings
type PostgresDatabase struct {
	Database
	Version	string
	MasterDBCount int
	CPULimit string
	MemoryLimit string
}

type simple struct{
	Check bool
}


