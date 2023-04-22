트라이 캐치처럼 사용하기

```go
package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestError(param string) (string, error) {
	if param == "1" {
		return "1", nil
	}
	return "2", errors.New("ㅁㄴㅇ")
}

func Service() string {
	var result string
	if val,err := TestError("2"); err !=nil {
		if err.Error() == "error" {
		panic("레코드가 없습니다")
	}
	panic(err)
	} else {
		result = val
	}
	return result
}

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				str:= fmt.Sprintf("%v",err)
				if str != "" {
					c.JSON(500,gin.H{"message": str})
					return
				}
				c.JSON(500,gin.H{"message": "서버 error"})
			}
			}()

		var result = Service()
		c.JSON(200, gin.H {
			"message": result,
		})


	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
