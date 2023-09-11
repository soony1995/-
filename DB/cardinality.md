## 카디널리티

> 중복도가 ‘낮으면’ 카디널리티가 ‘높다’고 표현한다.  
중복도가 ‘높으면’ 카디널리티가 ‘낮다’고 표현한다.   
카디널리티는 전체 행에 대한 특정 컬럼의 중복 수치를 나타내는 지표이다.

> 컬럼의 distinct의 수가 높다 => 중복도가 낮다. => 카디널리티가 높다. 라고 표현한다.
하지만 이것은 하나의 테이블 내에서의 상대적인 수치일 뿐이다. 

> 카디널리티는 인덱스를 지정하는 데에도 사용할 수 있는데, 중복되는 값이 적을 수록 검색을 하는 데에는 적은 비용이 소모된다.    
따라서, 인덱스를 지정할 때, 카디널리티가 높은 컬럼을 인덱스로 지정하는 것이 옳은 선택이다.

> 인덱스와 WEHRE절을 사용했을 때의 내부동작 
```
SELECT    *
FROM      users
use index (idx_id_first)
WHERE     id = '0'
AND       name = 'lee'
AND       location = 'seoul';
```
위의 예시와 같이 인덱스를 id에 걸어서 sql문을 실행하게 되면, where 절에서도 우선순위가 id가 가장 높은 순서로 필터링하게 된다. 

예시) https://itholic.github.io/database-cardinality/