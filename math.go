package gsl

import (
	"math"
)

const (
	M_E        = 2.71828182845904523536028747135 /* e */
	M_LOG2E    = 1.44269504088896340735992468100 /* log_2 (e) */
	M_LOG10E   = 0.43429448190325182765112891892 /* log_10 (e) */
	M_SQRT2    = 1.41421356237309504880168872421 /* sqrt(2) */
	M_SQRT1_2  = 0.70710678118654752440084436210 /* sqrt(1/2) */
	M_SQRT3    = 1.73205080756887729352744634151 /* sqrt(3) */
	M_PI       = 3.14159265358979323846264338328 /* pi */
	M_PI_2     = 1.57079632679489661923132169164 /* pi/2 */
	M_PI_4     = 0.78539816339744830961566084582 /* pi/4 */
	M_SQRTPI   = 1.77245385090551602729816748334 /* sqrt(pi) */
	M_2_SQRTPI = 1.12837916709551257389615890312 /* 2/sqrt(pi) */
	M_1_PI     = 0.31830988618379067153776752675 /* 1/pi */
	M_2_PI     = 0.63661977236758134307553505349 /* 2/pi */
	M_LN10     = 2.30258509299404568401799145468 /* ln(10) */
	M_LN2      = 0.69314718055994530941723212146 /* ln(2) */
	M_LNPI     = 1.14472988584940017414342735135 /* ln(pi) */
	M_EULER    = 0.57721566490153286060651209008 /* Euler constant */
)

func IsEven(n int) bool {
	return n&1 == 0
}
func IsOdd(n int) bool {
	return !IsEven(n)
}
func Sign(x float64) (sign float64) {
	if x >= 0.0 {
		sign = 1.0
	} else {
		sign = -1.0
	}
	return
}

func IsReal(x float64) bool {
	return !math.IsNaN(x) && !math.IsInf(x, 0)
}

func Max(in ...interface{}) float64 {
	var max_i int
	flin := make([]float64, len(in))
	for i, v := range in {
		switch v.(type) {
		case int:
			flin[i] = float64(v.(int))
		case float64:
			flin[i] = v.(float64)
		}
		if flin[i] > flin[max_i] {
			max_i = i
		}
	}
	return flin[max_i]
}

func Min(in ...interface{}) float64 {
	var max_i int
	flin := make([]float64, len(in))
	for i, v := range in {
		switch v.(type) {
		case int:
			flin[i] = float64(v.(int))
		case float64:
			flin[i] = v.(float64)
		}
		if flin[i] < flin[max_i] {
			max_i = i
		}
	}
	return flin[max_i]
}
