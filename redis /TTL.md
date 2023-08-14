### TTL 

TTL (time-to-live)은

IP 패킷 내에 있는 값으로서, 그 패킷이 네트웍 내에 너무 오래 있어서 버려져야 하는지의 여부를

라우터에게 알려준다.
패킷들은, 여러 가지 이유로 적당한 시간 내에 지정한 장소에 배달되지 못하는 수가 있다.
예를 들어, 부정확한 라우팅 테이블의 결합은 패킷을 끝없이 순환하게 만들 수도 있다.
일 정한 시간이 지나면 그 패킷을 버리고,재전송할 것인지를 결정하도록
그 사실을 발신인에게 알릴 수 있게 하기 위한 해결책으로 TTL이 사용된다.

TTL의 초기치는
시스템에 의해 대개 8 비트 길이의 패킷 헤더에 설정된다.
TTL 에 관한 원래의 아이디어는,
일정시간동안을 초단위로 지정함으로써 그 시간이 지나면
패킷을 버리도록 하기 위한 것이었다.
각 라우터는 TTL 필드로부터 적어도 하나의 숫자를 빼도록 되어 있으며,
그 계산은 대개 패킷이 버려지기 전에
허용되는 라우터 홉수를 의미하는데 사용된다.
패킷을 받는 각 라우터는 TTL 필드 내의 값에서 1을 뺀다.

그 값이 0이 되었을 때,
라우터는 그것을 감지하여 그 패킷을 버리고
ICMP 메시지를 발신지 호스트로 보낸다.

윈 도우95나 98의 TTL 기본 값은 32 홉이다.
만약 어떤 사이트에 도달하기에 어려움을 느끼는 사용자라면,
이 값을 128 정도로 바꾸는 것도 하나의 방법이다.

쉽게 말하지면...

ㅇ ping test는 망계층(OSI Layer3)의 ICMP(internet control message protocol)를
   이용하여 그 기능을 제공하고 있습니다....

ㅇ ICMP메세지는 type에 따라 여러종류가 있으나
   ping test는 이중 type 8,0(echo request, echo replay)번을 사용합니다.
    
ㅇ 송신측에서 상대편 ip address로 ping test를 하면
   icmp type 8번 packet이 상대편 호스트로 전달되고
   상대편 호스트는 icmp type 0번 packet을 송신측으로 재전송함

ㅇ 송신측에서 전송하고 수신할 때 까지의 시산을 RTT(Round Trip Time)이며
   그냥 간단히 "Time"이라고 쓰기도 합니다...

ㅇ icmp패킷은 인터넷망에서 목적지에 찾아가기 위해서
   ip패킷으로 encapsulation되어 전송되게 된다.
   이때 ip 헤더부분에 8bit의 TTL(Time To Live)필드가 존재합니다.
   TTL필드는 ip패킷이 한 홉(Router)을 지날 때 마다 1씩 감소하여 0이 되면
   그 패킷은 인터넷상에서 영원히 삭제됩니다(Router에서 패기처분)
   이는 루핑되는 ip패킷에 의한 인터넷망의 부하를 줄이도록 설계된 부분입니다.

ㅇ 결론적으로 TTL에 8bit가 할당되었으므로,
   TTL이 가질 수 있는 가장 큰 값은 2의 8승 = 255가 됩니다.
   그 값은 어플리케이션마다 default값이 다르나
   dos에서는 255가 default가 됩니다.

ㅇ ping test한 pc와 dns(168.126.63.1)간에
   Router가 존재하는 갯수(홉수)만큼 빼면은  TTL이 된다...

      255 - Router수(홉수) = TTL      

 