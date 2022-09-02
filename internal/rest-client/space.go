package client


type Space struct {
    DataReduction  float64 `json:"data_reduction"`
    Snapshots      float64 `json:"snapshots"`
    TotalPhysical  float64 `json:"total_physical"`
    Unique         float64 `json:"unique"`
    Virtual        float64 `json:"virtual"`
}
