
package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "Raw string literals are character sequences between back quotes, as in `foo`. Within the quotes, any character may appear except back quote. "
	re := regexp.MustCompile(`c.+`) //MustCompile는 에러를 반환해 주지 않는다.
	matches := re.FindAllStringSubmatch(a, -1) // 두번째 인자로 -1을 주게되면 매칭되는 모든 문자를 찾아준다.
	fmt.Print(matches)
}
