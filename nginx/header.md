## X-FORWARDED-FOR 와 X-REAL-IP의 차이에 대해서 알아보자

### X-FORWARDED-FOR
> client <-> nginx <-> server 

이 상황에서 server가 client의 주소를 알고 싶을 때 사용하는 헤더이다.

하지만 X-FORWARDED-FOR는 변조가 가능하다.

### X-REAL-IP

> X-REAL-IP 역시도 client의 주소를 알기 위해 사용하는 헤더이다.
X-REAL-IP: 192.168.1.1 의 형식으로 보여진다.

```
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://backend_server;
        proxy_set_header X-REAL-IP $remote_addr;
    }
}
```

