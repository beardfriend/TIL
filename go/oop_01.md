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