## journalctl
> systemd의 로그를 journal로 관리하는데 해당 journal들을 조회할 수 있는 소프트웨어를 말합니다.

### 예제 - 특정 유닛의 로그를 확인   
-u 옵션을 사용   
journalctl -u [systemd unit name]


### 예제 - 로그를 계속 확인하며 트래킹   
-f 옵션을 사용   
journalctl -f