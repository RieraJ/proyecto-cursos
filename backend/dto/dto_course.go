package dto

type Course struct {
	ID           uint    `json:"id"`           // Course ID
	Price        float64 `json:"price"`        // Course price
	Active       bool    `json:"active"`       // Course status
	Name         string  `json:"name"`         // Course name
	Description  string  `json:"description"`  // Course description
	Instructor   string  `json:"instructor"`   // Course instructor
	Length       string  `json:"length"`       // Course length
	Requirements string  `json:"requirements"` // Course requirements
	Image        string  `json:"image"`        // Course image
}

type SearchRequest struct {
	Name string `json:"name"`
}

type InscriptionRequest struct {
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}
