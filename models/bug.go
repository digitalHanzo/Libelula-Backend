package models

//Bug is a model
type Bug struct {
	//to json, not null
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	FamilyID    int     `json:"family_id,omitempty"`
	Family      *Family `json:"family,omitempty"`
}
