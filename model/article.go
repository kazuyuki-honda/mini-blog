package model

import (
    "time"
	"math"
	// "github.com/jinzhu/gorm"
    "gopkg.in/go-playground/validator.v9"
  )

// PageRequest
type PageRequest struct {
	Limit                 int     `json:"limit"`
	Page                  int     `json:"page"`
}

// Article ...
type Article struct {
	ID      int       `db:"id" form:"id" json:"id"`
    Title   string    `db:"title" form:"title" validate:"required,max=50" json:"title"`
    Body    string    `db:"body" form:"body" validate:"required" json:"body"`
    Created time.Time `db:"created" json:"created"`
    Updated time.Time `db:"updated" json:"updated"`
}

// ArticleResponse 記事一覧とページを返す
type ArticleResponse struct {
	Article []*Article `json:"articles"`
	Total       int           `json:"total"`
	Pages       int           `json:"pages"`
	Page        int           `json:"page"`
	Limit       int           `json:"limit"`
}



// ValidationErrors ...
func (a *Article) ValidationErrors(err error) []string {
	// メッセージを格納するスライスを宣言します。
	var errMessages []string
  
	// 複数のエラーが発生する場合があるのでループ処理を行います。
	for _, err := range err.(validator.ValidationErrors) {
	  // メッセージを格納する変数を宣言します。
	  var message string
  
	  // エラーになったフィールドを特定します。
	  switch err.Field() {
	  case "Title":
		// エラーになったバリデーションルールを特定します。
		switch err.Tag() {
		case "required":
		  message = "タイトルは必須です。"
		case "max":
		  message = "タイトルは最大50文字です。"
		}
	  case "Body":
		message = "本文は必須です。"
	  }
  
	  // メッセージをスライスに追加します。
	  if message != "" {
		errMessages = append(errMessages, message)
	  }
	}
  
	return errMessages
  }


//記事全件取得
// func GetAllArticle() (*[]Article, error) {

	
// 	res :=make([]Article,0)

// 	if err := db.Model(&Article{}).Find(&res).Error; err !=nil{
// 		return nil, err
// 	}

// 	return &res, nil
// }


func GetAllArticle(req *PageRequest) (*ArticleResponse ,error) {
	query := db.Model(&Article{})
	
	// 件数カウント
	var total int
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// 検索実行
	var res ArticleResponse
		// limit作成 / offset作成 / order作成
		searchQuery := query.Limit(req.Limit)
		searchQuery = searchQuery.Offset((req.Page - 1) * req.Limit)
		searchQuery = searchQuery.Order("articles.created desc") // order作成(降順固定
	
	article :=make([]*Article,0)
	if err := searchQuery.Find(&article).Error; err !=nil{
		return nil, err
	}

	res.Article = article
	res.Total = total
	res.Pages = int(math.Ceil(float64(total) / float64(req.Limit)))
	res.Page = req.Page
	res.Limit = req.Limit

	return &res, nil
}

func GetIDArticle(id int) (*Article, error) {
	var res Article
	if err := db.Model(&Article{}).Where("id=?", id).Find(&res).Error; err !=nil{
		return nil, err
	}

	return &res, nil
}


func CreateArticle(req *Article)error{
	req.Created = time.Now()
	req.Updated = time.Now()

	if err:= db.Create(req).Error; err !=nil{
		return err
	}

	return nil
}

func UpdateArticle(req *Article)error{
	req.Updated = time.Now()
	// defer db.Close()

	if err:= db.Model(&Article{}).Where("id=?", req.ID).Updates(req).Error; err !=nil{
		return err
	}

	return nil
}