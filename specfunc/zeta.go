package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

/* coefficients for Maclaurin summation in hzeta()
 * B_{2j}/(2j)!
 */
var hzeta_c = [15]float64{
	1.00000000000000000000000000000,
	0.083333333333333333333333333333,
	-0.00138888888888888888888888888889,
	0.000033068783068783068783068783069,
	-8.2671957671957671957671957672e-07,
	2.0876756987868098979210090321e-08,
	-5.2841901386874931848476822022e-10,
	1.3382536530684678832826980975e-11,
	-3.3896802963225828668301953912e-13,
	8.5860620562778445641359054504e-15,
	-2.1748686985580618730415164239e-16,
	5.5090028283602295152026526089e-18,
	-1.3954464685812523340707686264e-19,
	3.5347070396294674716932299778e-21,
	-8.9535174270375468504026113181e-23,
}

func Hzeta_e(s, q float64, result *Result) error {
	if s <= 1.0 || q <= 0.0 {
		return DomainError(result)
	}

	max_bits := 54.0
	ln_term0 := -s * math.Log(q)

	if ln_term0 < gsl.LOG_DBL_MIN+1.0 {
		return UnderflowError(result)
	} else if ln_term0 > gsl.LOG_DBL_MAX-1.0 {
		return OverflowError(result)
	} else if (s > max_bits && q < 1.0) || (s > 0.5*max_bits && q < 0.25) {
		result.val = math.Pow(q, -s)
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if s > 0.5*max_bits && q < 1.0 {
		p1 := math.Pow(q, -s)
		p2 := math.Pow(q/(1.0+q), s)
		p3 := math.Pow(q/(2.0+q), s)
		result.val = p1 * (1.0 + p2 + p3)
		result.err = gsl.DBL_EPSILON * (0.5*s + 2.0) * math.Abs(result.val)
		return err.SUCCESS
	}

	/* Euler-Maclaurin summation formula
	* [Moshier, p. 400, with several typo corrections]
	 */
	jmax := 12
	kmax := 10
	pmax := math.Pow(float64(kmax)+q, -s)
	scp := s
	pcp := pmax / (float64(kmax) + q)
	ans := pmax * ((float64(kmax)+q)/(s-1.0) + 0.5)

	for k := 0; k < kmax; k++ {
		ans += math.Pow(float64(k)+q, -s)
	}

	for j := 0; j <= jmax; j++ {
		delta := hzeta_c[j+1] * scp * pcp
		ans += delta
		if math.Abs(delta/ans) < 0.5*gsl.DBL_EPSILON {
			break
		}
		scp *= (s + float64(2*j) + 1.0) * (s + float64(2*j) + 2.0)
		pcp /= (float64(kmax) + q) * (float64(kmax) + q)
	}

	result.val = ans
	result.err = 2.0 * float64(jmax+1) * gsl.DBL_EPSILON * math.Abs(ans)
	return err.SUCCESS
}
