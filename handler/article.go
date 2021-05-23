package handler

import(
	"net/http"
	"mini-blog/model"
	"github.com/labstack/echo/v4"
	"strconv"
	_"fmt"
	_"time"
)

// ArticleListは記事一覧を返す
func ArticleList(c echo.Context) error {

	var req model.PageRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "不正なリクエストです。",
			"error":   err.Error(),
		})
	}

	res, err :=model.GetAllArticle(&req)
	if err !=nil{
		return err
	}

	return c.JSON(http.StatusOK, res)
}

//特定のID記事を取得
func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))
	res, err :=model.GetIDArticle(id)
	if err !=nil{
		return err
	}
	return c.JSON(http.StatusOK, res)
}

//記事を作成する
func ArticleEdit(c echo.Context) error {
	

	var req model.Article
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "不正なリクエストです。",
			"error":   err.Error(),
		})
	}
	//バリデーション
	if err :=c.Validate(&req); err !=nil{
		errMessages :=req.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": errMessages,
			"error":   err.Error(),	
		})
	}

		//articleIDがあればUpdate
		articleID:=c.Param("articleID")
		if len(articleID) >0{	
			if err:=model.UpdateArticle(&req); err !=nil{
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "登録に失敗しました",
					"error":   err.Error(),
				})
			}			
		}else{	
			if err:=model.CreateArticle(&req); err !=nil{
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "登録に失敗しました",
					"error":   err.Error(),
				})
			}
		}

	return c.JSON(http.StatusOK, "ok")
}