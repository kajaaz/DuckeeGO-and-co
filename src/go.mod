module main

go 1.20

require (
	github.com/aclements/go-z3 v0.0.0-20220809013456-4675d5f90ca5
	github.com/otiai10/copy v1.14.0
	gitlab.com/mgmap/maps v0.0.0-20150609092100-ce307106584b
	golang.org/x/tools v0.16.0
)

require (
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace concolicTypes => ../concolicTypes
