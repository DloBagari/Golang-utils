package example1

import (
	gc "gopkg.in/check.v1"
	"testing"
)

var _ = gc.Suite(new(SuiteBase))

func Test(t *testing.T) { gc.TestingT(t) }
