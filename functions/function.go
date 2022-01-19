package functions

type Function interface {
	F(x float64) float64
	Expression() string
	GetParamters() []float64
	SetParameters([]float64)
}
