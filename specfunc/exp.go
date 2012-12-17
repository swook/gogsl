package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

func Exprel_n_CF(N int, x float64, result *Result) error {
	RECUR_BIG := gsl.SQRT_DBL_MAX
	maxiter := 5000
	n := 1
	Anm2 := 1.0
	Bnm2 := 0.0
	Anm1 := 0.0
	Bnm1 := 1.0
	a1 := 1.0
	b1 := 1.0
	a2 := -x
	b2 := float64(N) + 1.0
	var an, bn, old_fn, del float64

	An := b1*Anm1 + a1*Anm2 /* A1 */
	Bn := b1*Bnm1 + a1*Bnm2 /* B1 */

	/* One explicit step, before we get to the main pattern. */
	n++
	Anm2 = Anm1
	Bnm2 = Bnm1
	Anm1 = An
	Bnm1 = Bn
	An = b2*Anm1 + a2*Anm2 /* A2 */
	Bn = b2*Bnm1 + a2*Bnm2 /* B2 */

	fn := An / Bn

	for n < maxiter {
		n++
		Anm2 = Anm1
		Bnm2 = Bnm1
		Anm1 = An
		Bnm1 = Bn
		if gsl.IsOdd(n) {
			an = (float64(n-1) / 2.0) * x
		} else {
			an = -(float64(N) + (float64(n) / 2) - 1.0) * x
		}
		bn = float64(N + n - 1)
		An = bn*Anm1 + an*Anm2
		Bn = bn*Bnm1 + an*Bnm2

		if math.Abs(An) > RECUR_BIG || math.Abs(Bn) > RECUR_BIG {
			An /= RECUR_BIG
			Bn /= RECUR_BIG
			Anm1 /= RECUR_BIG
			Bnm1 /= RECUR_BIG
			Anm2 /= RECUR_BIG
			Bnm2 /= RECUR_BIG
		}

		old_fn = fn
		fn = An / Bn
		del = old_fn / fn

		if math.Abs(del-1.0) < 2.0*gsl.DBL_EPSILON {
			break
		}
	}

	result.val = fn
	result.err = 2.0 * (float64(n) + 1.0) * gsl.DBL_EPSILON * math.Abs(fn)

	if n == maxiter {
		return err.EMAXITER
	}
	return err.SUCCESS
}

/*-*-*-*-*-*-*-*-*-*-*-* Functions with Error Codes *-*-*-*-*-*-*-*-*-*-*-*/

func Exp_e(x float64, result *Result) error {
	if x > gsl.LOG_DBL_MAX {
		return OverflowError(result)
	} else if x < gsl.LOG_DBL_MIN {
		return UnderflowError(result)
	}
	result.val = math.Exp(x)
	result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	return err.SUCCESS
}

func Exp_e10_e(x float64, result *Result_e10) error {
	if x > float64(math.MaxInt32-1) {
		return OverflowError_e10(result)
	} else if x < float64(math.MinInt32+1) {
		return UnderflowError_e10(result)
	}
	N := int(math.Floor(x / gsl.M_LN10))
	result.val = math.Exp(x - float64(N)*gsl.M_LN10)
	result.err = 2.0 * (math.Abs(x) + 1.0) * gsl.DBL_EPSILON * math.Abs(result.val)
	result.e10 = N
	return err.SUCCESS
}

func Exp_mult_e(x float64, y float64, result *Result) error {
	ay := math.Abs(y)

	if y == 0.0 {
		result.val = 0.0
		result.err = 0.0
		return err.SUCCESS
	} else if (x < 0.5*gsl.LOG_DBL_MAX && x > 0.5*gsl.LOG_DBL_MIN) && (ay < 0.8*gsl.SQRT_DBL_MAX && ay > 1.2*gsl.SQRT_DBL_MIN) {
		ex := math.Exp(x)
		result.val = y * ex
		result.err = (2.0 + math.Abs(x)) * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	}

	ly := math.Log(ay)
	lnr := x + ly

	if lnr > gsl.LOG_DBL_MAX-0.01 {
		return OverflowError(result)
	} else if lnr < gsl.LOG_DBL_MIN+0.01 {
		return UnderflowError(result)
	}

	sy := gsl.Sign(y)
	M := math.Floor(x)
	N := math.Floor(ly)
	a := x - M
	b := ly - N
	berr := 2.0 * gsl.DBL_EPSILON * (math.Abs(ly) + math.Abs(N))
	result.val = float64(sy) * math.Exp(M+N) * math.Exp(a+b)
	result.err = berr * math.Abs(result.val)
	result.err += 2.0 * gsl.DBL_EPSILON * (M + N + 1.0) * math.Abs(result.val)
	return err.SUCCESS
}

func Exp_mult_e10_e(x, y float64, result *Result_e10) error {
	ay := math.Abs(y)

	if y == 0.0 {
		result.val = 0.0
		result.err = 0.0
		result.e10 = 0
		return err.SUCCESS
	} else if (x < 0.5*gsl.LOG_DBL_MAX && x > 0.5*gsl.LOG_DBL_MIN) && (ay < 0.8*gsl.SQRT_DBL_MAX && ay > 1.2*gsl.SQRT_DBL_MIN) {
		ex := math.Exp(x)
		result.val = y * ex
		result.err = (2.0 + math.Abs(x)) * gsl.DBL_EPSILON * math.Abs(result.val)
		result.e10 = 0
		return err.SUCCESS
	}
	ly := math.Log(ay)
	l10_val := (x + ly) / gsl.M_LN10

	if l10_val > math.MaxInt32-1 {
		return OverflowError_e10(result)
	} else if l10_val < math.MinInt32+1 {
		return UnderflowError_e10(result)
	}

	sy := gsl.Sign(y)
	N := int(math.Floor(l10_val))
	arg_val := (l10_val - float64(N)) * gsl.M_LN10
	arg_err := 2.0 * gsl.DBL_EPSILON * math.Abs(ly)

	result.val = float64(sy) * math.Exp(arg_val)
	result.err = arg_err * math.Abs(result.val)
	result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	result.e10 = N

	return err.SUCCESS
}

func Exp_mult_err_e(x, dx, y, dy float64, result *Result) error {
	ay := math.Abs(y)

	if y == 0.0 {
		result.val = 0.0
		result.err = math.Abs(dy * math.Exp(x))
		return err.SUCCESS
	} else if (x < 0.5*gsl.LOG_DBL_MAX && x > 0.5*gsl.LOG_DBL_MIN) && (ay < 0.8*gsl.SQRT_DBL_MAX && ay > 1.2*gsl.SQRT_DBL_MIN) {
		ex := math.Exp(x)
		result.val = y * ex
		result.err = ex * (math.Abs(dy) + math.Abs(y*dx))
		result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	}

	ly := math.Log(ay)
	lnr := x + ly

	if lnr > gsl.LOG_DBL_MAX-0.01 {
		return OverflowError(result)
	} else if lnr < gsl.LOG_DBL_MIN+0.01 {
		return UnderflowError(result)
	}

	sy := gsl.Sign(y)
	M := math.Floor(x)
	N := math.Floor(ly)
	a := x - M
	b := ly - N
	eMN := math.Exp(M + N)
	eab := math.Exp(a + b)
	result.val = float64(sy) * eMN * eab
	result.err = eMN * eab * 2.0 * gsl.DBL_EPSILON
	result.err += eMN * eab * math.Abs(dy/y)
	result.err += eMN * eab * math.Abs(dx)
	return err.SUCCESS
}

func Exp_mult_err_e10_e(x, dx, y, dy float64, result *Result_e10) error {
	ay := math.Abs(y)

	if y == 0.0 {
		result.val = 0.0
		result.err = math.Abs(dy * math.Exp(x))
		result.e10 = 0
		return err.SUCCESS
	} else if (x < 0.5*gsl.LOG_DBL_MAX && x > 0.5*gsl.LOG_DBL_MIN) && (ay < 0.8*gsl.SQRT_DBL_MAX && ay > 1.2*gsl.SQRT_DBL_MIN) {
		ex := math.Exp(x)
		result.val = y * ex
		result.err = ex * (math.Abs(dy) + math.Abs(y*dx))
		result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		result.e10 = 0
		return err.SUCCESS
	}

	ly := math.Log(ay)
	l10_val := (x + ly) / gsl.M_LN10

	if l10_val > math.MaxInt32-1 {
		return OverflowError_e10(result)
	} else if l10_val < math.MinInt32+1 {
		return UnderflowError_e10(result)
	}

	sy := gsl.Sign(y)
	N := int(math.Floor(l10_val))
	arg_val := (l10_val - float64(N)) * gsl.M_LN10
	arg_err := dy/math.Abs(y) + dx + 2.0*gsl.DBL_EPSILON*math.Abs(arg_val)

	result.val = float64(sy) * math.Exp(arg_val)
	result.err = arg_err * math.Abs(result.val)
	result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	result.e10 = N

	return err.SUCCESS
}

func Expm1_e(x float64, result *Result) error {
	cut := 0.002

	if x < gsl.LOG_DBL_MIN {
		result.val = -1.0
		result.err = gsl.DBL_EPSILON
		return err.SUCCESS
	} else if x < -cut {
		result.val = math.Exp(x) - 1.0
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < cut {
		result.val = x * (1.0 + 0.5*x*(1.0+x/3.0*(1.0+0.25*x*(1.0+0.2*x))))
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < gsl.LOG_DBL_MAX {
		result.val = math.Exp(x) - 1.0
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	}
	return OverflowError(result)
}

func Exprel_e(x float64, result *Result) error {
	cut := 0.002

	if x < gsl.LOG_DBL_MIN {
		result.val = -1.0 / x
		result.err = gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < -cut {
		result.val = (math.Exp(x) - 1.0) / x
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < cut {
		result.val = (1.0 + 0.5*x*(1.0+x/3.0*(1.0+0.25*x*(1.0+0.2*x))))
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < gsl.LOG_DBL_MAX {
		result.val = (math.Exp(x) - 1.0) / x
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	}
	return OverflowError(result)
}

func Exprel_2_e(x float64, result *Result) error {
	cut := 0.002

	if x < gsl.LOG_DBL_MIN {
		result.val = -2.0 / x * (1.0 + 1.0/x)
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < -cut {
		result.val = 2.0 * (math.Exp(x) - 1.0 - x) / (x * x)
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < cut {
		result.val = (1.0 + 1.0/3.0*x*(1.0+0.25*x*(1.0+0.2*x*(1.0+1.0/6.0*x))))
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	} else if x < gsl.LOG_DBL_MAX {
		result.val = 2.0 * (math.Exp(x) - 1.0 - x) / (x * x)
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
		return err.SUCCESS
	}
	return OverflowError(result)
}

func Exprel_n_e(N int, x float64, result *Result) error {
	if N < 0 {
		return DomainError(result)
	} else if x == 0.0 {
		result.val = 1.0
		result.err = 0.0
		return err.SUCCESS
	} else if math.Abs(x) < gsl.ROOT3_DBL_EPSILON*float64(N) {
		result.val = 1.0 + x/float64(N+1)*(1.0+x/float64(N+2))
		result.err = 2.0 * gsl.DBL_EPSILON
		return err.SUCCESS
	} else if N == 0 {
		return Exp_e(x, result)
	} else if N == 1 {
		return Exprel_e(x, result)
	} else if N == 2 {
		return Exprel_2_e(x, result)
	}

	if x > float64(N) && (-x+float64(N)*(1.0+math.Log(x/float64(N))) < gsl.LOG_DBL_EPSILON) {
		/* x is much larger than n.
		* Ignore polynomial part, so
		* exprel_N(x) ~= e^x N!/x^N
		 */
		lnf_N := new(Result)
		Lnfact_e(N, lnf_N)
		lnterm := float64(N) * math.Log(x)
		lnr_val := x + lnf_N.val - lnterm
		lnr_err := gsl.DBL_EPSILON * (math.Abs(x) + math.Abs(lnf_N.val) + math.Abs(lnterm))
		lnr_err += lnf_N.err
		return Exp_err_e(lnr_val, lnr_err, result)
	} else if x > float64(N) {
		/* Write the identity
		*   exprel_n(x) = e^x n! / x^n (1 - Gamma[n,x]/Gamma[n])
		* then use the asymptotic expansion
		* Gamma[n,x] ~ x^(n-1) e^(-x) (1 + (n-1)/x + (n-1)(n-2)/x^2 + ...)
		 */
		ln_x := math.Log(x)
		lnf_N := new(Result)
		Lnfact_e(N, lnf_N)                       /* math.Log(N!)       */
		lg_N := lnf_N.val - math.Log(float64(N)) /* math.Log(Gamma(N)) */
		lnpre_val := x + lnf_N.val - float64(N)*ln_x
		lnpre_err := gsl.DBL_EPSILON * (math.Abs(x) + math.Abs(lnf_N.val) + math.Abs(float64(N)*ln_x))
		lnpre_err += lnf_N.err
		if lnpre_val < gsl.LOG_DBL_MAX-5.0 {
			bigG_ratio, pre := new(Result), new(Result)
			stat_ex := Exp_err_e(lnpre_val, lnpre_err, pre)
			ln_bigG_ratio_pre := -x + float64(N-1)*ln_x - lg_N
			bigGsum := 1.0
			term := 1.0
			for k := 1; k < N; k++ {
				term *= float64(N-k) / x
				bigGsum += term
			}
			stat_eG := Exp_mult_e(ln_bigG_ratio_pre, bigGsum, bigG_ratio)
			if stat_eG == err.SUCCESS {
				result.val = pre.val * (1.0 - bigG_ratio.val)
				result.err = pre.val * (2.0*gsl.DBL_EPSILON + bigG_ratio.err)
				result.err += pre.err * math.Abs(1.0-bigG_ratio.val)
				result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
				return stat_ex
			} else {
				result.val = 0.0
				result.err = 0.0
				return stat_eG
			}
		} else {
			OverflowError(result)
		}
	} else if x > -10.0*float64(N) {
		return Exprel_n_CF(N, x, result)
	}

	/* x . -Inf asymptotic:
	* exprel_n(x) ~ e^x n!/x^n - n/x (1 + (n-1)/x + (n-1)(n-2)/x + ...)
	*             ~ - n/x (1 + (n-1)/x + (n-1)(n-2)/x + ...)
	 */
	sum := 1.0
	term := 1.0
	for k := 1; k < N; k++ {
		term *= float64(N-k) / x
		sum += term
	}
	result.val = -float64(N) / x * sum
	result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	return err.SUCCESS
}

func Exp_err_e(x, dx float64, result *Result) error {
	adx := math.Abs(dx)

	if x+adx > gsl.LOG_DBL_MAX {
		OverflowError(result)
	} else if x-adx < gsl.LOG_DBL_MIN {
		UnderflowError(result)
	}

	ex := math.Exp(x)
	edx := math.Exp(adx)
	result.val = ex
	result.err = ex * gsl.Max(gsl.DBL_EPSILON, edx-1.0/edx)
	result.err += 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	return err.SUCCESS
}

func Exp_err_e10_e(x, dx float64, result *Result_e10) error {
	adx := math.Abs(dx)

	if x+adx > math.MaxInt32-1 {
		OverflowError_e10(result)
	} else if x-adx < math.MinInt32+1 {
		UnderflowError_e10(result)
	}

	N := int(math.Floor(x / gsl.M_LN10))
	ex := math.Exp(x - float64(N)*gsl.M_LN10)
	result.val = ex
	result.err = ex * (2.0*gsl.DBL_EPSILON*(math.Abs(x)+1.0) + adx)
	result.e10 = N
	return err.SUCCESS
}

/*-*-*-*-*-*-*-*-*-* Functions w/ Natural Prototypes *-*-*-*-*-*-*-*-*-*-*/

func Exp(x float64) (float64, error) {
	result := new(Result)
	status := Exp_e(x, result)
	return result.val, status
}

func Exp_mult(x, y float64) (float64, error) {
	result := new(Result)
	status := Exp_mult_e(x, y, result)
	return result.val, status
}

func Expm1(x float64) (float64, error) {
	result := new(Result)
	status := Expm1_e(x, result)
	return result.val, status
}

func Exprel(x float64) (float64, error) {
	result := new(Result)
	status := Exprel_e(x, result)
	return result.val, status
}

func Exprel_2(x float64) (float64, error) {
	result := new(Result)
	status := Exprel_2_e(x, result)
	return result.val, status
}

func Exprel_n(n int, x float64) (float64, error) {
	result := new(Result)
	status := Exprel_n_e(n, x, result)
	return result.val, status
}
