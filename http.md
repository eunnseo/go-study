# HTTP

HTTP(HyperText Transfer Protocol)는 **웹상에서 클라이언트와 서버 간에 요청/응답으로 데이터를 주고 받을 수 있는 프로토콜**이다.

- 하이퍼텍스트(HTML) 문서를 교환하기 위해 만들어진 protocol(통신 규약)이다.
- HTTP는 TCP/IP 기반으로 되어있다.


## 0. Web
인터넷 상에 정보가 얽혀있는 무형의 정보 네트워크로, 인터넷에서 사용되는 서비스 중 하나이다. 클라이언트와 서버가 상호작용하는 과정으로 웹이 이루어진다.
> 인터넷(Internet)은 컴퓨터가 서로 연결되어 통신을 주고 받는, 컴퓨터끼리의 네트워크를 일컫는 말이다.

<img src="https://user-images.githubusercontent.com/55284181/128598132-17f61245-219d-4bce-bf83-b5f6e4e0ff54.png" width="600" title="web">

- Web Client (Web Browser) : 요청(Request)을 보내는 클라이언트 컴퓨터에 설치된 프로그램이다.
- Web Server : 요청에 응답(Response)하는 서버 컴퓨터에 설치된 웹을 위한 프로그램이다. 일반적으로 데이터를 보내주는 컴퓨터를 의미한다.


## 1. HTTP 통신 방식
- 클라이언트가 HTTP request를 서버에 보내면 서버는 HTTP response를 보내는 구조로 되어 있다.
- HTTP는 Stateless, 즉 state(상태)를 저장하지 않는다. 그래서 만일 여러 요청과응답 의 진행과정이나 데이터가 필요할때는 쿠키나 세션 등등을 사용하게 된다.


## 2. HTTP Request 구조
HTTP request 메세지는 크게 3부분으로 구성된다.

- start line
    ```http
    GET /search HTTP/1.1
    ```
    - HTTP Method : 해당 request가 의도한 action을 정의하는 부분
    - Request target : 해당 request가 전송되는 목표 uri
    - HTTP Version
- headers
    - 해당 request에 대한 추가 정보(addtional information)를 담고 있는 부분
    - Key:Value 값으로 되어있다.
- body
    - 해당 reqeust의 실제 메세지/내용


## 3. HTTP Response 구조
HTTP response도 request와 마찬가지로 크게 3부분으로 구성된다.

- status line
    ```http
    HTTP/1.1 404 Not Found
    ```
    - HTTP Version
    - Status code : 응답 상태를 나타내는 숫자로 된 코드
    - Status text : 응답 상태를 간략하게 설명해주는 부분
- headers
- body


## 4. HTTP Methods
HTTP 요청에 포함되는 HTTP 메소드는 서버가 요청을 수행하기 위해 해야할 행동을 표시하는 용도로 사용한다. HTTP 메소드로는 대표적으로 GET과 POST가 있다.

#### GET
GET은 **서버로부터 정보를 조회하기 위해 설계된 메소드**이다.

GET으로 **서버에게 동일한 요청을 여러 번 전송하더라도 동일한 응답이 돌아와야 한다.** 그러므로 주로 조회를 할 때 사용해야 한다. 예를 들어, 브라우저에서 웹페이지를 열어보거나 게시글을 읽는 등 조회를 하는 행위는 GET으로 요청하게 된다.

- 쿼리스트링

    GET은 요청을 전송할 때 필요한 데이터를 Body에 담지 않고, 쿼리스트링을 통해 전송한다. URL의 끝에 ```?```와 함께 이름과 값으로 쌍을 이루는 요청 파라미터를 쿼리스트링이라고 부른다. 만약, 요청 파라미터가 여러 개이면 ```&```로 연결한다.

- 쿼리스트링을 포함한 URL 샘플

    여기서 요청 파라미터명은 name1, name2이고, 각각의 파라미터는 value1, value2라는 값으로 서버에 요청을 보내게 된다.
    ```
    www.example-url.com/resources?name1=value1&name2=value2
    ```

#### POST
POST는 **리소스를 생성/변경하기 위해 설계**되었기 때문에 GET과 달리 전송해야될 데이터를 HTTP 메세지의 Body에 담아서 전송한다.

POST는 **서버에게 동일한 요청을 여러 번 전송해도 응답은 항상 다를 수 있다.** 이에 따라 POST는 서버의 상태나 데이터를 변경시킬 때 사용된다. 이처럼 POST는 생성, 수정, 삭제에 사용할 수 있지만, 생성에는 POST, 수정은 PUT 또는 PATCH, 삭제는 DELETE가 더 용도에 맞는 메소드라고 할 수 있다.

- HTTP 메세지의 Body는 길이의 제한 없이 데이터를 전송할 수 있다.
- POST 요청은 크롬 개발자 도구, Fiddler와 같은 툴로 요청 내용을 확인할 수 있기 때문에 민감한 데이터의 경우에는 반드시 암호화해 전송해야 한다.


## 5. TCP/IP
데이터 전송을 위한 통신 규약

#### TCP (Transmission Protocol)
- 패킷을 **목적지까지 어떻게 안정적으로 보낼 것인가**에 대해 정의한 프로토콜이다.
- IP로 컴퓨터의 위치를 찾은 다음, 해당 프로토콜을 사용하여 패킷을 전송한다.
- TCP는 연결 확립과 보내진 패킷의 확인, 순서화, 전달 중 손상된 패킷을 복구하는 책임을 진다.

#### IP (Internet Protocol)
- 호스트에서 호스트까지의 통신, 즉 **보내는 컴퓨터에서 받는 컴퓨터까지의 통신을 책임**지는 프로토콜이다.
- 각 장치들을 구분하기 위해 부여한 주소를 IP 주소라고 한다.
- IP 주소는 4개의 숫자로 구성되며 숫자의 크기에 따라 IPv4, IPv6 로 나뉜다.
- IP 주소와 도메인을 매칭시키는 시스템은 DNS라고 한다.

#### TCP/IP 계층
<img src="https://user-images.githubusercontent.com/55284181/111994395-11b08d00-8b5b-11eb-863f-091e24a7b5de.png" width="400" title="TCP/IP 계층">

1. 네트워크 인터페이스 계층 (Network Interface Layer)
    - **Node-To-Node간의 신뢰성 있는 데이터 전송**을 담당하는 계층이다.
    - MAC 주소가 이 계층에서 사용된다.
    - NIC(Network Interface Card)가 있어야만 네트워크 통신이 가능하다.

2. 인터넷 계층 (Internet Layer)
    - **호스트간의 라우팅을 담당**하는 계층이다.
    - 대표적인 프로토콜 : IP, ARP ...

3. 전송 계층 (Transport Layer)
    - **프로세스간의 신뢰성 있는 데이터 전송**을 담당하는 계층이다.
    - process-to-process 전송을 담당하기 위해서는 논리적 주소가 필요한데, process가 사용하는 포트 번호를 그 논리적 주소로 사용한다.
    - 대표적인 프로토콜 : TCP, UDP

4. 응용 계층 (Application Layer)
    - 사용자와 가장 가까운 계층이다.
    - **서버나 클라이언트 응용 프로그램이 이 계층에서 동작**한다.
    - 동작을 위해서는 전송 계층의 주소, 즉 포트번호를 사용한다. http의 경우 포트번호 80번을 사용한다.
    - 대표적인 프로토콜 : HTTP, Telnet, SSH ...


## 6. URL
네트워크 상에서 자원이 어디 있는지를 알려주기 위한 규약

#### URL의 구성 요소
<img src="https://user-images.githubusercontent.com/55284181/111994472-2a20a780-8b5b-11eb-8c9d-d8224abe6e92.png" width="600" title="URL 구성요소">

1. 프로토콜
    - 컴퓨터끼리 네트워크 통신을 할 때 규격이다.
    - 웹을 이용할 때는 HTTP 프로토콜을 이용한다.

2. 호스트 주소
    - 도메인 네임 혹은 IP 주소 등 컴퓨터의 주소를 표시하는 부분이다.

3. 포트 번호
    - 컴퓨터에서 실행되고 있는 수많은 프로세스들의 주소이다.
    - 기본적으로 포트번호를 입력하지 않았을 때는 프로토콜이 가지고 있는 기본 포트번호가 적용된다.
    - HTTP의 경우 80번, HTTPS의 경우 443번의 포트번호가 기본으로 적용된다.

4. 경로
    - 서버 프로그램 내에 짜인 로직으로 가는 영역이다.

5. 쿼리
    - URL에서 추가적인 데이터를 표현할 때 사용된다.
    - query는 Path 뒤에 ```?```를 기점으로 해서 ```key=value```형태로 데이터를 표현하게 된다.



---
### reference
- [웹에서 HTTP란 뭘까?](https://velog.io/@rimu/%EC%9B%B9%EA%B3%BC-%EC%84%9C%EB%B2%84%EC%97%90-%EB%8C%80%ED%95%9C-%EA%B8%B0%EC%B4%88%EC%A7%80%EC%8B%9D-HTTP-%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C)
- [웹이란 무엇인가](https://www.betterweb.or.kr/blog/%EC%9B%B9%EA%B3%BC-%EC%9B%B9-%EA%B2%80%EC%83%89-%EC%9B%B9%EC%9D%B4%EB%9E%80-%EB%AC%B4%EC%97%87%EC%9D%B8%EA%B0%80/)
- [HTTP 구조 및 핵심 요소](https://velog.io/@teddybearjung/HTTP-%EA%B5%AC%EC%A1%B0-%EB%B0%8F-%ED%95%B5%EC%8B%AC-%EC%9A%94%EC%86%8C)
- [GET과 POST의 차이](https://hongsii.github.io/2017/08/02/what-is-the-difference-get-and-post/)
- [GET과 POST의 비교 및 차이](https://mangkyu.tistory.com/17)
- [TCP/IP 프로토콜](https://underground2.tistory.com/5)
- [TCP/IP 계층 기본 개념](https://reakwon.tistory.com/68)
- [URL 구조 이해하기](https://www.grabbing.me/URL-018cdd1bb4b541fab6246569244fcf93)
