## ERD? 

> ERD란 Entity Relationship Diagram의 약어로, 데이터베이스 구조를 한눈에 알아보기 위해서 쓰인다.

> ERD의 핵심은 세 가지 이다. "Entity" 와 "Relationship", 그리고 "Attribute"이다. 각 용어가 무엇을 의미하는지 살펴보자.

### Entitiy
> Entity는 관리하고자 하는 정보의 실체이며, 사람, 객체 혹은 개념이라고 이해하면 된다.

- 데이터베이스를 설계할때, 쉽게는 테이블이 Entity로 정의될 수 있다.

- 모든 Entity는 하나 이상의 식별자 (UID)를 지녀야 하며, UID가 없다면 Entity라고 할 수 없다.
### Attribute 
> Attribute는 Entity를 구성하고 있는 구성 요소이다.

- 데이터 타입을 반드시 같이 명시해줘야 한다.

- Key Attribute
다른 객체들과 중복되지 않는 고유한 값을 가진 Attribute로, 객체를 식별하는 데 사용된다.

- Composite Attribute
독립적인 Attribute들이 모여서 생성된 Attribute를 의미한다.
예를 들어 OO시,OO동,OO아파트 등의 독립적인 Attribute 네개가 모여서 생성된 주소 Attribute는 복합 attribute라고 할 수 있다.

- Multi-Valued Attribute
하나의 Attribute가 여러개의 값을 가지는 Attribute를 의미한다.
예를 들어 하나의 영상물에 로맨스, SF, 호러 등의 여러가지 장르가 공통적으로 존재할 수 있다.

- Derived Attribute
이는 다른 Attribute가 갖고 있는 값으로부터 유도된 속성을 의미한다.
예를 들어 모든 상품의 총 가격을 나타내는 total이라는 속성은 상품의 가격 attribute, 상품의 개수 attribute를 곱해서 계산된 값이다. 이는 Derived Attribute이 된다.

### Relationship

> Entity 간의 관계를 의미한다.


