package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

type Result struct {
	val float64
	err float64
}

type Result_e10 struct {
	val float64
	err float64
	e10 int
}

func result_smash_e(re *Result_e10, r *Result) error {
	if re.e10 == 0 {
		r.val = re.val
		r.err = re.err
		return err.SUCCESS
	}

	av := math.Abs(re.val)
	ae := math.Abs(re.err)

	if gsl.SQRT_DBL_MIN < av && av < gsl.SQRT_DBL_MAX && gsl.SQRT_DBL_MIN < ae && ae < gsl.SQRT_DBL_MAX && 0.49*gsl.LOG_DBL_MIN < float64(re.e10) && float64(re.e10) < 0.49*gsl.LOG_DBL_MAX {
		scale := math.Exp(float64(re.e10) * gsl.M_LN10)
		r.val = re.val * scale
		r.err = re.err * scale
		return err.SUCCESS
	}

	return Exp_mult_err_e(float64(re.e10)*gsl.M_LN10, 0.0, re.val, re.err, r)
}
