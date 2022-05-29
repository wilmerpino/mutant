package model

type DnaChain struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	DNA      string `json:"dna"`
	IsMutant bool   `json:"is_mutant"`
}

type Stats struct {
	CountMutant int32   `json:"count_mutant_dna"`
	CountHuman  int32   `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}

type InputDNA struct {
	DNA []string `json:"dna"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type DnaInfo struct {
	DNA      []string
	IsMutant bool
}

func (DnaChain) TableName() string {
	return "dna_chains"
}
