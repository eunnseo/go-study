## 2. 용어 정리

#### Go template
- 템플릿(template) : 대부분의 서버사이드 언어들은 정적인 페이지에 스크립팅을 지원하기 위한 언어들을 제공한다. JSP나 PHP 스크립팅 같은 것이 그런 예이다. Go에서는 이와 유사한 스크립팅 언어로 template 패키지를 제공한다.

#### debug mode <-> release mode
- release mode : 프로그램을 배포하기 위해 컴파일 하는 모드
- debug mode : 컴파일 시 들어가는 디버깅에 필요한 자질구리한 정보를 뺀 알짜 프로그램만 쏙 뽑아냄

#### MySQL Database
<img src="https://user-images.githubusercontent.com/55284181/128584205-901982fe-bb0b-4598-ad29-ca25501e4a23.png" width="700" title="database">

- database server(데이터베이스 서버) : 스키마들을 저장하는 곳
- database(데이터베이스) : 서로 연관된 표들을 grouping할 때 사용하는 일종의 폴더 (= schema)
- table(표) : 데이터가 실질적으로 저장되는 저장소

<img src="https://user-images.githubusercontent.com/55284181/128589772-0cd921e2-782d-4259-8280-6343c6fff37e.png" width="600" title="mysql table">

- column(열) : 속성 정보를 의미한다. 각각의 열은 유일한 이름을 가지고 있으며, 자신만의 타입을 가진다. (= field, attribute)
- row(행) : 관계된 데이터의 묶음을 의미한다. (= record, tuple)
- degree : 열의 수를 의미하며, 0이 될 수 없다.
- cardinality : 행의 수를 의미하며, 0이 될 수 있다.
- key
    - PK(Primary Key) : 테이블에서 행의 식별자로 이용되는 열
    - FK(Foreign Key) : 한 테이블의 키 중에서 다른 테이블의 PK
- relationship(관계) : 테이블 간의 관계는 관계를 맺는 테이블의 수에 따라 다음과 같이 나눌 수 있다.
    <img src="https://user-images.githubusercontent.com/55284181/128589688-a3db9878-a91f-4a91-9184-410d149902f9.png" width="600" title="mysql relationship">
    1. 일대일(one-to-one) 관계
    2. 일대다(one-to-many) 관계
    3. 다대다(many-to-many) 관계

#### Relational Database
- Model : 어떤 목적을 가지고 진짜를 모방한 것 (추상적 의미)


## 3. Gin

Gin은 Web Application과 Microservices를 만드는 데 사용되는 고성능 Micro-Framework이다. 실습 코드는 eunnseo/web-tuto-with-gin 에 위치한다.

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


## 4. MySQL

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

#### CRUD Interface

- **Create Record**

    - Create
        ```go
        user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
        result := db.Create(&user) // pass pointer of data to Create

        user.ID             // returns inserted data's primary key
        result.Error        // returns error
        result.RowsAffected // returns inserted records count
        ```

- **Create Record With Selected Fields**

    - Create : record를 만들고 지정된 필드에 값을 할당
        ```go
        db.Select("Name", "Age", "CreatedAt").Create(&user)
        // INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
        ```

    - Omit : record를 만들고 생략하도록 전달된 필드의 값은 무시
        ```go
        db.Omit("Name", "Age", "CreatedAt").Create(&user)
        // INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
        ```

- **Delete a Record**

    - Delete
        ```go
        // Email's ID is `10`
        db.Delete(&email)
        // DELETE from emails where id = 10;
        ```

#### Query

- **Retrieving a single object**

    - First
        ```go
        // Get the first record ordered by primary key
        db.First(&user)
        // SELECT * FROM users ORDER BY id LIMIT 1;
        ```

    - Take
        ```go
        // Get one record, no specified order
        db.Take(&user)
        // SELECT * FROM users LIMIT 1;
        ```

    - Last
        ```go
        // Get last record, ordered by primary key desc
        db.Last(&user)
        // SELECT * FROM users ORDER BY id DESC LIMIT 1;
        ```

    - check result
        ```go
        result := db.First(&user)
        result.RowsAffected // returns count of records found
        result.Error        // returns error or nil

        // check error ErrRecordNotFound
        errors.Is(result.Error, gorm.ErrRecordNotFound)
        ```

- **Retrieving objects with primary key**

    - First
        ```go
        db.First(&user, 10)
        // SELECT * FROM users WHERE id = 10;

        db.First(&user, "10")
        // SELECT * FROM users WHERE id = 10;
        ```

    - Find
        ```go
        db.Find(&users, []int{1,2,3})
        // SELECT * FROM users WHERE id IN (1,2,3);
        ```

- **Retrieving all objects**

    - Find
        ```go
        // Get all records
        result := db.Find(&users)
        // SELECT * FROM users;

        result.RowsAffected // returns found records count, equals `len(users)`
        result.Error        // returns error
        ```

- **Conditions**

    - String Conditions
        ```go
        // Get first matched record
        db.Where("name = ?", "jinzhu").First(&user)
        // SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

        // Get all matched records
        db.Where("name <> ?", "jinzhu").Find(&users)
        // SELECT * FROM users WHERE name <> 'jinzhu';

        // IN
        db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
        // SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

        // LIKE
        db.Where("name LIKE ?", "%jin%").Find(&users)
        // SELECT * FROM users WHERE name LIKE '%jin%';

        // AND
        db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
        // SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

        // Time
        db.Where("updated_at > ?", lastWeek).Find(&users)
        // SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

        // BETWEEN
        db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
        // SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
        ```

#### Eager Loading (Preloading)

- Preload
    ```go
    type User struct {
        gorm.Model
        Username string
        Orders   []Order
    }

    type Order struct {
        gorm.Model
        UserID uint
        Price  float64
    }

    // Preload Orders when find users
    db.Preload("Orders").Find(&users)
    // SELECT * FROM users;
    // SELECT * FROM orders WHERE user_id IN (1,2,3,4);

    db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
    // SELECT * FROM users;
    // SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
    // SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
    // SELECT * FROM roles WHERE id IN (4,5,6); // belongs to
    ```

- Preload All
    ```clause.Associations```는 creating/updating 시 ```Preload```와 함께 작동할 수 있으며, 모든 연결을 ```Preload```하는 데 사용할 수 있다.
    ```go
    type User struct {
        gorm.Model
        Name       string
        CompanyID  uint
        Company    Company
        Role       Role
        Orders     []Order
    }

    db.Preload(clause.Associations).Find(&users)
    ```

## 5. 관계형 데이터베이스

#### 테이블 분리
테이블을 분리하여 데이터를 관리하면 추후 데이터의 수정 및 유지보수가 편리하다.

- topic table
    ```sql
    DESC topic;
    ```
    | Field       | Type        | Null | Key | Default | Extra          |
    |-------------|-------------|------|-----|---------|----------------|
    | id          | int         | NO   | PRI | NULL    | auto_increment |
    | title       | varchar(30) | NO   |     | NULL    |                |
    | description | text        | YES  |     | NULL    |                |
    | created     | datetime    | NO   |     | NULL    |                |
    | author_id   | int         | YES  |     | NULL    |                |
    ```sql
    SELECT * FROM topic;
    ```
    | id | title      | description       | created             | author_id |
    |---:|:-----------|:------------------|:--------------------|----------:|
    |  1 | MySQL      | MySQL is...       | 2018-01-01 12:10:11 |         1 |
    |  2 | Oracle     | Oracle is ...     | 2018-01-03 13:01:10 |         1 |
    |  3 | SQL Server | SQL Server is ... | 2018-01-20 11:01:10 |         2 |
    |  4 | PostgreSQL | PostgreSQL is ... | 2018-01-23 01:03:03 |         3 |
    |  5 | MongoDB    | MongoDB is ...    | 2018-01-30 12:31:03 |         1 |

- author table
    ```sql
    DESC author;
    ```
    | Field   | Type         | Null | Key | Default | Extra          |
    |---------|--------------|------|-----|---------|----------------|
    | id      | int          | NO   | PRI | NULL    | auto_increment |
    | name    | varchar(20)  | NO   |     | NULL    |                |
    | profile | varchar(200) | YES  |     | NULL    |                |
    ```sql
    SELECT * FROM author;
    ```
    | id | name   | profile                   |
    |---:|:-------|:--------------------------|
    |  1 | egoing | developer                 |
    |  2 | duru   | database administrator    |
    |  3 | taeho  | data scientist, developer |

#### Join
```sql
SELECT * FROM topic LEFT JOIN author ON topic.author_id = author.id;
```
| id | title      | description       | created             | author_id | id   | name   | profile                   |
|---:|------------|-------------------|---------------------|----------:|-----:|--------|---------------------------|
|  1 | MySQL      | MySQL is...       | 2018-01-01 12:10:11 |         1 |    1 | egoing | developer                 |
|  2 | Oracle     | Oracle is ...     | 2018-01-03 13:01:10 |         1 |    1 | egoing | developer                 |
|  3 | SQL Server | SQL Server is ... | 2018-01-20 11:01:10 |         2 |    2 | duru   | database administrator    |
|  4 | PostgreSQL | PostgreSQL is ... | 2018-01-23 01:03:03 |         3 |    3 | taeho  | data scientist, developer |
|  5 | MongoDB    | MongoDB is ...    | 2018-01-30 12:31:03 |         1 |    1 | egoing | developer                 |

```sql
SELECT topic.id AS topic_id,title,description,created,name,profile FROM topic LEFT JOIN author ON topic.author_id = author.id;
```
| topic_id | title      | description       | created             | name   | profile                   |
|---------:|------------|-------------------|---------------------|--------|---------------------------|
|        1 | MySQL      | MySQL is...       | 2018-01-01 12:10:11 | egoing | developer                 |
|        2 | Oracle     | Oracle is ...     | 2018-01-03 13:01:10 | egoing | developer                 |
|        3 | SQL Server | SQL Server is ... | 2018-01-20 11:01:10 | duru   | database administrator    |
|        4 | PostgreSQL | PostgreSQL is ... | 2018-01-23 01:03:03 | taeho  | data scientist, developer |
|        5 | MongoDB    | MongoDB is ...    | 2018-01-30 12:31:03 | egoing | developer                 |

#### 관계형 데이터 모델링
업무파악 -> 개념적 데이터 모델링 -> 논리적 데이터 모델링 -> 물리적 데이터 모델링



---
### reference
- [debug <=> release 모드의 차이점과 배포](https://killsia.tistory.com/entry/debug-release-%EB%AA%A8%EB%93%9C%EC%9D%98-%EC%B0%A8%EC%9D%B4%EC%A0%90%EA%B3%BC-%EB%B0%B0%ED%8F%AC)
- [관계형 데이터베이스](http://tcpschool.com/mysql/mysql_intro_relationalDB)

- [gorm Query](https://gorm.io/docs/query.html)
- [DATABASE2 MySQL - 생활코딩 유튜브 강의](https://www.youtube.com/watch?v=-w1vJgslUG0&list=PLuHgQVnccGMCgrP_9HL3dAcvdt8qOZxjW&index=21)
- [관계형 데이터 모델링 - 생활코딩 유튜브 강의](https://www.youtube.com/playlist?list=PLuHgQVnccGMDF6rHsY9qMuJMd295Yk4sa)
- [Eager Loading & Options in ORM](https://velog.io/@minho/Eager-Loading-Options-in-ORM)

- [Gin을 사용하여 웹앱과 마이크로서비스 만들기](https://earntrust.tistory.com/entry/Gin%EC%9D%84-%EC%82%AC%EC%9A%A9%ED%95%98%EC%97%AC-%EC%9B%B9%EC%95%B1%EA%B3%BC-%EB%A7%88%EC%9D%B4%ED%81%AC%EB%A1%9C%EC%84%9C%EB%B9%84%EC%8A%A4-%EB%A7%8C%EB%93%A4%EA%B8%B0)
