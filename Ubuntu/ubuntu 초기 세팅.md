## useradd vs adduser

### useradd 
> 계정만 생성, 비밀번호, 홈 디렉토리 등 기타 부가적인 부분들은 따로 진행해야 함.

### adduser
> 기본 계정 정보, 홈 디렉토리, 쉘 설정 등 한 번에 진행되면서 생성된다.

결론) adduser를 사용하자. 

### ping
> 'ping'은 네트워크 상태를 점검하기 위한 유용한 도구 중 하나입니다. 다음은 'ping'을 사용하여 네트워크 상태를 점검하는 이유입니다:
    
    ping은 몇번 포트에다가 요청을 하는 걸까?

    => 'ping'은 ICMP(Internet Control Message Protocol)를 사용하여 목표 호스트에 데이터 패킷을 보냅니다. ICMP는 포트 번호를 사용하지 않고 IP 패킷을 전송하므로 'ping' 명령은 포트 번호를 명시적으로 지정하지 않습니다.

    ICMP란?

    => ICMP는 "Internet Control Message Protocol"의 약자로, 인터넷 프로토콜 스위트에서 사용되는 네트워크 계층 프로토콜 중 하나입니다. ICMP는 네트워크 장치 간에 메시지를 교환하여 네트워크에서 발생하는 다양한 문제를 진단하고 통신을 관리하기 위해 설계되었습니다. ICMP는 IP(Internet Protocol)와 함께 작동하여 네트워크 상황을 모니터링하고 제어하는 데 사용됩니다.
### package download server 
/etc/apt/sources.list에서 변경이 가능하다.
```
# See http://help.ubuntu.com/community/UpgradeNotes for how to upgrade to
# newer versions of the distribution.
deb http://archive.ubuntu.com/ubuntu/ focal main restricted
# deb-src http://archive.ubuntu.com/ubuntu focal main restricted

## Major bug fix updates produced after the final release of the
## distribution.
deb http://archive.ubuntu.com/ubuntu/ focal-updates main restricted
# deb-src http://archive.ubuntu.com/ubuntu focal-updates main restricted

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team. Also, please note that software in universe WILL NOT receive any
## review or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu/ focal universe
# deb-src http://archive.ubuntu.com/ubuntu focal universe
deb http://archive.ubuntu.com/ubuntu/ focal-updates universe
# deb-src http://archive.ubuntu.com/ubuntu focal-updates universe

## N.B. software from this repository is ENTIRELY UNSUPPORTED by the Ubuntu
## team, and may not be under a free licence. Please satisfy yourself as to
## your rights to use the software. Also, please note that software in
## multiverse WILL NOT receive any review or updates from the Ubuntu
## security team.
deb http://archive.ubuntu.com/ubuntu/ focal multiverse
# deb-src http://archive.ubuntu.com/ubuntu focal multiverse
deb http://archive.ubuntu.com/ubuntu/ focal-updates multiverse
# deb-src http://archive.ubuntu.com/ubuntu focal-updates multiverse

## N.B. software from this repository may not have been tested as
## extensively as that contained in the main release, although it includes
## newer versions of some applications which may provide useful features.
## Also, please note that software in backports WILL NOT receive any review
## or updates from the Ubuntu security team.
deb http://archive.ubuntu.com/ubuntu/ focal-backports main restricted universe multiverse
# deb-src http://archive.ubuntu.com/ubuntu focal-backports main restricted universe multiverse

## Uncomment the following two lines to add software from Canonical's
## 'partner' repository.
## This software is not part of Ubuntu, but is offered by Canonical and the
## respective vendors as a service to Ubuntu users.
# deb http://archive.canonical.com/ubuntu focal partner
# deb-src http://archive.canonical.com/ubuntu focal partner

deb http://security.ubuntu.com/ubuntu focal-security main restricted
# deb-src http://security.ubuntu.com/ubuntu focal-security main restricted
deb http://security.ubuntu.com/ubuntu focal-security universe
# deb-src http://security.ubuntu.com/ubuntu focal-security universe
deb http://security.ubuntu.com/ubuntu focal-security multiverse
# deb-src http://security.ubuntu.com/ubuntu focal-security multiverse
```
:%s/현재 사용 하고있는 서버/변경할 서버주소/ 
을 통해서 변경하면 된다.

### DNS 서버 관리
1. DNS의 서버의 구성 정보 체크
> sudo systemd-resolve --status
2. DNS의 서버 상태 체크
> sudo systemctl status systemd-resolved
3. DNS 서버의 conf를 변경
> sudo vi /etc/resolved.conf 에서 변경하도록 한다.
4. DNS 서버 reload
> sudo systemctl restart systemd-resolved



