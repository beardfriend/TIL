https://jdm.kr/blog/2 을 보고 정리했습니다.

> "특정 시간에 특정 작업을 해야 한다"

라고 한다면, 리눅스에서 크론 탭을 이용할 수 있다.

# 크론탭 기본

크론탭을 설정할 수 있는 파일을 열어준다.

```bash
crontab -e
```

현재 크론탭에 무엇이 들어있는지 조회하는 명렁어다.

```bash
crontab -l
```

크론탭을 지우고 싶을 때 사용하는 명령어입니다.

```bash
crontab -r
```

```bash
* * * * * ls -al
```

별이 다섯 개가 있고 뒤에는 리눅스 명령어가 있다.
별이 다섯 개 있는 경우에는 매 분마다 뒤의 명령어를 실행한다.

# 주기 설정

```bash
*　　　　　　*　　　　　　*　　　　　　*　　　　　　*
분(0-59)　　시간(0-23)　　일(1-31)　　월(1-12)　　　요일(0-7)
```

# 주기별 예제

```bash
# 매분 test.sh 실행
* * * * * /home/script/test.sh
```

```bash
# 매주 금요일 오전 5시 45분에 test.sh 를 실행
45 5 * * 5 /home/script/test.sh
```

```bash
# 매일 매시간 0분, 20분, 40분에 test.sh 를 실행
0,20,40 * * * * /home/script/test.sh
```

```bash
# 매일 1시 0분부터 30분까지 매분 tesh.sh 를 실행
0-30 1 * * * /home/script/test.sh
```

```bash
# 매 10분마다 test.sh 를 실행
*/10 * * * * /home/script/test.sh
```

```bash
# 5일에서 6일까지 2시,3시,4시에 매 10분마다 test.sh 를 실행
*/10 2,3,4 5-6 * * /home/script/test.sh
```

# 크론탭 사용 팁

한 줄에 한 명령어만 쓰자.
\# 을 입력하여 주석을 달자.

# 크론탭 로깅

```bash
* * * * * /home/script/test.sh > /home/script/test.sh.log 2>&1
```

# 크론탭 주기적으로 백업

```bash
50 23 * * * crontab -l > /home/bak/crontab_bak.txt
```
