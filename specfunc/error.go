package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

func OverflowError(result ...*Result) error {
	for _, r := range result {
		r.val = math.Inf(1)
		r.err = math.Inf(1)
	}
	return err.EOVRFLW
}

func UnderflowError(result ...*Result) error {
	for _, r := range result {
		r.val = 0.0
		r.err = gsl.DBL_MIN
	}
	return err.EUNDRFLW
}

func InternalOverflowError(result *Result) error {
	result.val = math.Inf(1)
	result.err = math.Inf(1)
	return err.EOVRFLW
}

func InternalUnderflowError(result *Result) error {
	result.val = 0.0
	result.err = gsl.DBL_MIN
	return err.EUNDRFLW
}

func DomainError(result ...*Result) error {
	for _, r := range result {
		r.val = math.NaN()
		r.err = math.NaN()
	}
	return err.EDOM
}

func DomainErrorMsg(msg string, result *Result) error {
	result.val = math.NaN()
	result.err = math.NaN()
	return err.EDOM
}

func DomainError_e10(result *Result_e10) error {
	result.val = math.NaN()
	result.err = math.NaN()
	result.e10 = 0
	return err.EDOM
}

func OverflowError_e10(result *Result_e10) error {
	result.val = math.Inf(1)
	result.err = math.Inf(1)
	result.e10 = 0
	return err.EOVRFLW
}

func UnderflowError_e10(result *Result_e10) error {
	result.val = 0.0
	result.err = gsl.DBL_MIN
	result.e10 = 0
	return err.EUNDRFLW
}
