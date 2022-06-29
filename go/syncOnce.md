특정 함수를 한 번만 호출해야 될 때 sync.Once를 사용한다.

```go
func (o *Once) Do (f func())
```

한 번만 수행해야 하는 함수를 Do() 메서드의 매개변수로 전달하여
실행하면 여러 고루틴에서 실행한다 해도 해당 함수는 한 번만 수행된다.  

싱글턴 인터페이스에서 사용하면 좋을 듯하다.

```go
type authService struct{}

var singleton AuthService
var once sync.Once

func GetAuthService() AuthService {
	once.Do(func() {
		singleton = &authService{}
	})
	return singleton
}
```

참조 : https://thebook.io/006806/ch05/03/03/