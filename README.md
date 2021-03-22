# go_study

# 1. HTTP 프로토콜
> HyperText Transfer Protocol
> * 하이퍼텍스트(HTML) 문서를 교환하기 위해 민들어진 protocol(통신 규약)
> * 클라이언트가 HTTP request를 서버에 보내면 서버는 HTTP response를 보내는 구조
> * HTTP는 TCP/IP 기반으로 되어있다.

### TCP/IP
* 인터넷에서 컴퓨터들이 서로 정보를 주고받는 데 쓰이는 통신규약(프로토콜)의 모음
* 패킷 통신 방식의 인터넷 프로토콜인 IP와 전송 조절 프로토콜인 TCP로 이루어져 있다.   
* **TCP**
    - 패킷을 목적지까지 어떻게 안정적으로 보낼 것인가에 대해 정의한 프로토콜
    - IP로 컴퓨터의 위치를 찾은 다음, 해당 프로토콜을 사용하여 패킷을 전송
    - TCP는 연결 확립과 보내진 패킷의 확인, 순서화, 전달 중 손상된 패킷을 복구하는 책임을 짐
* **IP**
    - 인터넷 프로토콜(Internet Protocol)의 약자
    - 호스트에서 호스트까지의 통신, 즉 보내는 컴퓨터에서 받는 컴퓨터까지의 통신을 책임지는 프로토콜
    - 각 장치들을 구분하기 위해 IP 주소를 부여함
    - IP 주소는 4개의 숫자로 구성되며 숫자의 크기에 따라 IPv4, IPv6 로 나뉨
* **TCP/IP 계층**
    1. 네트워크 인터페이스 계층
        - Node-To-Node간의 신뢰성 있는 데이터 전송을 담당하는 계층
        - OSI7 계층의 물리 계층과 데이터링크 계층의 역할 담당
    2. 인터넷 계층 (Internet Layer)
        - 호스트간의 라우팅을 담당하는 계층
        - OSI7 계층의 네트워크 계층의 역할 담당
        - 대표적인 프로토콜 : IP, ARP ...
    3. 전송 계층 (Transport Layer)
        - 프로세스간의 신뢰성 있는 데이터 전송을 담당하는 계층
        - OSI7 계층의 전송계층의 역할 담당
        - 대표적인 프로토콜 : TCP, UDP
    4. 응용 계층 (Application Layer)
        - 사용자와 가장 가까운 계층
        - 서버나 클라이언트 응용 프로그램이 이 계층에서 동작
        - OSI7계층의 5계층부터 7계층까지의 기능을 담당
        - 대표적인 프로토콜 : HTTP, Telnet, SSH ...
![TCP/IP 계층](tcp_ip_protocol.png)

### HTTP header
* 해당 request에 대한 추가 정보(addtional information)를 담고 있는 부분
* Key:Value 값으로 되어있다 (: 이 사용됨)

* **자주 사용되는 headers**
    - Host : 요청이 전송되는 target의 host url
    - User-Agent : 요청을 보내는 클라이언트의 대한 정보
    - Accept : 해당 요청이 받을 수 있는 응답(response) 타입
    - Connection : 해당 요청이 끝난후에 클라이언트와 서버가 계속해서 네트워크 컨넥션을 유지 할것인지 아니면 끊을것인지에 대해 지시하는 부분
    - Content-Type : 해당 요청이 보내는 메세지 body의 타입
    - Content-Length : 메세지 body의 길이

### HTTP Body
* 해당 reqeust의 실제 메세지/내용
* GET request들은 body가 없는 경우가 많음

### HTTP Method
* 해당 request가 의도한 action을 정의하는 부분
* HTTP Methods에는 GET, POST, PUT, DELETE, OPTIONS 등등이 있다.

* **자주 사용되는 HTTP Methods**
    - GET : 데이터를 서버로부터 받아올 때 주로 사용하는 Method. 데이터 생성/수정/삭제 없이 받아오기만 할때 사용됨
    - POST : 데이터를 생성/수정/삭제 할때 주로 사용되는 Method. 대부분의 경우 request body가 포함되서 보내짐

* **알고 있으면 좋은 HTTP Methods**
    - OPTIONS : 주로 요청 URI에서 사용할 수 있는 Method를 받아올때 사용
    - PUT : POST와 비슷하다. 데이터를 생성 할때 사용되는 Method
    - DELETE : 특정 데이터를 서버에서 삭제 요청을 보낼때 쓰이는 Method

### URL
* 네트워크 상에서 자원이 어디 있는지를 알려주기 위한 규약
* **URL의 구성 요소**
    1. 프로토콜
        - 컴퓨터끼리 네트워크 통신을 할 때 규격
        - 웹을 이용할 때는 HTTP 프로토콜을 이용
    2. 호스트 주소
        - 도메인 네임 혹은 IP 주소 등 컴퓨터의 주소 표시
    3. 포트 번호
        - 컴퓨터에서 실행되고 있는 수많은 프로세스들의 주소
        - 기본적으로 포트번호를 입력하지 않았을 때는 프로토콜이 가지고 있는 기본 포트번호가 적용
    4. 경로
        - 서버 프로그램 내에 짜인 로직으로 가는 영역
    5. 쿼리
        - URL에서 추가적인 데이터를 표현할 때 사용
![URL 구성요소](url.png)

# 2. 웹 프레임워크
> * 

### 라우터

### 핸들러, 미들웨어

### 세션, 쿠키

### Gin
* Gin은 Golang으로 작성된 웹 프레임워크이다.

***
### 참조링크
* <https://velog.io/@teddybearjung/HTTP-%EA%B5%AC%EC%A1%B0-%EB%B0%8F-%ED%95%B5%EC%8B%AC-%EC%9A%94%EC%86%8C>
* <https://ko.wikipedia.org/wiki/%EC%9D%B8%ED%84%B0%EB%84%B7_%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C_%EC%8A%A4%EC%9C%84%ED%8A%B8>
* <https://underground2.tistory.com/5>
* <https://reakwon.tistory.com/68>
* <https://www.grabbing.me/URL-018cdd1bb4b541fab6246569244fcf93>
