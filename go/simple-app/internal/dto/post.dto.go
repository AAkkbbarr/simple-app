package dto

type UpdatePostDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title,string"`
	Body  string `json:"body"`
}

type CreatePostDTO struct {
	Title string `json:"title,string"`
	Body  string `json:"body"`
} 