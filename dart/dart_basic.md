### final

final과 const는 선언할 때 사용한다.

final, const 둘 다 한 번 선언하면 값을 변경할 수 없다.

final은 실행 중에 값이 결정 됨.
const는 컴파일 시에 값이 결정 됨.

예시 1
```
일반 String은 언제든지 바꿀 수 있음
final은 여행을 가기 전에 정하는게 아니라, 여행 중에 결정할 수 있음. 단 결정하고 나서는 변경 불가
const는 먹을 것을 여행을 가기 전에 미리 정한 것, 가서 딴거 먹을래라고 물어보면 안됨
```
예시 2
```
어떤 프로그램이 실행될 때 시간에 대한 로그를 남기고 싶을 때는 DateTime.now()를 사용
-- var log1 = DateTime.now();
실행 시에 시간이 결정되므로 사용 가능
-- final log2 = DateTime.now();
컴파일 시의 시간을 담을 수 없으므로 사용 불가. 실행은 지금이 아니라 내일도 할 수 있음
-- const log3 = DateTime.now();

```

참고 : https://velog.io/@ruinak_4127/Dart-final%EA%B3%BC-const%EC%9D%98-%EC%B0%A8%EC%9D%B4