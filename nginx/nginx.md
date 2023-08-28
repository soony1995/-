# nginx의 sample code를 보고 알아보자.
```
user       www www;  ## Default: nobody 
// Nginx 프로세스가 실행될 때 사용자 및 그룹을 지정하는 옵션입니다. 이 경우 www 사용자와 www 그룹을 사용하도록 설정되어 있습니다.

worker_processes  5;  ## Default: 1 
error_log  logs/error.log;
pid        logs/nginx.pid;
worker_rlimit_nofile 8192;
// 각 워커 프로세스가 열 수 있는 파일 디스크립터(파일 핸들)의 최대 수를 제한하는 옵션입니다.

events {
  worker_connections  4096;  ## Default: 1024
}

http {
  include    conf/mime.types;
  include    /etc/nginx/proxy.conf;
  include    /etc/nginx/fastcgi.conf;
  index    index.html index.htm index.php;

  default_type application/octet-stream;
  //  클라이언트로 보내는 기본 Content-Type을 설정합니다.
  log_format   main '$remote_addr - $remote_user [$time_local]  $status '
    '"$request" $body_bytes_sent "$http_referer" '
    '"$http_user_agent" "$http_x_forwarded_for"';
  access_log   logs/access.log  main;
  sendfile     on;
  tcp_nopush   on;
  server_names_hash_bucket_size 128; # this seems to be required for some vhosts

  server { # php/fastcgi
    listen       80;
    server_name  domain1.com www.domain1.com;
    // 이 서버 블록이 domain1.com 및 www.domain1.com 호스트 이름으로 들어오는 요청을 처리하도록 설정합니다.
    access_log   logs/domain1.access.log  main;
    root         html;

    location ~ \.php$ {
      fastcgi_pass   127.0.0.1:1025;
      // .php로 끝나는 모든 URL 경로에 일치하도록 설정됩니다. 
    }
  }

  server { # simple reverse-proxy
    listen       80;
    server_name  domain2.com www.domain2.com;
    access_log   logs/domain2.access.log  main;

    # serve static files
    location ~ ^/(images|javascript|js|css|flash|media|static)/  {
      root    /var/www/virtual/big.server.com/htdocs;
      expires 30d;
    }

    # pass requests for dynamic content to rails/turbogears/zope, et al
    location / {
      proxy_pass      http://127.0.0.1:8080;
    }
  }

  upstream big_server_com {
    server 127.0.0.3:8000 weight=5;
    server 127.0.0.3:8001 weight=5;
    server 192.168.0.1:8000;
    server 192.168.0.1:8001;
  }

  server { # simple load balancing
    listen          80;
    server_name     big.server.com;
    access_log      logs/big.server.access.log main;

    location / {
      proxy_pass      http://big_server_com;
    }
  }
}
```

workerprocess의 개수를 늘린다고 해서 성능이 향상 되지 않는다.

결론적으로 nginx server에서 이상적인 worker process의 갯수는 본인 server의 core의 갯수입니다.    
각 core는 서로 process를 공유하지 않으므로 모든 core당 1개의 worker process를 사용하면 전체 server의 resource를 최대로 사용하는 셈입니다.


