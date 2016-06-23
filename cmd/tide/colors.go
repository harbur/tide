package main // import "github.com/harbur/tide/cmd/tide"

import (
	"github.com/fatih/color"
)

var colorPrefix = color.New(color.FgWhite, color.Bold).SprintFunc()
var colorDebug = color.New(color.FgBlue).SprintFunc()
var colorInfo = color.New(color.FgGreen).SprintFunc()
var colorWarn = color.New(color.FgYellow).SprintFunc()
var colorErr = color.New(color.FgRed).SprintFunc()
