package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-app/internal/entities"
	// "simple-app/internal/entities"
)

type postController struct {
}

func NewPostController() *postController {
	return &postController{}
}

func (p *postController) GetAll(res http.ResponseWriter, req *http.Request) {

	//defer res.Header().Set("content-type", "application/json")

	clinet := &http.Client{}
	clientReq, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		res.Write([]byte("Prearing posts request"))
		return
	}

	resData, err := clinet.Do(clientReq)
	fmt.Println(resData.StatusCode)
	postsData, err := ioutil.ReadAll(resData.Body)
	fmt.Println("sbdviuSBFGIUWS", string(postsData))
	posts := &entities.Posts{}
	parseErr := json.Unmarshal(postsData, &posts)
	if parseErr != nil {
		fmt.Println("err,parseErr")
	} else {
		fmt.Println("posts---->", posts)
		res.Write(postsData)
		return

	}

	if err != nil {
		res.Write([]byte("Error getting posts from server" + err.Error()))
		return
	}

	fmt.Println(resData.Body)

	if resData.Body != nil {
		resData.Body.Close()
		res.Write([]byte("Error: content is null" + err.Error()))
		return
	}

	if resData.StatusCode != 200 {
		res.Write([]byte("Error: status code is't 200"))
		return
	}

	// postsData, err := ioutil.ReadAll(resData.Body)

	if err != nil {
		res.Write([]byte("Error: status code is't 200"))
		return
	}

	// posts := &entities.Posts{}
	// parseErr := json.Unmarshal(postsData, &posts)

	// if parseErr != nil {
	// 	res.Write([]byte("Error: status code is't 200"))
	// 	return
	// }

	res.Header().Set("content-type", "application/json")
	res.Write([]byte(postsData))
}

func (p *postController) GetOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("get one"))
}

func (p *postController) CreatePost(res http.ResponseWriter, req *http.Request) {
	// post := entities.Post{
	// 	Title: req.FormValue("title"),
	// 	Body:  req.FormValue("body"),
	// }

	res.Header().Set("content-type", "application/json")
	res.Write([]byte("create post"))
}

func (p *postController) DeletePost(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("content-type", "application/json")
	res.Write([]byte("delete post"))
}
