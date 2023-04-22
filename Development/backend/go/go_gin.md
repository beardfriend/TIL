### Bind

Bind는 묶음이라는 뜻을 지닌다.
Request를 묶음으로 받을 수 있다.
Requset Header, Request Param, Request Query, Request Body 를 받을 수 있다.

- Bind Option

```go
type HeaderParameter struct {
    XRequestID string `header:"X-Request-ID"`
}

type Body struct {
   FirstName string `json:"firstName" binding:"required"`
   LastName string `json:"lastName" binding:"required"`
   Email string `json:"email" binding:"required,email"`
   Phone string `json:"phone" binding:"required,e164"`
   CountryCode string `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

```

헤더 옵션을 설정할 수도 있고,
binding:"required"를 표시하면 Bind할 때 값이 없으면 에러를 던져준다.
또한 메서드의 최대 최소길이 형태가 이메일인지 등을 확인할 수 있다.

자세한 옵션을 보려면
[클릭](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples)

- ShoudBind vs Bind

Bind는 바로 400에러를 던지지만, ShoudBind는 그렇지 않다.

```go
req := Body{}

c.Bind(&req)
  // 에러 처리를 안 해도 에러가 날 경우 400을 던짐.

if err := c.ShouldBind(&req) {
  //에러 처리를 꼭 해줘야 함.
  panic(err)
}
```

ShoudBind에는

ShouldBindJson
SholudBindUri
SholdBindHeader 등이 존재한다.

Custom Validator를 작성할 수도 있다.
params, json, query, header를 struct으로 받으면 장점은
타입 변환이 되고, validating 기능이 잘 되기 때문에,
좀더 명확한 API가 될 수 있다.