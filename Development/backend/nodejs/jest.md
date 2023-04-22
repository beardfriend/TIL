## Mock Function

실제 객체를 만들기엔 비용과 시간이 많이 들거나 의존성이 길게 걸쳐져 있어  
제대로 구현하기 어려울 경우 가짜객체(Mock)을 이용한다.

### 언제 필요한가?

- 테스트 작성을 위한 환경 구축이 어려운 경우
- 테스트가 특정 경우나 순간에 의존적인 경우
- 테스트 시간이 오래 걸리는 경우
- 개인 PC나 서버의 성능문제로 오래 걸릴 수 있는 경우 시간을 단축하기 위헤

### jest.fn()

```javascript
const mockFn = jest.fn();
```

- Return value 지정 가능 (mockReturnValue)
- Param을 받는 함수 생성 가능 (mockImplementation)
- 비동기 함수 생성가능 (mockResolvedValue, mockRejectedValue)

### jest.mock()

jest.fn()이 여러개라 한번에 묶고 싶을 때 사용

```javascript
export const add = jest.fn();
export const subtract = jest.fn();
export const multiply = jest.fn();
export const divide = jest.fn();
```

```javascript
import * as app from './app';
import * as math from './math';

// Set all module functions to jest.fn
jest.mock('./math.js');

test('calls math.add', () => {
  app.doAdd(1, 2);
  expect(math.add).toHaveBeenCalledWith(1, 2);
});

test('calls math.subtract', () => {
  app.doSubtract(1, 2);
  expect(math.subtract).toHaveBeenCalledWith(1, 2);
});
```

참조 : https://dev.to/dylanju/jest-mocks-18l9

### jest.spyOn()

메소드가 실행되는 것뿐만 아니라 기존의 구현 값을 보고 싶을 때,

```javascript
import * as app from './app';
import * as math from './math';

test('calls math.add', () => {
  const addMock = jest.spyOn(math, 'add');

  // calls the original implementation
  expect(app.doAdd(1, 2)).toEqual(3);

  // and the spy stores the calls to add
  expect(addMock).toHaveBeenCalledWith(1, 2);
});
```
