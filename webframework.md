# Go Web Framework
본 문서는 [Go 언어 웹 프로그래밍 철저 입문](https://thebook.io/006806/ch01/) 자료를 공부한 뒤 정리한 내용이다.



## 1. 웹 프레임워크

+ **웹 요청 처리 과정**

    <img src="https://user-images.githubusercontent.com/55284181/125286690-1215a380-e357-11eb-918a-c38c415c7eae.jpg" width="500">

    1. 특정 URL이 호출되면 호출된 URL에 매핑된 핸들러가 실행된다.
    2. 핸들러에서는 요청한 내용을 분석해서 비즈니스 로직을 실행한다.
    3. 비즈니스 로직 수행이 끝나면 그 결과를 응답으로 전송한다.

+ **웹 프레임워크 동작 방식**

    <img src="https://user-images.githubusercontent.com/55284181/125286687-10e47680-e357-11eb-9537-6af68d41f244.jpg" width="500">

    1. 라우터 : 웹 요청을 들어오면 URL 기반으로 특정 핸들러에 전달한다.
    2. 컨텍스트 : 웹 요청의 처리 상태를 저장하는 공간이다.
    3. 미들웨어 : 핸들러 로직을 수행하기 전에 공통으로 수행할 코드 조각이고 재사용이 가능하다. 미들웨어는 로그 처리, 에러 처리, 정적 파일 처리, 사용자 인증과 권한 관리, 보안 처리, 세션 상태 관리, 웹 요청 정보 파싱과 같은 기능을 처리한다.
    4. 렌더러 : 핸들러 로직 수행 결과를 다양한 형태(JSON, XML, Html Template 등)로 응답한다.


## 2. 웹 애플리케이션

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
