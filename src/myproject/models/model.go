package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	Id      int
	Name    string `orm:"unique"`
	Pwd     string
	Article []*Article `orm:"rel(m2m)"`
}

// Article ...
type Article struct {
	Id           int       `orm:"pk;auto"`
	ArtiName     string    `orm:"size(20)"`
	Atime        time.Time `orm:"auto_now"`
	Acontent     string    `orm:"size(500)"`
	Acount       int       `orm:"default(0);null"`
	Aimg         string
	ArticleTypes *ArticleType `orm:"rel(fk)"` //外键
	Users        []*User      `orm:"reverse(many)"`
}

// ArticleType ...
type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"` //一对多
}

/*
mysql> desc article;
+------------------+--------------+------+-----+---------+----------------+
| Field            | Type         | Null | Key | Default | Extra          |
+------------------+--------------+------+-----+---------+----------------+
| id               | int(11)      | NO   | PRI | NULL    | auto_increment |
| arti_name        | varchar(20)  | NO   |     |         |                |
| atime            | datetime     | NO   |     | NULL    |                |
| acontent         | varchar(500) | NO   |     |         |                |
| acount           | int(11)      | YES  |     | 0       |                |
| aimg             | varchar(255) | NO   |     |         |                |
| article_types_id | int(11)      | NO   |     | NULL    |                |
+------------------+--------------+------+-----+---------+----------------+
7 rows in set (0.04 sec)

mysql> desc article_type;
+-----------+-------------+------+-----+---------+----------------+
| Field     | Type        | Null | Key | Default | Extra          |
+-----------+-------------+------+-----+---------+----------------+
| id        | int(11)     | NO   | PRI | NULL    | auto_increment |
| type_name | varchar(20) | NO   |     |         |                |
+-----------+-------------+------+-----+---------+----------------+
2 rows in set (0.00 sec)

mysql> desc user_articles;
+------------+------------+------+-----+---------+----------------+
| Field      | Type       | Null | Key | Default | Extra          |
+------------+------------+------+-----+---------+----------------+
| id         | bigint(20) | NO   | PRI | NULL    | auto_increment |
| user_id    | int(11)    | NO   |     | NULL    |                |
| article_id | int(11)    | NO   |     | NULL    |                |
+------------+------------+------+-----+---------+----------------+
3 rows in set (0.00 sec)

mysql> desc user;
+-------+--------------+------+-----+---------+----------------+
| Field | Type         | Null | Key | Default | Extra          |
+-------+--------------+------+-----+---------+----------------+
| id    | int(11)      | NO   | PRI | NULL    | auto_increment |
| name  | varchar(255) | NO   | UNI |         |                |
| pwd   | varchar(255) | NO   |     |         |                |
+-------+--------------+------+-----+---------+----------------+
3 rows in set (0.00 sec)
*/
func init() {
	orm.RegisterDataBase("default", "mysql", "root:tester@tcp(127.0.0.1:3306)/test?charset=utf8")

	orm.RegisterModel(new(User), new(Article), new(ArticleType))

	orm.RunSyncdb("default", false, true)
	// true: 每次运行程序时, 重新创建表
	// orm.RunSyncdb("default", true, true)
}
