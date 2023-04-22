본 내용은 
https://esbook.kimjmin.net/ 를 보고 정리하였습니다.

# 1. 설치 

brew를 사용한.. 

```bash
brew tap elastic/tap
brew install elastic/tap/elasticsearch-full
brew install elastic/tap/kibana-full
```

# 나중에 정리할 것들..

서버 구성시 왜 3대를 구성해야 하는지
서버 설정 방법 등등


엘라스틱 - 
키바나 - 시각적으로 만들어주고 다양한 서비스와 결합해서 사용 가능.


# 2. 데이터 처리

RESTFUL한 명령어를 지원한다.

키바나에서 http 데브 툴즈를 지원



```json
PUT my_index/_doc/1
{
  "name":"Jongmin Kim",
  "message":"안녕하세요 Elasticsearch"
}
```

Method 문서이름/명령/문서index번호. 형태

```json
GET myindex/_update/1
{
	"doc" :{
		"message" : "hello"
	}
}
```

도큐먼트 필드 1번문서 업데이트

LogStash에서는 대량의 입력을 할 때 Bulk insert를 사용한다
 더 빠르다.

 ```json
 POST _bulk
{"index":{"_index":"test", "_id":"1"}}
{"field":"value one"}
{"index":{"_index":"test", "_id":"2"}}
{"field":"value two"}
{"delete":{"_index":"test", "_id":"2"}}
{"create":{"_index":"test", "_id":"3"}}
{"field":"value three"}
{"update":{"_index":"test", "_id":"1"}}
{"doc":{"field":"value two"}}
 ```



## 2.1.검색

```json
GET test/_search?q=value
GET test/_search?q=field:value
GET test/_search?q=value AND three
```


match쿼리를 가장 많이 사용한다.
```json
GET test/_search
{
  "query": {
    "match": {
      "field": "value"
    }
  }
}

```

### 2.1.1.멀티테넌시

```json
GET logs-2018-01,2018-02,2018-03/_search
GET logs-2018-*/_search
```


# 3. Query DSL


---

## 3.1. Full Text Query

요약
match로 검색
띄어쓰기가 있으면 OR연산자가 들어감
AND 연산자를 쓰고 싶으면 operator필드에 AND를 추가
띄어쓰기를 문자열로 취급하고 싶으면 match_phrase로 검색

---

데이터 시스템에서의 검색은 
수많은 대상 데이터 중에서 조건에 부합하는 범위를 축소하는 행위

```json
POST my_index/_bulk
{"index":{"_id":1}}
{"message":"The quick brown fox"}
{"index":{"_id":2}}
{"message":"The quick brown fox jumps over the lazy dog"}
{"index":{"_id":3}}
{"message":"The quick brown fox jumps over the quick dog"}
{"index":{"_id":4}}
{"message":"Brown fox brown dog"}
{"index":{"_id":5}}
{"message":"Lazy jumping dog"}
```

검색은 match가 가장 기본

### 3.1.1. match
메세지 필드에 개가 들어간 것들을 검색.
```json
GET my_index/_search
{
  "query": {
    "match": {
      "message": "dog"
    }
  }
}
```
띄어쓰기로 검색어를 추가할 수 있는데
OR연산자가 기본으로 설정 됨.
```json
GET my_index/_search
{
  "query": {
    "match": {
      "message": "dog quick"
    }
  }
}
```
AND 연산자로 하고 싶다면

```json
GET my_index/_search
{
  "query": {
    "match": {
      "message": "dog quick",
			"operator": "and"
    }
  }
}
```
### 3.1.2. match_phrase
띄어쓰기를 정확하게 검색하고 싶을 땐 match_phrase

```json
GET my_index/_search
{
  "query": {
    "match_phrase": {
      "message": "lazy dog"
    }
  }
}
```

slop은 띄어쓰기 사이에 몇 개의 단어를 끼울 지 정한다.

```json
GET my_index/_search
{
  "query": {
    "match_phrase": {
      "message": {
        "query": "lazy dog",
        "slop": 1
      }
    }
  }
}
```

이 때는 lazy jumping dog lazy sexy dog
이런 단어들이 걸러지고

slop이 2일 경우 lazy jumping sexy dog 
단어 사이에 2개가 들어간다.

slop은 1이상은 사용하지 않기를 권장


### 3.1.3. query_string
```json
GET my_index/_search
{
  "query": {
    "query_string": {
      "default_field": "message",
      "query": "(jumping AND lazy) OR \"quick dog\""
    }
  }
}
```
lazy jumping을 모두 포함하거나 quick dog 구문을 포함하는 도큐먼트



## 3.2. Bool 복합쿼리 

must : 쿼리가 참인 도큐먼트
must_not : 쿼리가 거짓인 도큐먼트
should : 검색 결과 중 이 쿼리에 해당하는 도큐먼트의 점수를 높입니다. 
filter : 쿼리가 참인 도큐먼트를 검색하지만 스코어를 계산하지 않습니다. must 보다 검색 속도가 빠르고 캐싱이 가능합니다.

example

```json
GET my_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "message": "quick"
          }
        },
        {
          "match_phrase": {
            "message": "lazy dog"
          }
        }
      ]
    }
  }
}
```


![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LnkDY-4MZD1V45rlVio%2F-LnkDaO7ZsRAg85vnPxB%2F5.2-01.png?alt=media&token=bfe13e1d-01e6-4d01-b36c-9112d3ad4a7d)
출처 https://esbook.kimjmin.net/

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LnkDY-4MZD1V45rlVio%2F-LnkDo_R96o_gyd7QWJk%2F5.2-02.png?alt=media&token=69136151-e064-4887-b9ec-90ecd422a0d3)
출처 https://esbook.kimjmin.net/

## 3.3. 정확도 Relavancy

요약 

점수가 높으면 검색결과 상단에 올라온다.

점수가 높을 수 있는 조건 
1. 문서에 단어가 많이 포함될 경우 (최대 25개까지 점수에 반영) TF (Term Frequnecy)
2. 중요한 단어 희소성 있는 단어가 포함될 경우 점수가 높다. IDF (Inverse Document Frequency) 
ex) 코딩하는 사람  "코딩하는"이 "사람"보다 전체문서에서 더 적게 포함되기 때문에 높은 점수를 준다. 
3. 내용이 짧을수록 높은 점수를 준다. (Field Length)
내용이 짧을수록 제목일 가능성이 높기 때문에
ex) 블로그 제목, 내용이 모두 같은 필드에 있으면,
제목이 더 짧으니 제목처럼 생긴 녀석이 위로 올라옴.

---

## 3.4. Bool : Should

```json 
GET my_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "message": "fox"
          }
        }
      ],
      "should": [
        {
          "match": {
            "message": "lazy"
          }
        }
      ]
    }
  }
}
```
fox를 검색할 때, sholud에 키워드에 더 높은 점수를 부여
```json
GET my_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "message": {
              "query": "lazy dog"
            }
          }
        }
      ],
      "should": [
        {
          "match_phrase": {
            "message": "lazy dog"
          }
        }
      ]
    }
  }
}
```
이렇게 should와 match_phrase를 응용하면 쇼핑몰에서 "스키 장갑" 같은 단어로 검색했을 때 스키 용품들과 각종 장갑들을 모두 가져오면서 그 중 스키 장갑을 가장 상위에 표시할 수 있습니다. slop:1을 이용하면 "스키 보드 장갑", "스키 벙어리 장갑" 같이 스키와 장갑 사이에 다른 값이 들어간 결과에도 가중치를 부여할 수 있습니다.

## 3.5. 정확값 쿼리 - Exact Value Query

**검색 조건의 참/거짓을 판별해 결과를 가져오는 것도 가능하다.**
이 특성을 *정확값* Exact Value라고 부른다.
exact value에는 term, range가 있으며,
스코어를 계산하지 않는다.

스코어에 영향을 주지 않는다.

스코어에 영향을 주지 않으면, 검색 결과를 그대로 반영하고,
필터는 필터만으로 작용을 할 수 있다.
### 3.5.1. bool:filter

bool쿼리의 filter 안에 하위 쿼리를 사용하면 스코어에 영향을 주지 않습니다. 

```json
GET my_index/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "message": "fox"
          }
        }
      ],
      "filter": [
        {
          "match": {
            "message": "quick"
          }
        }
      ]
    }
  }
}

```

### 3.5.2. keyword

```json
GET my_index/_search
{
  "query": {
    "bool": {
      "filter": [
        {
          "match": {
            "message.keyword": "Brown fox brown dog"
          }
        }
      ]
    }
  }
}

```

정확히 Brown fox brown dog 에 일치하는 결과만 리턴한다.

# 3.6. 범위쿼리

elastic은 숫자, 날짜 형식의 형태 저장이 가능하다.
이때 range쿼리를 사용한다.

gte (Greater-than or equal to) - 이상 (같거나 큼)
gt (Greater-than) – 초과 (큼)
lte (Less-than or equal to) - 이하 (같거나 작음)
lt (Less-than) - 미만 (작음)

```json
GET phones/_search
{
  "query": {
    "range": {
      "price": {
        "gte": 700,
        "lt": 900
      }
    }
  }
}
```

```json
GET phones/_search
{
  "query": {
    "range": {
      "date": {
        "gt": "2016-01-01"
      }
    }
  }
}

```

range는 기본적으로 score계산을 하지 않는다

가격이 900원에 가까울수록 점수가 높은 게 아니다.

오직 true/false로 구분


# 4.데이터 색인과 텍스트 분석

풀 텍스트 검색을 하기 위해서는 **데이터를 검색에 맞게 가공**하는 작업을 필요로 하는데 Elasticsearch는 **데이터를 저장하는 과정**에서 이 작업을 처리합니다. 이번 장에서는 Elasticsearch가 검색을 위해 텍스트 데이터를 어떻게 처리하고 데이터를 색인 할 때 Elasticsearch에서 어떤 과정이 이루어지는지에 대해 살펴보겠습니다.

# 4.1. 역인덱스

전통적인 RDMBS에서는 데이터가 늘어날수록 검색해야될 대상이 늘어날수록, 시간이 오래걸린다. 그리고 모든 rows를 읽어야 하기 떄문에, 기본적으로 속도가 느리다.

엘라스틱은 `역인덱스`라는 구조를 만들어 저장한다.


![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntL_BGpuFbNXy_sFtK%2F-LntLbibpXHABupWvXtu%2F6.1-03.png?alt=media&token=d2726f20-a7ea-4219-bcb0-340cbe1d21f1)
출처 https://esbook.kimjmin.net/


  이 역 인덱스는 책의 맨 뒤에 있는 주요 키워드에 대한 내용이 몇 페이지에 있는지 볼 수 있는 찾아보기 페이지에 비유할 수 있습니다. Elasticsearch에서는 추출된 각 키워드를 텀(term) 이라고 부릅니다. 이렇게 역 인덱스가 있으면 fox를 포함하고 있는 도큐먼트들의 id를 바로 얻어올 수 있습니다.
	  Elasticsearch는 데이터가 늘어나도 찾아가야 할 행이 늘어나는 것이 아니라 역 인덱스가 가리키는 id의 배열값이 추가되는 것 뿐이기 때문에 큰 속도의 저하 없이 빠른 속도로 검색이 가능합니다. 이런 역 인덱스를 데이터가 저장되는 과정에서 만들기 때문에 Elasticsearch는 데이터를 입력할 때 저장이 아닌 색인을 한다고 표현합니다.


# 4.2. 텍스트 분석 Text analysis

엘라스틱서치는 문자열 필드가 저장될 때 검색어 토큰을 저장하기 위해 여러 단계를 거친다.

이 전체과정을 텍스트 분석(text analysis)이라고 한다.

이 과정을 처리하는 기능을 애널라이저(Analyzer)라고 한다.

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntYrdKmTe441TqYAJl%2F-LntZ63SAIfHu6Q_OgzJ%2F6.2-02.png?alt=media&token=52213afe-e6ab-4bc2-b9e0-20027542a79e)
출처 https://esbook.kimjmin.net/

캐릭터 필터 : 0~3개
토크나이저 1개
토큰필터 : 0 ~ N개

캐릭터필터는 특수문자를 제거하거나, HTML tag를 제거한다거나 등의 과정을 처리할 수 있다.

토크나이저는 쪼개는 작업을 담당한다.
whitespace는 공백기준으로 쪼개는 작업을 할 수 있다.

이제 필터에 들어오면

lowercase는 문자열에서 대문자를 소문자로 바꿔 대소문자 구분없이 검색할 수 있게 도와주고

불용어 stopword by, of,are 이런 것들을 날리는 필터에 들어가면 다 날라간다.

snowball필터는 영어쓸 때 많이 사용하는데
복수형이라던가 ing가 붙은 것들 일반형태로 만들어준다.

synonym는 동의어를 추가해준다.

# 4.3. 애널라이저 Analyzer

### 4.3.1. _analyze API


```json
GET _analyze
{
  "text": "The quick brown fox jumps over the lazy dog",
  "tokenizer": "whitespace",
  "filter": [
    "lowercase",
    "stop",
    "snowball"
  ]
}
```

토큰 필터, 토크나이저가 적용된다.
필터의 순서가 중요하다.
lowercase보다 stop이 먼저오면 "The"는 불용어로 간주되지 않아
역 색인에 남게 된다.



```json
GET _analyze
{
  "text": "The quick brown fox jumps over the lazy dog",
  "analyzer": "snowball"
}
```
analzyer 항목으로 적용 가능하다.
snowball은 whitespace + lowercase +stop +snowball 을 합친 애널라이저다.
엘라스틱서치에서 정의해놓았기 때문에 바로 사용가능하다.

직접 애널라이즈를 정의할 수도 있다.

```json
PUT my_index2
{
  "mappings": {
    "properties": {
      "message": {
        "type": "text",
        "analyzer": "snowball"
      }
    }
  }
}
```
my_index2의 message field에 snowball 애널라이즈 적용


### 4.3.2. Term 쿼리

엘라스틱에는 term쿼리가 존재한다.
match와다른 점은 애널라이즈를 적용하지 않고 검색한다.


### 4.3.3. 사용자 정의 애널라이저 - Custom Analyzer

_analyze API로 애널라이저, 토크나이저, 토큰필터들의 테스트가 가능하지만, 실제로 인덱스에 저장되는 데이터의 처리에 대한 설정은 애널라이저만 적용할 수 있습니다

```json
PUT my_index3
{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "my_custom_analyzer": {
            "type": "custom",
            "tokenizer": "whitespace",
            "filter": [
              "lowercase",
              "stop",
              "snowball"
            ]
          }
        }
      }
    }
  }
}
```

```json
PUT my_index3
{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "my_custom_analyzer": {
            "type": "custom",
            "tokenizer": "whitespace",
            "filter": [
              "lowercase",
              "my_stop_filter",
              "snowball"
            ]
          }
        },
        "filter": {
          "my_stop_filter": {
            "type": "stop",
            "stopwords": [
              "brown"
            ]
          }
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "message": {
        "type": "text",
        "analyzer": "my_custom_analyzer"
      }
    }
  }
}

```

매핑에 넣으면 PUT할 때 텍스트 필드는 my_custom_analyzer가 적용된다.

즉 검색할 때 brwon이라 검색하면 

```json
GET my_index3/_search
{
  "query": {
    "match": {
      "message": "brown"
    }
  }
}
```
응답 결과가 없다.


### 4.3.4. 텀 벡터 - _termvectors API

```json
GET my_index3/_termvectors/1?fields=message
```

my_index3/_doc1/1 번의 message 필드의 termvectors확인

## 4.4. 캐릭터 필터 - Character Filter

캐릭터필터는 텍스트 분석 중 가장 먼저 처리되는 과정으로
색인된 텍스트가 토크나이저에 의해 텀으로 분리되기 전에
전체 문장에 대해 적용되는 일종의 전처리 도구입니다.
이 책에서 설명하는 7.0 버전 기준으로 캐릭터 필터는 
HTML Strip, Mapping, Pattern Replace 총 3개가 존재합니다. 
char_filter 항목에 배열로 입력하며 하나만 적용하거나 
차례대로 입력해서 3개를 모두 적용할 수도 있습니다.

## 4.4.1. HTML Strip

HTML 구문 분석

### 4.4.2. Mapping

C++같은 경우 ++는 불용어인데
대부분의 애널라이저들은 +를 불용어취급을 한다.
따라서 +를 다른 단어로 치환해서 색인을 해야 하는데
이 때 사용하는 것이 mapping이다.

```json
PUT coding
{
  "settings": {
    "analysis": {
      "analyzer": {
        "coding_analyzer": {
          "char_filter": [
            "cpp_char_filter"
          ],
          "tokenizer": "whitespace",
          "filter": [ "lowercase", "stop", "snowball" ]
        }
      },
      "char_filter": {
        "cpp_char_filter": {
          "type": "mapping",
          "mappings": [ "+ => _plus_", "- => _minus_" ]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "language": {
        "type": "text",
        "analyzer": "coding_analyzer"
      }
    }
  }
}
```

### 4.4.3. Pattern Replace

 Pattern Replace 캐릭터 필터는 정규식(Regular Expression)을 이용해서 
 좀더 복잡한 패턴들을 치환할 수 있는 캐릭터 필터입니다. 
 다음은 카멜 표기법(camelCase)으로 된 단어를 대문자가 시작하는 단위 마다 
 공백을 삽입하여 세부 단어별로 토크나이징 될 수 있도록 
 camel 인덱스에 camel_analyzer 라는 애널라이저를 생성하는 예제입니다.

## 4.5. 토크나이저 - Tokenizer

토크나이저들 중 NGram, Lowercase 같은 토크나이저들은
대부분은 Standard 토크나이저에 같은 이름의 토큰 필터를 내장한 들입니다.
이 책에서 다루지 않는 토크나이저들은 공식 홈페이지의 도큐먼트를 확인하시기 바랍니다.
### 4.5.1. Standard, Letter, Whitespace

보통은 Standard를 많이 사용한다.
```json
GET _analyze
{
  "tokenizer": "letter",
  "text": "THE quick.brown_FOx jumped! @ 3.5 meters."
}
```
아래와 같이 사용.

### 4.5.2 UAX URL Email

주로 사용되는 Standard는 `@` `/`같은 특수문자를 제거하고 분리한다.

uax_url_email토크나이저는 이메일이나 주소를 분리하지 않고 그대로 저장한다.

### 4.5.3. 패턴 

일반적인 토크나이저는 공백을 기준으로 분리하지만,
패턴 같은 경우에는  분리할 패턴을 기호 또는 Java 정규식 형태로 지정할 수 있습니다

### 4.5.4. Path Hierarchy

디렉토리 경로를 분석할 때 상위경로를 포함하여 저장한다.

```json
{
  "tokens" : [
    {
      "token" : "/usr",
      "start_offset" : 0,
      "end_offset" : 4,
      "type" : "word",
      "position" : 0
    },
    {
      "token" : "/usr/share",
      "start_offset" : 0,
      "end_offset" : 10,
      "type" : "word",
      "position" : 0
    },
    {
      "token" : "/usr/share/elasticsearch",
      "start_offset" : 0,
      "end_offset" : 24,
      "type" : "word",
      "position" : 0
    },
    {
      "token" : "/usr/share/elasticsearch/bin",
      "start_offset" : 0,
      "end_offset" : 28,
      "type" : "word",
      "position" : 0
    }
  ]
}
```

## 4.6. 토큰필터

토크나이저를 이용하여 텀을 분리한 뒤에
분리한 텀들을 지정한 규칙에 따라 처리를 해준다.
이 과정을 담당하는 것이 토큰필터다.

### 4.6.1. Lowercase, Uppercase

대소문자 구분없이 처리해주는 필터

### 4.6.2. Stop (불용어)

블로그 포스트나 뉴스 기사 같은 글에는 검색에서는 큰 의미가 없는 조사나 전치사 등이 많습니다. 
영문에서도 the, is, a 같은 단어들은 대부분 검색어로 쓰이지 않는데 
이런 단어를 한국어로는 불용어, 영어로는 stopword라고 합니다. 
Stop 토큰 필터를 적용하면 불용어에 해당되는 텀 들을 제거합니다.

### 4.6.3 Synonym (동의어)

검색 서비스에 따라서 동의어를 제공해야 하는 경우가 있다.
예를들면 AWS를 검색했을 때, Amazon, 아마존도 검색을 하도록 하면
더 많은 정보를 찾을 수 있다.
Synonym필터를 사용하면 텀의 동의어 저장이 가능하다.


동의어를 설정하는 옵션은 synonyms 항목에서 직접 동의어 목록을 입력하는 방법과 
동의어 사전 파일을 만들어 synonyms_path 로 지정하는 방법이 있습니다. 
동의어 사전 명시 규칙에는 다음의 것들이 있습니다.

"A, B => C" : 왼쪽의 A, B 대신 오른쪽의 C 텀을 저장합니다. 
A, B 로는 C 의 검색이 가능하지만 C 로는 A, B 가 검색되지 않습니다.

"A, B" : A, B 각 텀이 A 와 B 두개의 텀을 모두 저장합니다. 
A 와 B 모두 서로의 검색어로 검색이 됩니다.

```json

{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_syn": {
          "tokenizer": "whitespace",
          "filter": [
            "lowercase",
            "syn_aws"
          ]
        }
      },
      "filter": {
        "syn_aws": {
          "type": "synonym",
          "synonyms": [
            "amazon => aws"
          ]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "message": {
        "type": "text",
        "analyzer": "my_syn"
      }
    }
  }
}
```