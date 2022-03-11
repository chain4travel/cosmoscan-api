package dmodels

type ValidatorValue struct {
	Validator string `db:"validator" json:"validator"`
	Value     string `db:"value" json:"value"`
}
