package client

type Space struct {
	DataReduction        float64 `json:"data_reduction"`
	Snapshots            int64   `json:"snapshots"`
	TotalPhysical        int64   `json:"total_physical"`
	Unique               int64   `json:"unique"`
	TotalProvisioned     int64   `json:"total_provisioned"`
	AvailableProvisioned int64   `json:"available_provisioned"`
	AvailableRatio       float64 `json:"available_ratio"`
	Destroyed            int64   `json:"destroyed"`
	DestroyedVirtual     int64   `json:"destroyed_virtual"`
	Virtual              int64   `json:"virtual"`
}
