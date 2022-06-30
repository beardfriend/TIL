# MySQL Replication

MySQL Replication은 말 그대로 DB안 데이터를 갖다가 물리적으로 복사해
다른 곳에 넣어두는 기술을 말한다.

읽기 처리 ( 동시성이 꼭 필요하지 않은 경우는 )
Slave DB에 부하를 분산시킨다.

## 복제 메커니즘

1. Master DB에 변경이 일어난다.
2. Master DB에 반영한다
3. 변경 이력을 Binary Log로 저장한다
4. 관련 이벤트를 Slave DB들에게 넘긴다.
5. Slave IO Thread에서 이벤트를 캐치하면 
6. Binary Log를 Slave DB 각각의 Relay Log에 저장한다
7. Slave SQL Thread에서 Relay Log를 읽어
8. Slave DB를 업데이트한다