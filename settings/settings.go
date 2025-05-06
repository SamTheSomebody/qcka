// Copyright © 2025 Sam Muller gamedevsam@pm.me
package settings

type runSettings struct {
	Force   bool
	Quiet   bool
	Verbose bool
}

var Settings *runSettings

func New() {
	Settings = &runSettings{}
}
