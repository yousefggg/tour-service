package dto
import("github.com/google/uuid")


type CreateTourRequest struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	Country     string `json:"country"`
	Season      string `json:"season"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Includes    string `json:"includes"`
}

type UpdateTourRequest struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	Country     string `json:"country"`
	Season      string `json:"season"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Includes    string `json:"includes"`
}

type TourCardResponse struct {
	ID       uuid.UUID `json:"id"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Season   string `json:"season"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

type TourDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Country     string `json:"country"`
	Season      string `json:"season"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Includes    string `json:"includes"`
}