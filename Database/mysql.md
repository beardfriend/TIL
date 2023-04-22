# DATE_ADD

지금으로부터 하루를 더한 Date.Time이 나온다.
DATE_SUB(NOW() INTERVAL 1 DAY)

# DATE_SUB

지금으로부터 3달을 뺀 Date.Time이 나온다.
DATE_SUB(NOW() INTERVAL 3 MONTH)

# 문법

## CREATE

CREATE TABLE `테이블이름` (
`속성`
`데이터타입`
`NOT NULL`
`UNIQUE`
`DEFAULT 값`
`CHECK (값) //값에 대한 조건 설정 (조건에 괄호)`
`PRIMARY KEY 속성이름(들) //속성 뒤에 PRIMARY KEY를 추가하여 기본키 설정`
`FOREIGN KEY 속성이름 REFERENCES 테이블이름(속성이름)`
`CONSTRAINT 키이름 FOREIGN KEY(키이름) REFERENCES 테이블(키이름)`
`ON UPDATE [NO ACTION | CASCADE | SET NULL | SET DEFAULT]`
`ON DELETE [NO ACTION | CASCADE | SET NULL | SET DEFAULT]`
);

example

```SQL
CREATE TABLE ORDER(
	it int,
	cust_id int NOT NULL,
	book_id int NOT NULL,
	saleprice int DEFAULT 10000 CHECK(price > 1000),
	PRIMARY KEY (id),
	FOREIGN KEY (cust_id) REFERENCES CUSTOMER(id) ON DELETE CASCADE
);
```

## ALTER

테이블의 제약조건이나 속성을 수정하는 명령이다.

ALTER TABLE 테이블이름
[ADD 속성이름 데이터타입]
[DROP COLUMN 속성이름]
[ALTER COLUMN 속성이름 데이터타입]
[ALTER COLUMN 속성이름 [NULL | NOT NULL]]
[ADD PRIMARY KEY(속성이름)];

## DROP

DROP TABLE 테이블이름;

참조 되어지는 테이블도 함께 제거하고 싶은 경우
DROP TABLE department CASCADE CONSTRAINTS;

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
