package main

import (
	"github.com/wrk-grp/errnie"
	"github.com/wrk-grp/wrkspc/cmd"
)

func main() {
	errnie.Kills(cmd.Execute())
}
