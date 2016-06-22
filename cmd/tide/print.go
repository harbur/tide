package main // import "github.com/harbur/tide/cmd/tide"

import "fmt"

func info(text string, arg ...interface{}) {
	text = colorInfo("[") + colorPrefix("TIDE") + colorInfo("]") + " " + text + "\n"
	s := arg
	for i := range s {
		s[i] = colorInfo(s[i])
	}
	fmt.Printf(text, arg...)
}

func print_error(text string, arg ...interface{}) {
	text = colorErr("[") + colorPrefix("TIDE") + colorErr("]") + " " + text + "\n"
	s := arg
	for i := range s {
		s[i] = colorErr(s[i])
	}
	fmt.Printf(text, s...)
}

func debug(text string, arg ...interface{}) {
	if verbose {
		text = colorDebug("[") + colorPrefix("TIDE") + colorDebug("]") + " " + text + "\n"
		s := arg
		for i := range s {
			s[i] = colorDebug(s[i])
		}
		fmt.Printf(text, s...)
	}
}
