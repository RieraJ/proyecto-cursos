package dto

type Course struct {
	ID           uint       `json:"id"`           // Course ID
	Price        float64    `json:"price"`        // Course price
	Active       bool       `json:"active"`       // Course status
	Name         string     `json:"name"`         // Course name
	Description  string     `json:"description"`  // Course description
	Instructor   string     `json:"instructor"`   // Course instructor
	Length       string     `json:"length"`       // Course length
	Requirements string     `json:"requirements"` // Course requirements
	Categories   []Category `json:"categories"`   // Course category
}

type Category struct {
	ID   uint   `json:"id"`   // Category ID
	Name string `json:"name"` // Category name
}

type InscriptionRequest struct {
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}
