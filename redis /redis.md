# Redis

 ## **Key word**
 - cache   
  정의: **자주 사용되는 정보들을 미리 저장해놓은 뒤에 임시로 저장하는 것을 말한다**   
  redis를 이용해 처음 DB에서 꺼내온 데이터일 경우 redis에 저장한 후에  다른 사용자가 똑같은 데이터를 요청할 때 빠르게 서비스 할 수 있게 해준다.  
 - 싱글 스레드  
 - key-value
   조회속도가 O(1)~O(n)으로 빠른 편에 속하므로, 자주 사용하는 정보를 가져올 때 유용하다.      
   참고자료: 
   - https://bravenamme.github.io/2020/11/05/key-value-store/
   - https://velog.io/@xowen96/Chaining%EC%9D%84-%EC%93%B0%EB%8A%94-%ED%95%B4%EC%8B%9C-%ED%85%8C%EC%9D%B4%EB%B8%94-%EC%8B%9C%EA%B0%84%EB%B3%B5%EC%9E%A1%EB%8F%84-%ED%83%90%EC%83%89%EC%82%BD%EC%9E%85%EC%82%AD%EC%A0%9C
 - AOF(append on file) -> 트랜잭션을 로그에 남긴다.
 - RDB(snapshotting) 


## **사용법**   

```
golang으로 작성한 코드입니다.

######<redis-server>#####
https://hayden-archive.tistory.com/429을 참고해 서버를 설치했다. 

######<redis-client>#####
package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var DB struct {
	client *redis.Client
}

func initDBconn() {
	DB.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {

	initDBconn()
	fmt.Println(DB.client.Set("set-redis-key", "redis-value", 0))

	fmt.Println(DB.client.Get("set-redis-key"))
}

```
## **실제로 사용면서 느낀 점**
장점:
현재 로그인 여부, 파일 IO 여부, 트래픽량 트래킹을 redis에 저장하여 사용했을 때 확실히 maria db에 넣어서 쓸 때 보다 조회나 삭제가 자유로워서 편했다.

단점: 
value 값이 다양하게 들어갈 때가 있는데, 구조가 없다보니 가져다 쓰기 불편함이 존재한다. 

