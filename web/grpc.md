| 기능                    | gRPC               | JSON을 사용하는 RESTAPI      |
| ----------------------- | ------------------ | ---------------------------- |
| **계약**                | 필수(.proto)       | 선택사항 (OpenAPI)           |
| **프로토콜**            | HTTP/2             | HTTP                         |
| **Payload**             | Protobuf(소형,2진) | JSON(대형, 사람읽을 수 있음) |
| **규범**                | 엄격한 사양        | 느슨함, HTTP가 유효합니다    |
| **스트리밍**            | 아니오             | 예                           |
| **보안**                | 전송(TLS)          | 전송(TLS)                    |
| **클라이언트 코드생성** | 예                 | OpenAPI + 타사도구           |

gRPC는 직렬화하여 작은 메세지 페이로드가 발생되기 때문에 제한된 대역폭에서 유용하다.

이진 프레이밍 압축, HTTP/2 간단 보내고 받을 때 효율적  
HTTP/2는 gRPC에만 국한되지 않습니다. JSON을 사용한 HTTP API를 포함하여 다양한 요청 형식에서 HTTP/2를 사용하고 성능 개선을 활용할 수 있습니다.

# 장점

## 코드생성

gRPC 개발의 핵심은 gRPC 서비스 및 메세지의 계약을 정의하는 .proto파일이다.
서버 클라이언트간에 .proto파일을 공유하여 메세지와 클라이언트 코드를 END to END로 생성 가능.
클라이언트를 작성하지 않아도 되므로 많은 서비스를 갖춘 응용 프로그램의 개발 시간이 상당히 절감됩니다.

## 엄격한 사양

REST API는 공식 사양이 없지만, gRPC는 사양이 존재하기 때문에 논쟁의 여지가 없다.

## 권장 시나리오

- 마이크로 서비스
  대기시간이 짧고 처리량이 높은 통신을 위해 설계됨. 효율이 중요한 경량 마이크로 서비스에 적합

- 실시간 통신
  양방향 스트리밍을 위한 뛰어난 기능을 지원. 폴링하지 않고 실시간으로 메세지 푸싱

- 네트워크 제한 환경
