package types

//go:generate genny -pkg bytes -in=$HOME/go/src/github.com/itsmontoya/turtle/. -out=bytes/turtle.go gen "Value=[]byte"
//go:generate genny -pkg bytes -in=$HOME/go/src/github.com/itsmontoya/turtle/. -out=generic/. gen "Value=interface{}"
