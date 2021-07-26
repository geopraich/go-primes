package model

type PrimesResponse struct {
	Initial int32   `json:"initial"`
	Primes  []int32 `json:"primes"`
}
