package main

import (
	//"strings"
  //s"fmt"
  "testing"
)

func TestParser(t *testing.T) {
  ivyCache := "/Users/daniev/.ivy2/cache/"
  s := Parse2Map(ivyCache, "x")

  Write2File(s)

}
