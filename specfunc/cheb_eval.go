package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

type cheb_series struct {
	c        []float64
	order    int
	a, b     float64
	order_sp int
}

func Cheb_eval_e(cs cheb_series, x float64, result *Result) error {
	var temp float64
	d := 0.0
	dd := 0.0

	y := (2.0*x - cs.a - cs.b) / (cs.b - cs.a)
	y2 := 2.0 * y

	e := 0.0

	for j := cs.order; j >= 1; j-- {
		temp = d
		d = y2*d - dd + cs.c[j]
		e += math.Abs(y2*temp) + math.Abs(dd) + math.Abs(cs.c[j])
		dd = temp
	}

	temp = d
	d = y*d - dd + 0.5*cs.c[0]
	e += math.Abs(y*temp) + math.Abs(dd) + 0.5*math.Abs(cs.c[0])

	result.val = d
	result.err = gsl.DBL_EPSILON*e + math.Abs(cs.c[cs.order])

	return err.SUCCESS
}
