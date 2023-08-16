## journalctl
> systemd의 로그를 journal로 관리하는데 해당 journal들을 조회할 수 있는 소프트웨어를 말합니다.

### 특정 유닛의 로그를 확인   
-u 옵션을 사용   
journalctl -u [systemd unit name]


### 로그를 계속 확인하며 트래킹   
-f 옵션을 사용   
journalctl -f

### 최근 10 개 메시지 표시
-n 옵션을 사용
journalctl -n

journalctl -n 7 (7개의 최근 로그 확인)

### journalctl vs syslog 

- 로그 포맷 및 저장 위치:
    - syslog: syslog은 리눅스 시스템의 기본 로깅 시스템으로, 다양한 프로그램 및 시스템 컴포넌트에서 생성된 로그를 ***/var/log/ 디렉터리에 텍스트 파일 형태로 저장***합니다. syslogd 또는 rsyslogd 같은 데몬을 통해 관리됩니다.

    - journalctl: journalctl은 systemd의 기능 중 하나로, ***로그를 바이너리 형식의 저널 파일에 저장***합니다. 이 파일은 /run/log/journal/ 또는 /var/log/journal/에 위치하며, ***시스템의 부팅 과정부터 종료까지의 로그를 기록***합니다.

- 로그 형식:
    - syslog: syslog 로그는 주로 텍스트 형식으로 저장됩니다. 여러 개의 필드로 구성되어 있으며, 일반적으로 날짜, 시간, 호스트명, 프로그램 이름, 로그 메시지 등의 정보를 포함합니다.

    - journalctl: journalctl은 로그를 바이너리 형식으로 저장하며, 로그 항목에는 JSON 형식으로 구조화된 메타데이터와 메시지가 포함됩니다. 이는 로그를 분석하고 필터링하기에 더 효율적인 구조를 가지게 합니다.

### 결론: 
    1. Syslog는 텍스트 파일, journalctl은 바이너리 파일로 로그를 저장한다.
    2. journalctl은 시스템의 부팅부터 기록한다.


