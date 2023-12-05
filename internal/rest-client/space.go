package client

type Space struct {
	DataReduction float64 `json:"data_reduction"`
	Snapshots     int64   `json:"snapshots"`
	TotalPhysical int64   `json:"total_physical"`
	Unique        int64   `json:"unique"`
	Virtual       int64   `json:"virtual"`
}
