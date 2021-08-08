# Database


## 1. 용어 정리

<img src="https://user-images.githubusercontent.com/55284181/128584205-901982fe-bb0b-4598-ad29-ca25501e4a23.png" width="650" title="database">

- database server(데이터베이스 서버) : 스키마들을 저장하는 곳
- database(데이터베이스) : 서로 연관된 표들을 grouping할 때 사용하는 일종의 폴더 (= schema)
- table(표) : 데이터가 실질적으로 저장되는 저장소

<img src="https://user-images.githubusercontent.com/55284181/128589772-0cd921e2-782d-4259-8280-6343c6fff37e.png" width="600" title="mysql table">

- column(열) : 속성 정보를 의미한다. 각각의 열은 유일한 이름을 가지고 있으며, 자신만의 타입을 가진다. (= field, attribute)
- row(행) : 관계된 데이터의 묶음을 의미한다. (= record, tuple)
- degree : 열의 수를 의미하며, 0이 될 수 없다.
- cardinality : 행의 수를 의미하며, 0이 될 수 있다.

+ key
    - PK(Primary Key) : 테이블에서 행의 식별자로 이용되는 열
    - FK(Foreign Key) : 한 테이블의 키 중에서 다른 테이블의 PK

+ relationship(관계) : 테이블 간의 관계는 관계를 맺는 테이블의 수에 따라 다음과 같이 나눌 수 있다.
    <img src="https://user-images.githubusercontent.com/55284181/128589688-a3db9878-a91f-4a91-9184-410d149902f9.png" width="600" title="mysql relationship">
    1. 일대일(one-to-one) 관계
    2. 일대다(one-to-many) 관계
    3. 다대다(many-to-many) 관계


## 2. MySQL

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

#### 관계형 데이터베이스
- Model : 어떤 목적을 가지고 진짜를 모방한 것 (추상적 의미)

#### 관계형 데이터 모델링
업무파악 -> 개념적 데이터 모델링 -> 논리적 데이터 모델링 -> 물리적 데이터 모델링



## 3. gorm

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



---
### reference
- [관계형 데이터베이스](http://tcpschool.com/mysql/mysql_intro_relationalDB)
- [ORM이란](https://gmlwjd9405.github.io/2019/02/01/orm.html)
- [Eager Loading & Options in ORM](https://velog.io/@minho/Eager-Loading-Options-in-ORM)
- [gorm Query](https://gorm.io/docs/query.html)
- [DATABASE2 MySQL - 생활코딩 유튜브 강의](https://www.youtube.com/watch?v=-w1vJgslUG0&list=PLuHgQVnccGMCgrP_9HL3dAcvdt8qOZxjW&index=21)
- [관계형 데이터 모델링 - 생활코딩 유튜브 강의](https://www.youtube.com/playlist?list=PLuHgQVnccGMDF6rHsY9qMuJMd295Yk4sa)
