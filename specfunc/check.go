package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

func CheckUnderflow(res *Result) error {
	if math.Abs(res.val) < gsl.DBL_MIN {
		return err.EUNDRFLW
	}
	return err.SUCCESS
}
