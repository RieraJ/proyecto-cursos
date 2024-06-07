package dto

type Course struct {
	ID          uint    `json:"id"`          // Course ID
	Price       float64 `json:"price"`       // Course price
	Active      bool    `json:"active"`      // Course status
	Name        string  `json:"name"`        // Course name
	Description string  `json:"description"` // Course description
}

type SearchResponse struct {
	Courses []Course `json:"courses"`
}

type InscriptionRequest struct {
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}
