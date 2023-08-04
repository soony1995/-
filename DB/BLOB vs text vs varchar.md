# varchar 

- 최대 저장 크기 : 65,535 bytes ~= 65KB

- 영어,숫자는 1byte로 처리하고, 한글은 3bytes로 처리한다.

- 따라서 한글로만 21,845의 글자만 저장이 가능하다. 

- varchar(21845) -> mysql 4.1 이후 버전부터 괄호 안의 숫자는 글자 수를 나타낸다. 21845의 글자를 저장할 수 있다는 뜻이다. 

# BLOB (Binary Large Object)

- 최대 저장 크기 : 2,147,483,647 bytes ~= 2.1GB

- 이미지, 비디오, 사운드 등과 같은 멀티미디어 객체들을 담을 때 사용.

- BLOB 타입 : TINYBLOB, BLOB, MEDIUMBLOB, LONGBLOB (저장할 수 있는 최대치만 다름)

- BLOB 데이터는 일반적으로 기본 테이블 저장소 외부에 저장되며 사용 중인 데이터베이스 시스템에 따라 데이터베이스 또는 파일 시스템에서 별도의 위치를 가질 수 있습니다.

- BLOB 값 : binary strings (byte strings) character set, collation : binary character set 과 collation 만 가질 수 있습니다.

# TEXT

- BLOB과 비슷함.

- character set, collation : binary character set 과 collation 만 가질 수 있습니다.

## 
요약하면 주요 차이점은 BLOB는 이진 데이터용이고 TEXT는 문자 기반 텍스트 데이터용이라는 것입니다.      

이미지 및 멀티미디어 파일을 저장할 때 권장되는 방법은 이진 데이터를 처리하도록 특별히 설계된 BLOB(Binary Large OBject) 데이터 형식을 사용하는 것입니다.     

문서 또는 기사와 같은 큰 텍스트 기반 콘텐츠를 저장해야 하는 경우 TEXT를 사용합니다.