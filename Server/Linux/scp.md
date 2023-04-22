# scp란?

SecureCopy의 약자.

ssh 원격 접속 프로토콜을 기반으로 만들어짐.

로컬 -> 원격 (파일명이 ip보다 앞에 있으면)
원격 -> 로컬 (파일명이 ip보다 뒤에 있으면)

파일 전송이 가능한 프로그램

## 로컬 -> 원격

scp [옵션] [파일명] [원격지_id]@[원격지_ip]:[받는 위치]

- directory 를 보내고 싶을 때.

```
scp -P 3000 -r dist/ xx@xx.xx.xx.xx:/home/nickname
```

- 파일 만 보낼 때

```
scp testfile root@xx.xx.xx.xx:/home/nickname
```

## 원격 -> 로컬

scp [옵션] [원격지_id]@[원격지_ip]:[원본 위치] [받는 위치]
