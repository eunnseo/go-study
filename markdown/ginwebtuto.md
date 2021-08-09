# web-tuto-with-gin
본 문서는 nclab 커리큘럼 웹서버-실습-3 ~ 웹서버-실습-5 내용을 공부하고 정리한 문서이다.

[소스코드](https://github.com/eunnseo/web-tuto-with-gin)


## 1. Gin

Gin은 Web Application과 Microservices를 만드는 데 사용되는 고성능 Micro-Framework이다.

#### Handler

- render : http 요청의 헤더를 조회하여 json과 html 두가지 형태로 응답하게 하는 함수이다.
    ```go
    func render(c *gin.Context, data gin.H, templateName string) {
        switch c.Request.Header.Get("Accept") {
        case "application/json":
            // Respond with JSON
            c.JSON(http.StatusOK, data["payload"])
        default:
            // Respond with HTML
            c.HTML(http.StatusOK, templateName, data)
        }
    }
    ```

- gin.Context : 하나의 요청을 처리하는 모든 핸들러에서 함수 인자로 사용하는 변수 타입이다.
    + http 요청을 처리하는 일련의 과정에서 key:value 형태로 값을 저장하고 조회할 수 있다.
    + Go routine의 생성, 중단등 flow control 의 역할도 수행한다.
    + "context", "context.Context"는 Golang에서 지원하는 유용한 패키지이다. gin.Context 는 해당 패키지를 활용한 타입이다.

- gin.H : ```Type H map[string]interface{}```와 같다.

- payload : 전송되는 데이터를 의미한다. 바디(body)와 같다.

#### view
Go template : 대부분의 서버사이드 언어들은 정적인 페이지에 스크립팅을 지원하기 위한 언어들을 제공한다. JSP나 PHP 스크립팅 같은 것이 그런 예이다. Go에서는 이와 유사한 스크립팅 언어로 template 패키지를 제공한다.

#### main.go

```go
func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.LoadHTMLGlob("view/*")

    r.GET("/", handler.ShowIndexPage)
    article := r.Group("/article")
    {
        article.GET("/view/:article_id", handler.GetArticle)
        article.GET("/create", handler.ShowArticleCreationPage)
        article.POST("/create", handler.CreateArticle)
        article.GET("/delete/:article_id", handler.DeleteArticle)
    }

    r.Run(":8080")
}
```

- ```gin.Default()``` : Gin 프레임워크의 라우터 생성

- ```r.LoadHTMLGlob("view/*")``` : "./view/" 경로에 있는 html파일들을 요청 처리에 사용할 수 있도록 로드
    * ex) render 함수의 c.HTML()

- ```r.GET("/", handler.ShowIndexPage)``` : URL path가 "/" 인 경우 handler.ShowIndexPage 핸들러로 처리하도록 등록

- ```article := r.Group("/article")``` : URL path가 "/article" 으로 시작하는 요청에 대하여 그룹으로 묶어 article이라는 라우터로 관리

- ```article.GET("/view/:article_id", handler.GetArticle)```
    * Path : /article/view/:article_id
    * Path example : "/article/view/3", "/article/view/11" 
    * Path에 ":"를 지정하면 Gin의 라우터가 파라미터로 처리
    * handler.GetArticle 핸들러 함수에서 해당 파라미터는 정수타입으로 활용하고 있음
    * Gin에서 요청의 파라미터, 쿼리 데이터는 스트링 타입으로 처리하기 때문에 적절한 변환 필요 

- ```r.Run(":8080")``` : 루프백 주소(localhost)의 8080포트로 소켓을 열고 서버를 실행

#### Clean Architecture

<img src="https://user-images.githubusercontent.com/55284181/128448586-7bbe2aef-6c55-4bc4-89d3-b6abada1ac40.png" width="600" title="clean architecture">

<img src="https://user-images.githubusercontent.com/55284181/128449644-bcd91adc-4682-4280-90c2-eb6dcfdf9f5d.png" width="700" title="repository">

- v2
    + Entity(domain) : model
        ```go
        type Article struct {
            ID        int       `json:"id" gorm:"primary_key"`
            Title     string    `json:"title" gorm:"type:varchar(64);"`
            Content   string    `json:"content" gorm:"type:varchar(128)"`
            CreatedAt time.Time `json:"created_at"`
        }
        ```
    + Repository : model과 usecase 계층 연결. usecase 계층에게 데이터에 접근할 수 있는 인터페이스를 제공한다.
        ```go
        type ArticleRepo interface {
            GetAll() ([]model.Article, error)
            GetByID(id int) (*model.Article, error)
            Create(article *model.Article) (*model.Article, error)
            Delete(article *model.Article) error
        }
        ```
    + Usecase : 어플리케이션 작업에 대한 함수. handler 함수에서 사용한다.
        ```go
        type ManageArticleUsecase interface {
            GetAllArticles() ([]model.Article, error)
            GetArticleByID(id int) (*model.Article, error)
            CreateNewArticle(title, content string) (*model.Article, error)
            DeleteArticleByID(id int) error
        }
        ```
    + Handler
        ```go
        func (h *GinHandler) ShowIndexPage(c *gin.Context)
        func (h *GinHandler) ShowArticleCreationPage(c *gin.Context)
        func (h *GinHandler) ShowArticle(c *gin.Context)
        func (h *GinHandler) NewArticle(c *gin.Context)
        func (h *GinHandler) RemoveArticle(c *gin.Context)
        ```
    + 의존성 주입 : main에서 인터페이스로 추상화 되어있는 계층들을 실제 구현한 인스턴스로 연결한다.
        ```go
        ar := mysql.NewArticleRepo()
        mauc := manageArticle.NewManageArticleUsecase(ar)
        h := rest.NewGinHandler(mauc)
        ```

#### mysql connection
AutoMigrate를 통하여 mysql 데이터베이스에 Article 테이블 자동으로 생성한다.

> AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It WON’T delete unused columns to protect your data.

```go
var dbConn *gorm.DB

func Setup() {
	var err error
	conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(dbms, conn)
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&model.Article{},
	)
}
```

#### gorm
gorm을 통해 직접 SQL Query를 작성하지 않고 DB를 사용한다. 메소드로 쿼리 생성 후 객체 인스턴스에 자동으로 바인딩한다. ```db.Where("--").Find(&article)```

```go
func NewArticleRepo() *articleRepo {
    return &articleRepo{
        db: dbConn,
    }
}

func (ar *articleRepo) GetAll() (al []model.Article, err error) {
    return al, ar.db.Find(&al).Error
    // SELECT * FROM articles;
}

func (ar *articleRepo) GetByID(id int) (a *model.Article, err error) {
    a = new(model.Article)
    return a, ar.db.Where("id=?", id).First(a).Error
    // SELECT * FROM articles WHERE id = 'id' ORDER BY id LIMIT 1;
}

func (ar *articleRepo) Create(article *model.Article) (*model.Article, error) {
    return article, ar.db.Create(article).Error
}

func (ar *articleRepo) Delete(article *model.Article) error {
    return ar.db.Delete(article).Error
    // DELETE FROM articles WHERE id = 'id';
}
```

- v3
    + domain
        Article 테이블에 writer_id FK를 추가하고, User 테이블을 추가하였다.

        <img src="https://user-images.githubusercontent.com/55284181/128635256-f12b82e5-595b-408c-9939-9ff7719c875f.png" width="300" title="v3 schema">




---
### reference
- [Go 패키지 시리즈 1 템플릿 사용하기](http://chanlee.github.io/2016/04/21/golang-template-package/)
- [Gin을 사용하여 웹앱과 마이크로서비스 만들기](https://earntrust.tistory.com/entry/Gin%EC%9D%84-%EC%82%AC%EC%9A%A9%ED%95%98%EC%97%AC-%EC%9B%B9%EC%95%B1%EA%B3%BC-%EB%A7%88%EC%9D%B4%ED%81%AC%EB%A1%9C%EC%84%9C%EB%B9%84%EC%8A%A4-%EB%A7%8C%EB%93%A4%EA%B8%B0)
