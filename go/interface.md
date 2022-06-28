- 여러번 읽는 것보다 직접 만들어보는 것이 이해하는 데 더 적은 시간이 걸린다.

내가 만든 예시
```go
// service.go
package services

import "fmt"

type AuthInterface interface {
	Find()
}

type AuthService struct{}

func NewAuth() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Find() {
	fmt.Println("hello")
}
```

```go
//controllers.go 
package controllers

import "template/services"

type AuthCTRL struct{}

func (a AuthCTRL) Auth() {
	var i services.AuthInterface
	svc := services.NewAuth()
	i = svc
	i.Find()
}
```

```go
// main.go
package main

import (
	"template/controllers"
)

func main() {
	ctrl := &controllers.AuthCTRL{}
	ctrl.Auth()
}
```

interface는 껍데기다.
껍데기에 같은 내용을 갖고 있는 구조체를 입혀주면,
구조체에 있는 내용을 사용할 수 있다.


---
아래는 https://brownbears.tistory.com/306 의 글을 정리했다.

구조체 (strcut) = 필드의 집합
interface = 메서드의 집합

interface는 구현해야 하는 메서들의 (prototype)원형을 정의한다.

예제 1.


```go

package main

import "fmt"

type MyInt int // int 형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(r.width, r.height)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = r     // r을 인터페이스 p에 대입
	p.Print() // 10 20: 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출
}
```

인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우가 있다.
함수 파라미터가 interface인 경우 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면 
모두 입력 파라미터로 사용될 수 있다는 것을 의미한다.

```go
package main

import "fmt"

type Duck struct { // 오리(Duck) 구조체 정의
}

func (d Duck) quack() {     // 오리의 quack 메서드 정의
	fmt.Println("꽥~!") // 오리 울음 소리
}

func (d Duck) feathers() { // 오리의 feathers 메서드 정의
	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

type Person struct { // 사람(Person) 구조체 정의
}

func (p Person) quack() {                           // 사람의 quack 메서드 정의
	fmt.Println("사람은 오리를 흉내냅니다. 꽥~!") // 사람이 오리 소리를 흉내냄
}

func (p Person) feathers() { // 사람의 feathers 메서드 정의
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface { // quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feather 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feather 메서드 호출
}

```