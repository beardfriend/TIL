https://codesk.tistory.com/102
를 요약했습니다.

OOP란?

Object Oriented Programming

객체 지향 프로그래밍

지향 = Direct = 방향성

Oriented = 중심의 = 기원의 = 근간

객체가 아닌데 객체 쪽으로 가겠다는 의미

OOP로 넘어간지 20년이 훌쩍 넘었다.

OOP는 개념이다
절대법칙, 기술, 필수 아니다.

Anti OOP 움직임도 나타나고 있다.

OOP로 해결할 수 있는 문제가 있었고, 프로그래밍 사회에 큰 기여를 했기 떄문에 배울 의미가 있다.

아래는 절차지향적 프로그래밍의 예이다.

```go
package main

import "fmt"

type Bread struct {
    val string
}

type StrawbrreyJam struct {
    opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
    val string
}

func GetBreads(num int) []*Bread {
    breads := make([]*Bread, num)
    for i := 0; i < num; i++ {
        breads[i] = &Bread{val: "bread"}
    }

    return breads
}

func OpenStrawberryJam(jam *StrawbrreyJam) {
    jam.opened = true
}

func GetOneSpoon(_ *StrawbrreyJam) *SpoonOfStrawberry {
    return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
    bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
    sandwitch := &Sandwitch{}
    for i := 0; i < len(breads); i++ {
        sandwitch.val += breads[i].val + " + "
    }
    return sandwitch
}

func main() {
    // 1. 빵 두개를 꺼낸다.
    breads := GetBreads(2)

    jam := &StrawbrreyJam{}
    // 2. 딸기잼 뚜껑을 연다.
    OpenStrawberryJam(jam)

    // 3. 딸기잼을 한 스푼 퍼서 빵위에 올린다.
    spoon := GetOneSpoon(jam)

    // 4. 딸기잼을 잘 바른다.
    PutJamOnBread(breads[0], spoon)

    // 5. 빵을 덮는다.
    sandwitch := MakeSandwitch(breads)
    // 6. 완성.
    fmt.Println(sandwitch.val)
}

```

# abstraction

```go
package abstraction

import "fmt"

type SloganSayer interface {
	Slogan()
}

// Campaign은 자기 자신의 인스턴스에서 SloganSayer를 받을 수 있음
// Campaign은 또한 SloganSayer 인터페이스를 구현하고 있으므로 SloganSayer이기도 함
// 이는 체이닝시 유용하다
type Campaign struct {
	SloganSayer
}

// SaySlogan은 파라미터로 Campaign 또한 받을 수 있음
func SaySlogan(s SloganSayer) {
	s.Slogan()
}

// Hillary는 SloganSayer 인터페이스를 구현함
// Hillary는 SloganSayer임
type Hillary struct{}

func (h *Hillary) Slogan() {
	fmt.Println("Stronger together.")
}

// Trump는 SloganSayer 인터페이스를 구현함
// Trump는 SloganSayer임
type Trump struct{}

func (t *Trump) Slogan() {
	fmt.Println("Make American great again.")
}

func Exec() {
	hillary := new(Hillary)
	trump := new(Trump)

	h := Campaign{hillary}
	t := Campaign{trump}

	// Slogan의 Hillary와 Trump 구현체는 추상화 되어있다
	// 대신, Campaign은 단지 이게 SloganSayer이고 따라서 Slogan을 호출할 수 있음을 알고있음
	h.Slogan() // "Stronger together."
	t.Slogan() // "Make America great again."

	// SloganSayer를 SaySlogan의 파라미터로 주입할 수 있음
	SaySlogan(hillary) // "Stronger together."
	SaySlogan(trump)   // "Make America great again."

	// h와 t는 또한 Campaign 타입이라는걸 기억하라
	SaySlogan(h) // "Stronger together."
	SaySlogan(t) // "Make America great again."
}

```

# composition

```go
package composition

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	CanSwim   bool
}

// Amy는 Human 타입으로 임베딩 되어있으며 따라서 Human의 메서드 집합에 속하는 메서드를 실행할 수 있음
type Amy struct {
	Human
}

// Alan 또한 Human 타입으로 임베딩 되어있음
type Alan struct {
	Human
}

func (h *Human) Name() {
	fmt.Printf("Hello! My name is %v %v", h.FirstName, h.LastName)
}

func (h *Human) Swim() {
	if h.CanSwim {
		fmt.Println("I can swim!")
	} else {
		fmt.Println("I can not swim.")
	}
}

func Exec() {
	// amy는 Human 타입으로 구성됨
	amy := Amy{
		Human: Human{
			FirstName: "Amy",
			LastName:  "Chen",
			CanSwim:   true,
		},
	}

	// alan은 Human 타입으로 구성됨
	alan := Alan{
		Human: Human{
			FirstName: "Alan",
			LastName:  "Chen",
			CanSwim:   false,
		},
	}

	// Human의 메서드 집합은 Amy 타입으로 전달됨
	amy.Name()  // "Hello! My name is Amy Chen"
	amy.Swim()  // "I can swim!"
	alan.Name() // "Hello! My name is Alan Chen"
	alan.Swim() // "I can't swim"
}

```

# Encapsulation

```go
package encapuslation

import "fmt"

/*
CLASS 기반 OOP에서는 public private을 통해 이뤄진다.
Go에서는 Package수준에서 이뤄진다.
Public은 대문자, private는 소문자로 시작한다.
*/

// Encapsulation 구조체는 이 패키지 밖으로 노출될 수 있음
type Encapsulation struct{}

// Expose 메서드는 패키지 밖을 노출될 수 있음
func (e *Encapsulation) Expose() {
	fmt.Println("AHHHH! I'm exposed!")
}

// hide 메서드는 패키지 내부에서만 사용할 수 있음
func (e *Encapsulation) hide() {
	fmt.Println("Shhhh... this is super secret")
}

// Unhide는 노출되지 않은 hide 메서드를 사용함
func (e *Encapsulation) Unhide() {
	e.hide()
	fmt.Println("...jk")
}
```

# Polymorphism

```go
package polymorphism

import "fmt"

/*
	Go에서는 인터페이스가 암시적으로 충족된다.
	인터페이스 또한 타입이다.
	이 두 문장은 많은 의미를 담고 있다.

	인터페이스가 임시적으로 충족된다 -> 인터페이스의 모든 메서드가 어떤 타입의
	메서드 집합에 모두 포함되어 있을 경우, 해당 타입은 인터페이스를 만족한다. Go에는 implements 키워드가 없다.


	인터페이스는 타입이다 -> 어떤 타입이 한 인터페이스를 만족하면,
	그 타입은 또한 타입이 만족하는 인터페이스에 의해 제한되는 모든 타입을 만족한다.


*/

type SloganSayer interface {
	Slogan()
}

// SaySlogan은 SloganSayer 타입을 파라미터로 받음

func SaySlogan(sloganSayer SloganSayer) {
	sloganSayer.Slogan()
}

// Hillary는 Slogan 함수를 구현함으로써 암묵적으로 SloganSayer 인터페이스를 만족
// 따라서, Hillary도 SloganSayer 타입이다
type Hillary struct{}

func (h *Hillary) Slogan() {
	fmt.Println("Stronger together")
}

type Trump struct{}

func (t *Trump) Slogan() {
	fmt.Println("Make America great again")
}

func Exec() {
	hillary := new(Hillary)
	hillary.Slogan()   // "Stronger together."
	SaySlogan(hillary) // "Stronger together."
	trump := new(Trump)
	SaySlogan(trump) //
}

```
