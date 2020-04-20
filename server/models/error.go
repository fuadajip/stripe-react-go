package models

// ExperimentError is a general Experiment error response struct
type ExperimentError struct {
	Code      int
	ErrorCode string
	Message   string
	Status    string
}
