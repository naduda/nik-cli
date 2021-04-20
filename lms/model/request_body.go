package model

func NewRequestBody(mode, date string, id int) RequestBody {
	return RequestBody{
		Mode: mode,
		Data: RequestBodyData{
			GenerationId: id,
			Date:         date,
			Body:         make(map[int]float64),
		},
	}
}

type RequestBody struct {
	Mode string          `json:"mode"`
	Data RequestBodyData `json:"data"`
}

type RequestBodyData struct {
	Date          string          `json:"date"`
	GenerationId  int             `json:"generation_id"`
	SourceVersion int             `json:"sourceVersion"`
	Body          map[int]float64 `json:"body"`
}

type GetResponse struct {
	Success bool           `json:"success"`
	Prs     GetResponsePrs `json:"prs"`
}

type GetResponsePrs struct {
	Version           int                      `json:"version"`
	Date              string                   `json:"date"`
	Body              map[int]float64          `json:"body"`
	AvailableVersions []GetResponsePrsVersions `json:"availableVersions"`
}

type GetResponsePrsVersions struct {
	Version           int    `json:"version"`
	CreatedAt         string `json:"created_at"`
	CreatedByUsername string `json:"created_by__username"`
}
