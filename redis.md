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

## **실제로 사용면서 느낀 점**

단점: 
구조가 없기때문에 value내에서 정교한 검색이 불가능.   
key값으로만 검색해야하므로 key 디자인을 잘해야함. (구조가 없는게 장점이자 단점 )
