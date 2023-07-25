<<<<<<< HEAD
// regex에서 많이 사용하는 문법들
=======
>>>>>>> 4a996e63a392bf0ddad4d572f915309b689da50a

<<<<<<< HEAD
// 1. . 
// 	-> 모든 단일 문자와 일치
// 2. +
// 	-> 앞의 나온 문자가 1개 이상일 때
// 3. *
// 	-> 앞의 나온 문자가 0개 이상일 때
// 4. [A-Za-z]
// 	-> 알파벳 A~Z 또는 a~z 까지의 문자중에 하나라도 맞을 때
// 5. [문자1|문자2]
// 	-> 문자1 또는 문자2번이 매칭 될 때
// 6. ^
// 	-> 뒤에 나오는 문자로 시작할 때
// 7. [^]
// 	-> 대괄호 안에 문자를 포함하지 않을 때
// 8. $
// 	-> 앞에 나온 문자로 끝날 때

a := "Raw string literals are c`haracter sequences between back quotes, as in `foo`. Within the quotes, any character may appear except back quote. "
re := regexp.MustCompile(``)
matches := re.FindAllStringSubmatch(a, -1)

fmt.Println(matches)


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
