# 파일 다운로드 기능

엑셀 파일을 클라이언트가 다운받는 로직이다.

Golang : ( gin, excelize)을 사용했다.

```go
	index := f.NewSheet(cellName)
  f.SetActiveSheet(index)
	excelbufferFile, err := f.WriteToBuffer()
  filename := url.QueryEscape("한글파일명.xlsx")
	c.Header("Content-Disposition", " attachment; filename="+filename+"; filename*=UTF-8''"+filename)
	c.Data(http.StatusOK, "application/octet-stream", excelbufferFile.Bytes())
```

MDN ...
application/octet-stream는 다른 모든 경우를 위한 기본값입니다. 알려지지 않은 파일 타입은 이 타입을 사용해야 합니다. 브라우저들은 이런 파일들을 다룰 때, 사용자를 위험한 동작으로부터 보호하도록 개별적인 주의를 기울여야 합니다

엑셀 파일이면 application/vnd.ms-excel을 사용해도 좋겠다.

크롬에서는 "Content-Disposition", " attachment; filename="한글.xslx" 를 지원하지만,
사파리에서 이렇게 작성할 경우 파일 이름이 깨지기 때문에
문자열을 인코딩한 뒤에 filename\*=UTF-8'' 에 꼭 넣어줘야 한다.
