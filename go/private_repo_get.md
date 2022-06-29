# go get private repository

## lib으로 사용할 repository 생성

```go
// go.mod
module bitbucket.org/${workspace}/${repository_name}

go ${version}

require gopkg.in/yaml.v2 v${version} // indirect
```

```go
// mysecret.go

package mysecret

import "fmt"

func SecretProcess() {
	fmt.Println("Running the secret process! version 1.1.1")
}

```

### 비트버켓에 올리기

올릴 때 repository 이름을 **꼭  go.mod 파일에 있는 repo 이름과 일치시켜야 한다.**

## 환경설정

### 비트버켓 URL 우회

```bash
git config --global url."git@bitbucket.org:".insteadOf "https://api.bitbucket.org/"
```

### GoENV에 사용할 Lib URL 등록

```bash
go env -w GOPRIVATE=bitbucket.org/${bucket_name}/${repo_name}
```

go.mod  <module_name>  === GOPRIVATE value

go env -w  === change the default settings of the named enviroment

없을 경우 휘발 됨.

```bash
GOPRIVATE=bitbucket.org/${bucket_name}/${repo_name}
```

# git tag version

### 태그 생성

push할 경우 현재 origin에 태그가 올라감.

```basb
git tag -a v1.1.1 -m "1.1.1 relase"
git push origin v1.1.1
```

### 다운로드 버전

다운로드 시 버전 붙혀주면 됨.

```go
bitbucket.org/${bucket_name}/${repo_name} v1.1.0
```

버전이 없을 경우 → v0.0.0-20200823014737-9f7001d12a5f 이런 식으로 붙음.