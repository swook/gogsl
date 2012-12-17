package err

import "errors"

/* Port from err/gsl_errno.h */
var (
	CONTINUE = errors.New("Iteration has not converged")
	FAILURE  = errors.New("General failure")
	SUCCESS  = error(nil)
	EDOM     = errors.New("Input domain error, e.g sqrt(-1)")
	ERANGE   = errors.New("Output range error, e.g. exp(1e100)")
	EFAULT   = errors.New("Invalid pointer")
	EINVAL   = errors.New("Invalid argument supplied by user")
	EFAILED  = errors.New("Generic failure")
	EFACTOR  = errors.New("Factorization failed")
	ESANITY  = errors.New("Sanity check failed - shouldn't happen")
	ENOMEM   = errors.New("Malloc failed")
	EBADFUNC = errors.New("Problem with user-supplied function")
	ERUNAWAY = errors.New("Iterative process is out of control")
	EMAXITER = errors.New("Exceeded max number of iterations")
	EZERODIV = errors.New("Tried to divide by zero")
	EBADTOL  = errors.New("User specified an invalid tolerance")
	ETOL     = errors.New("Failed to reach the specified tolerance")
	EUNDRFLW = errors.New("Underflow")
	EOVRFLW  = errors.New("Overflow")
	ELOSS    = errors.New("Loss of accuracy")
	EROUND   = errors.New("Failed because of roundoff error")
	EBADLEN  = errors.New("Matrix, vector lengths are not conformant")
	ENOTSQR  = errors.New("Matrix not square")
	ESING    = errors.New("Apparent singularity detected")
	EDIVERGE = errors.New("Integral or series is divergent")
	EUNSUP   = errors.New("Requested feature is not supported by the hardware")
	EUNIMPL  = errors.New("Requested feature not (yet) implemented")
	ECACHE   = errors.New("Cache limit exceeded")
	ETABLE   = errors.New("Table limit exceeded")
	ENOPROG  = errors.New("Iteration is not making progress towards solution")
	ENOPROGJ = errors.New("Jacobian evaluations are not improving the solution")
	ETOLF    = errors.New("Cannot reach the specified tolerance in F")
	ETOLX    = errors.New("Cannot reach the specified tolerance in X")
	ETOLG    = errors.New("Cannot reach the specified tolerance in gradient")
	EOF      = errors.New("End of file")
)

// If multiple errors, select first non-Success one
func Select(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return SUCCESS
}
