package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatchString(t *testing.T) {
	ok, _ := regexp.MatchString("[a-z]{3}", "ppl")
	fmt.Println(ok)
	ok, _ = regexp.MatchString("[a-z]{3}", "ppl342")
	fmt.Println(ok)
	ok, _ = regexp.MatchString("[a-z]{3}", "p3pl")
	fmt.Println(ok)
	ok, _ = regexp.MatchString("[a-z]{3}", "432ppl423")
	fmt.Println(ok)
}

func TestMatchCompile(t *testing.T) {
	reg := regexp.MustCompile("[a-z]{3}")
	res := reg.FindAllString("abcfda4fd45", -1)
	res2 := reg.FindAllStringSubmatch("abcfda4fd45", -1)
	fmt.Println(res)
	fmt.Println(res2)
}
