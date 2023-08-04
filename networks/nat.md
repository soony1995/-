# NAT의 종류

1. static nat

    - private IP : public IP 1:1로 매핑 됨
    - IP 절약 효과는 없다.
2. dynamic nat
    - private IP < public IP 일 때, 랜덤으로 매핑 해준다.
    - 고정적으로 IP 를 매핑해 줄 필요가 없기 때문에 IP가 계속 달라진다.
3. pat nat (port address translation)
    - private IP > public IP 일 때, IP 주소는 한정 되어 있기 때문에 하나의 IP 주소에서 port번호로 각각의 private IP를 구분.