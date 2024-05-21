package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceHolder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceHolder, traceLevel-1)
}
