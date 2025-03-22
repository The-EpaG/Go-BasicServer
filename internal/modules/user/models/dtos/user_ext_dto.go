package dtos

type UserExtDTO struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type UserExtPageDTO struct {
	Page       int          `json:"page"`
	PerPage    int          `json:"per_page"`
	Total      int          `json:"total"`
	TotalPages int          `json:"total_pages"`
	Data       []UserExtDTO `json:"data"`
}
