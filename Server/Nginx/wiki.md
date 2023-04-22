# server_name

서버네임과 일치하는 요청이 있으면,

서버네임이 작성된 곳으로 요청을 보낸다.

\*.example.com 처럼 와일드 카드를 쓸 수도 있고,
정규식을 사용할 수도 있다.

# try files

try files은
모든 request에 대해 일치하는 path를
앞에서 부터 순서대로 비교한뒤
root에 존재하는 path를 rewrite 하는 명령어입니다.

```conf
try_files $uri $uri/ index.html =404;
```

uri 가 없으면, index.html로 보내겠다는 뜻.

# upsteeam

https://opentutorials.org/module/384/4328 참고
