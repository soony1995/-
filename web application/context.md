### echo framework에서 context의 개념에 대해서 정리.

> echo.Context represents the context of the current HTTP request. It holds request and response reference, path, path parameters, data, registered handler and APIs to read request and write response. As Context is an interface, it is easy to extend it with custom APIs.

라고 echo 공식 docs에 기재되어 있다. 

즉, 사용자가 http 요청을 할 때 기입한 정보들을 가지고 있는 객체라고 보면 된다.

### 이러한 context를 사용해 router를 타기전에 필터링 할 수 있게 해주는 것이 middleware 이다. 