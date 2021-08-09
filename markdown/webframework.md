# Go Web Framework
본 문서는 [Go 언어 웹 프로그래밍 철저 입문](https://thebook.io/006806/ch01/) 자료를 공부한 뒤 정리한 내용이다.



## 1. Web Framework
웹(web)을 만드는 데 있어서 개발을 할 수 있도록 만들어 놓은 틀

#### 웹 요청 처리 과정
<img src="https://user-images.githubusercontent.com/55284181/125286690-1215a380-e357-11eb-918a-c38c415c7eae.jpg" width="500">

1. 특정 URL이 호출되면 호출된 URL에 매핑된 핸들러가 실행된다.
2. 핸들러에서는 요청한 내용을 분석해서 비즈니스 로직을 실행한다.
3. 비즈니스 로직 수행이 끝나면 그 결과를 응답으로 전송한다.

#### 웹 프레임워크 동작 방식
<img src="https://user-images.githubusercontent.com/55284181/125286687-10e47680-e357-11eb-9537-6af68d41f244.jpg" width="500">

1. 라우터 (router) : 웹 요청을 들어오면 URL 기반으로 특정 핸들러에 전달한다.

2. 컨텍스트 (context) : 웹 요청의 처리 상태를 저장하는 공간이다.

3. 미들웨어 (middleware) : 핸들러 로직을 수행하기 전에 공통으로 수행할 코드 조각이고 재사용이 가능하다. Go Web Application에서 미들웨어는 HTTP 요청을 처리하는 동안 아무데서나 실행할 수 있는 코드이다. 일반적으로 여러 경로에 적용하려는 공통 기능을 캡슐화하는 데 사용된다. 일반적인 사용은 승인, 검증 등이 있다.

4. 핸들러 (handler) : 라우터가 요청한 내용을 분석하여 비즈니스 로직을 실행한다.

5. 렌더러 (renderer) : 핸들러 로직 수행 결과를 다양한 형태(JSON, XML, Html Template 등)로 응답한다.


## 2. Web Application

- 웹 서버를 구동하고 웹 요청을 받아 처리하는 함수
    ```go
    func ListenAndServe(addr string, handler http.Handler) error
    ```

- URL별로 요청을 처리할 핸들러 함수를 등록하는 함수
    ```go
    func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
    ```

- URL별로 요청을 처리할 핸들러를 등록하는 함수
    ```go
    func Handle(pattern string, handler http.Handler)
    ```

- 웹 요청에 대한 응답으로 특정 경로의 파일 내용을 전달하는 함수
    ```go
    func ServeFile(w http.ResponseWriter, r *http.Request, name string)
    ```

#### 마이크로 프레임워크
마이크로 프레임워크는 웹 개발에 필요한 기능을 최소한만 제공하고, 나머지 기능은 다른 라이브러리나 프레임워크를 확장하여 사용하게 한다.


## 3. Cookie and Session

#### HTTP의 특징과 쿠키와 세션을 사용하는 이유
HTTP 프로토콜이 connectionless, stateless한 특성이 있기 때문에 HTTP 프로토콜 환경에서 서버는 클라이언트가 누구인지 확인해야 한다.

- connectionless : 클라이언트가 요청을 한 후 응답을 받으면 그 연결을 끊어 버리는 특징
- stateless : 통신이 끝나면 상태를 유지하지 않는 특징

### 3-1. Cookie
<img src="https://user-images.githubusercontent.com/55284181/128598750-89a70669-061e-4df8-86fb-48d1ad69c38e.png" width="600">


### 3-2. Session



---
### reference
- [Web framework](https://velog.io/@lucasonestar/Web-framework%EC%9B%B9-%ED%94%84%EB%A0%88%EC%9E%84%EC%9B%8C%ED%81%AC)
- [쿠키와 세션 개념](https://interconnection.tistory.com/74)
- [쿠키와 세션의 개념/차이/용도/작동방식](https://devuna.tistory.com/23)
