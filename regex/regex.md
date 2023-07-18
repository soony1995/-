## regex에서 많이 사용하는 문법들

  1. "."   
    -> 모든 단일 문자와 일치
  2. "+"   
     -> 앞의 나온 문자가 1개 이상일 때
  3. "*"   
     -> 앞의 나온 문자가 0개 이상일 때
  4. [A-Za-z]   
     -> 알파벳 A~Z 또는 a~z 까지의 문자중에 하나라도 맞을 때
  5. [문자1|문자2]   
     -> 문자1 또는 문자2번이 매칭 될 때
  6. ^   
     -> 뒤에 나오는 문자로 시작할 때
  7. [^]   
     -> 대괄호 안에 문자를 포함하지 않을 때
  8. $   
     -> 앞에 나온 문자로 끝날 때

## regex flag 

- i 
> i 또는 Insensitive: 대소문자를 구분하지 않는 플래그입니다. 이 플래그가 설정된 경우, 정규 표현식은 대소문자를 구분하지 않고 일치 항목을 찾습니다. 예를 들어, /abc/i는 "Abc", "ABC", "abc" 등을 모두 찾아냅니다.

- g	
> 정규 표현식에서 'g' 또는 'global' 플래그는 정규 표현식이 문자열 전체에서 모든 일치하는 항목을 찾아야 함을 나타냅니다.   
예시)   
    /foo/ --> "foo bar foo baz" --> matches: ["foo"]   
    /foo/g --> "foo bar foo baz" --> matches: ["foo", "foo"]   
'g' 플래그는 문자열에서 일치하는 모든 항목을 찾는 데 유용하므로, 일치하는 항목이 여러 개 있을 수 있는 경우나 모든 일치하는 항목을 바꾸려는 경우에 사용합니다

- s	
> s 또는 Singleline (Dotall in some languages): 단일 줄 모드를 나타내는 플래그입니다. 이 플래그가 설정된 경우, "dot" 메타문자 (.)는 줄바꿈 문자를 포함한 모든 문자를 일치시키게 됩니다. 일반적으로, "dot" 메타문자는 줄바꿈 문자를 일치시키지 않습니다.

- m	
> 정규 표현식에서 'm' 또는 'multiline' 플래그는 여러 줄 문자열에서 '^'와 '$' 메타문자가 각각 문자열의 시작과 끝 뿐만 아니라 각 줄의 시작과 끝을 의미하게 함을 나타냅니다.   
예시)   
Hello   
World   
이 경우, /^H.*/ 정규 표현식은 "Hello"만 찾지만, /^H.*/m 정규 표현식은 "Hello"와 "World" 모두를 찾습니다.   
이 'm' 플래그는 여러 줄 문자열에서 각 줄을 개별적으로 처리하려는 경우에 유용합니다.


- y	
>y 또는 Sticky: "Sticky" 검색을 수행하는 플래그입니다. 이 플래그가 설정된 경우, 정규 표현식은 대상 문자열에서 마지막으로 일치한 위치 이후에만 일치 항목을 찾습니다. 이 플래그는 대부분의 정규 표현식 구현에서 사용되지 않지만, JavaScript에서는 지원됩니다.

- u	
> u 또는 Unicode: 유니코드 일치를 수행하는 플래그입니다. 이 플래그가 설정된 경우, 정규 표현식은 대상 문자열을 유니코드 문자열로 취급하고, 특별한 유니코드 이스케이프 시퀀스 (\u{...})를 인식합니다. 또한, "dot" 메타문자 (.)는 모든 유니코드 문자를 포함한 모든 문자를 일치시키게 됩니다.

## golang 에서 regexp에서의 flag 사용법

(re)           numbered capturing group (submatch)   
(?P<name>re)   named & numbered capturing group (submatch)   
(?:re)         non-capturing group   
(?flags)       set flags within current group; non-capturing -> (?m,i,s,u) 를 붙여서 사용하면 된다.  
(?flags:re)    set flags during re; non-capturing  
  
Flag syntax is xyz (set) or -xyz (clear) or xy-z (set xy, clear z). The flags are:  

i              case-insensitive (default false)   

m              multi-line mode: ^ and $ match begin/end line in addition to begin/end text 
 (default false)

s              let . match \n (default false)

U              ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)