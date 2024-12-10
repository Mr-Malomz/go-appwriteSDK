package data

type Project struct {
	Id          string `json:"$id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProjectResponse struct {
	Id           string `json:"$id,omitempty"`
	CollectionId string `json:"$collectionId,omitempty"`
}

type JsonAPIBody struct {
	Data       *ProjectRequest `json:"data,omitempty"`
}
