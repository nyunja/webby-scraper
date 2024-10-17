package store

// Struct to hold car page details
type CarPageDetails struct {
    Title       string `json:"title"`
    Details     []Detail `json:"details"`
}

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
