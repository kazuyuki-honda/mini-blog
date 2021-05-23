package main

import (
	"log"
	"os"
	// "net/http"
	"fmt"
	"time"
	"mini-blog/handler"
	// "mini-blog/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// _ "github.com/go-sql-driver/mysql" // Using MySQL driver
	// "github.com/jmoiron/sqlx"
	"gopkg.in/go-playground/validator.v9"
)



const(
	AdminPort ="8888"
)

func init(){
	// デフォルトのタイムゾーン設定
	location := "Asia/Tokyo"

	loc, err := time.LoadLocation(location)
	if err != nil {
		// 失敗の場合は、9時間をアナログにずらす
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	// ログの設定
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	// ログは標準出力に出す
	log.SetOutput(os.Stdout)
}

func newRouter() *echo.Echo {
	e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
	e.Use(middleware.Gzip())



    e.Static("/assets", "public/assets")
	

	e.File("/", "public/index.html")

	// e.File("/article/:articleID", "public/show.html")
	
    e.File("/signup", "public/signup.html")
    e.POST("/signup", handler.Signup)
    e.File("/login", "public/login.html")
    e.POST("/login", handler.Login)

	e.File("/admin","public/admin.html")
	e.File("/admin/articles/new","public/edit.html")
	e.File("/admin/articles/edit/:articleID","public/edit.html")
	e.File("/admin/articles/:articleID","public/show.html")



	api := e.Group("/api")
	api.POST("/articles/list", handler.ArticleList)
	api.GET("/articles/:articleID",handler.ArticleShow)
	api.POST("/articles/new", handler.ArticleEdit)
	api.POST("/articles/:articleID", handler.ArticleEdit)


    
    // api.Use(middleware.JWTWithConfig(handler.Config))
    // api.GET("/todos", handler.GetTodos)
    // api.POST("/todos", handler.AddTodo)
    // api.DELETE("/todos/:id", handler.DeleteTodo)
    // api.PUT("/todos/:id/completed", handler.UpdateTodo)


	 // TOP ページに記事の一覧を表示します。
	//  e.GET("/", handler.ArticleIndex)

	 // 記事に関するページは "/articles" で開始するようにします。
   	// 記事一覧画面には "/" と "/articles" の両方でアクセスできるようにします。
  	// パスパラメータの ":id" も ":articleID" と明確にしています。

  	// e.GET("/articles", handler.ArticleIndex)         // 一覧画面
  	// e.GET("/articles/new", handler.ArticleNew)       // 新規作成画面
  	// e.GET("/articles/:articleID", handler.ArticleShow)      // 詳細画面
  	// e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

  // HTML ではなく JSON を返却する処理は "/api" で開始するようにします。
  // 記事に関する処理なので "/articles" を続けます。

//   e.GET("/api/articles", handler.ArticleList)          // 一覧
//   e.POST("/api/articles", handler.ArticleCreate)       // 作成
//   e.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 削除
//   e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)  // 更新

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}


  // CustomValidator ...
  type CustomValidator struct {
	validator *validator.Validate
  }
 
  // Validate ...
  func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
  }
  



func main() {

	// echoサーバの起動
	router := newRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf("0.0.0.0:%s", AdminPort)))
}





