package specfunc

import (
	gsl "github.com/swook/gogsl"
	"github.com/swook/gogsl/err"
	"math"
)

const (
	LogRootTwoPi_ = 0.9189385332046727418
	GAMMA_XMAX    = 171.0
	/* The maximum n such that gsl_sf_fact(n) does not give an overflow. */
	FACT_NMAX = 170.0
	/* The maximum n such that gsl_sf_doublefact(n) does not give an overflow. */
	DOUBLEFACT_NMAX = 297.0
)

type fact struct {
	n int
	f float64
	i int64
}

var fact_table = []fact{
	{0, 1.0, 1},
	{1, 1.0, 1},
	{2, 2.0, 2},
	{3, 6.0, 6},
	{4, 24.0, 24},
	{5, 120.0, 120},
	{6, 720.0, 720},
	{7, 5040.0, 5040},
	{8, 40320.0, 40320},

	{9, 362880.0, 362880},
	{10, 3628800.0, 3628800},
	{11, 39916800.0, 39916800},
	{12, 479001600.0, 479001600},

	{13, 6227020800.0, 0},
	{14, 87178291200.0, 0},
	{15, 1307674368000.0, 0},
	{16, 20922789888000.0, 0},
	{17, 355687428096000.0, 0},
	{18, 6402373705728000.0, 0},
	{19, 121645100408832000.0, 0},
	{20, 2432902008176640000.0, 0},
	{21, 51090942171709440000.0, 0},
	{22, 1124000727777607680000.0, 0},
	{23, 25852016738884976640000.0, 0},
	{24, 620448401733239439360000.0, 0},
	{25, 15511210043330985984000000.0, 0},
	{26, 403291461126605635584000000.0, 0},
	{27, 10888869450418352160768000000.0, 0},
	{28, 304888344611713860501504000000.0, 0},
	{29, 8841761993739701954543616000000.0, 0},
	{30, 265252859812191058636308480000000.0, 0},
	{31, 8222838654177922817725562880000000.0, 0},
	{32, 263130836933693530167218012160000000.0, 0},
	{33, 8683317618811886495518194401280000000.0, 0},
	{34, 2.95232799039604140847618609644e38, 0},
	{35, 1.03331479663861449296666513375e40, 0},
	{36, 3.71993326789901217467999448151e41, 0},
	{37, 1.37637530912263450463159795816e43, 0},
	{38, 5.23022617466601111760007224100e44, 0},
	{39, 2.03978820811974433586402817399e46, 0},
	{40, 8.15915283247897734345611269600e47, 0},
	{41, 3.34525266131638071081700620534e49, 0},
	{42, 1.40500611775287989854314260624e51, 0},
	{43, 6.04152630633738356373551320685e52, 0},
	{44, 2.65827157478844876804362581101e54, 0},
	{45, 1.19622220865480194561963161496e56, 0},
	{46, 5.50262215981208894985030542880e57, 0},
	{47, 2.58623241511168180642964355154e59, 0},
	{48, 1.24139155925360726708622890474e61, 0},
	{49, 6.08281864034267560872252163321e62, 0},
	{50, 3.04140932017133780436126081661e64, 0},
	{51, 1.55111875328738228022424301647e66, 0},
	{52, 8.06581751709438785716606368564e67, 0},
	{53, 4.27488328406002556429801375339e69, 0},
	{54, 2.30843697339241380472092742683e71, 0},
	{55, 1.26964033536582759259651008476e73, 0},
	{56, 7.10998587804863451854045647464e74, 0},
	{57, 4.05269195048772167556806019054e76, 0},
	{58, 2.35056133128287857182947491052e78, 0},
	{59, 1.38683118545689835737939019720e80, 0},
	{60, 8.32098711274139014427634118320e81, 0},
	{61, 5.07580213877224798800856812177e83, 0},
	{62, 3.14699732603879375256531223550e85, 0},
	{63, 1.982608315404440064116146708360e87, 0},
	{64, 1.268869321858841641034333893350e89, 0},
	{65, 8.247650592082470666723170306800e90, 0},
	{66, 5.443449390774430640037292402480e92, 0},
	{67, 3.647111091818868528824985909660e94, 0},
	{68, 2.480035542436830599600990418570e96, 0},
	{69, 1.711224524281413113724683388810e98, 0},
	{70, 1.197857166996989179607278372170e100, 0},
	{71, 8.504785885678623175211676442400e101, 0},
	{72, 6.123445837688608686152407038530e103, 0},
	{73, 4.470115461512684340891257138130e105, 0},
	{74, 3.307885441519386412259530282210e107, 0},
	{75, 2.480914081139539809194647711660e109, 0},
	{76, 1.885494701666050254987932260860e111, 0},
	{77, 1.451830920282858696340707840860e113, 0},
	{78, 1.132428117820629783145752115870e115, 0},
	{79, 8.946182130782975286851441715400e116, 0},
	{80, 7.156945704626380229481153372320e118, 0},
	{81, 5.797126020747367985879734231580e120, 0},
	{82, 4.753643337012841748421382069890e122, 0},
	{83, 3.945523969720658651189747118010e124, 0},
	{84, 3.314240134565353266999387579130e126, 0},
	{85, 2.817104114380550276949479442260e128, 0},
	{86, 2.422709538367273238176552320340e130, 0},
	{87, 2.107757298379527717213600518700e132, 0},
	{88, 1.854826422573984391147968456460e134, 0},
	{89, 1.650795516090846108121691926250e136, 0},
	{90, 1.485715964481761497309522733620e138, 0},
	{91, 1.352001527678402962551665687590e140, 0},
	{92, 1.243841405464130725547532432590e142, 0},
	{93, 1.156772507081641574759205162310e144, 0},
	{94, 1.087366156656743080273652852570e146, 0},
	{95, 1.032997848823905926259970209940e148, 0},
	{96, 9.916779348709496892095714015400e149, 0},
	{97, 9.619275968248211985332842594960e151, 0},
	{98, 9.426890448883247745626185743100e153, 0},
	{99, 9.332621544394415268169923885600e155, 0},
	{100, 9.33262154439441526816992388563e157, 0},
	{101, 9.42594775983835942085162312450e159, 0},
	{102, 9.61446671503512660926865558700e161, 0},
	{103, 9.90290071648618040754671525458e163, 0},
	{104, 1.02990167451456276238485838648e166, 0},
	{105, 1.08139675824029090050410130580e168, 0},
	{106, 1.146280563734708354534347384148e170, 0},
	{107, 1.226520203196137939351751701040e172, 0},
	{108, 1.324641819451828974499891837120e174, 0},
	{109, 1.443859583202493582204882102460e176, 0},
	{110, 1.588245541522742940425370312710e178, 0},
	{111, 1.762952551090244663872161047110e180, 0},
	{112, 1.974506857221074023536820372760e182, 0},
	{113, 2.231192748659813646596607021220e184, 0},
	{114, 2.543559733472187557120132004190e186, 0},
	{115, 2.925093693493015690688151804820e188, 0},
	{116, 3.393108684451898201198256093590e190, 0},
	{117, 3.96993716080872089540195962950e192, 0},
	{118, 4.68452584975429065657431236281e194, 0},
	{119, 5.57458576120760588132343171174e196, 0},
	{120, 6.68950291344912705758811805409e198, 0},
	{121, 8.09429852527344373968162284545e200, 0},
	{122, 9.87504420083360136241157987140e202, 0},
	{123, 1.21463043670253296757662432419e205, 0},
	{124, 1.50614174151114087979501416199e207, 0},
	{125, 1.88267717688892609974376770249e209, 0},
	{126, 2.37217324288004688567714730514e211, 0},
	{127, 3.01266001845765954480997707753e213, 0},
	{128, 3.85620482362580421735677065923e215, 0},
	{129, 4.97450422247728744039023415041e217, 0},
	{130, 6.46685548922047367250730439554e219, 0},
	{131, 8.47158069087882051098456875820e221, 0},
	{132, 1.11824865119600430744996307608e224, 0},
	{133, 1.48727070609068572890845089118e226, 0},
	{134, 1.99294274616151887673732419418e228, 0},
	{135, 2.69047270731805048359538766215e230, 0},
	{136, 3.65904288195254865768972722052e232, 0},
	{137, 5.01288874827499166103492629211e234, 0},
	{138, 6.91778647261948849222819828311e236, 0},
	{139, 9.61572319694108900419719561353e238, 0},
	{140, 1.34620124757175246058760738589e241, 0},
	{141, 1.89814375907617096942852641411e243, 0},
	{142, 2.69536413788816277658850750804e245, 0},
	{143, 3.85437071718007277052156573649e247, 0},
	{144, 5.55029383273930478955105466055e249, 0},
	{145, 8.04792605747199194484902925780e251, 0},
	{146, 1.17499720439091082394795827164e254, 0},
	{147, 1.72724589045463891120349865931e256, 0},
	{148, 2.55632391787286558858117801578e258, 0},
	{149, 3.80892263763056972698595524351e260, 0},
	{150, 5.71338395644585459047893286526e262, 0},
	{151, 8.62720977423324043162318862650e264, 0},
	{152, 1.31133588568345254560672467123e267, 0},
	{153, 2.00634390509568239477828874699e269, 0},
	{154, 3.08976961384735088795856467036e271, 0},
	{155, 4.78914290146339387633577523906e273, 0},
	{156, 7.47106292628289444708380937294e275, 0},
	{157, 1.17295687942641442819215807155e278, 0},
	{158, 1.85327186949373479654360975305e280, 0},
	{159, 2.94670227249503832650433950735e282, 0},
	{160, 4.71472363599206132240694321176e284, 0},
	{161, 7.59070505394721872907517857094e286, 0},
	{162, 1.22969421873944943411017892849e289, 0},
	{163, 2.00440157654530257759959165344e291, 0},
	{164, 3.28721858553429622726333031164e293, 0},
	{165, 5.42391066613158877498449501421e295, 0},
	{166, 9.00369170577843736647426172359e297, 0},
	{167, 1.50361651486499904020120170784e300, 0},
	{168, 2.52607574497319838753801886917e302, 0},
	{169, 4.26906800900470527493925188890e304, 0},
	{170, 7.25741561530799896739672821113e306, 0},
}

var lanczos_7_c = [9]float64{
	0.99999999999980993227684700473478,
	676.520368121885098567009190444019,
	-1259.13921672240287047156078755283,
	771.3234287776530788486528258894,
	-176.61502916214059906584551354,
	12.507343278686904814458936853,
	-0.13857109526572011689554707,
	9.984369578019570859563e-6,
	1.50563273514931155834e-7,
}

func Lngamma_lanczos(x float64, result *Result) error {
	var Ag, term1, term2 float64

	x -= 1.0 /* Lanczos writes z! instead of Gamma(z) */

	Ag = lanczos_7_c[0]
	for k := 1; k <= 8; k++ {
		Ag += lanczos_7_c[k] / (x + float64(k))
	}

	/* (x+0.5)*math.Log(x+7.5) - (x+7.5) + LogRootTwoPi_ + math.Log(Ag(x)) */
	term1 = (x + 0.5) * math.Log((x+7.5)/gsl.M_E)
	term2 = LogRootTwoPi_ + math.Log(Ag)
	result.val = term1 + (term2 - 7.0)
	result.err = 2.0 * gsl.DBL_EPSILON * (math.Abs(term1) + math.Abs(term2) + 7.0)
	result.err += gsl.DBL_EPSILON * math.Abs(result.val)

	return err.SUCCESS
}

func Lngamma_sgn_0(eps float64, lng *Result, sgn *float64) error {
	/* calculate series for g(eps) = Gamma(eps) eps - 1/(1+eps) - eps/2 */
	c1 := -0.07721566490153286061
	c2 := -0.01094400467202744461
	c3 := 0.09252092391911371098
	c4 := -0.01827191316559981266
	c5 := 0.01800493109685479790
	c6 := -0.00685088537872380685
	c7 := 0.00399823955756846603
	c8 := -0.00189430621687107802
	c9 := 0.00097473237804513221
	c10 := -0.00048434392722255893
	g6 := c6 + eps*(c7+eps*(c8+eps*(c9+eps*c10)))
	g := eps * (c1 + eps*(c2+eps*(c3+eps*(c4+eps*(c5+eps*g6)))))

	/* calculate Gamma(eps) eps, a positive quantity */
	gee := g + 1.0/(1.0+eps) + 0.5*eps

	lng.val = math.Log(gee / math.Abs(eps))
	lng.err = 4.0 * gsl.DBL_EPSILON * math.Abs(lng.val)
	*sgn = gsl.Sign(eps)

	return err.SUCCESS
}

func Lngamma_sgn_sing(N int, eps float64, lng *Result, sgn *float64) error {
	if eps == 0.0 {
		lng.val = 0.0
		lng.err = 0.0
		*sgn = 0.0
		return err.EDOM
	} else if N == 1 {
		/* calculate series for
		* g = eps gamma(-1+eps) + 1 + eps/2 (1+3eps)/(1-eps^2)
		* double-precision for |eps| < 0.02
		 */
		c0 := 0.07721566490153286061
		c1 := 0.08815966957356030521
		c2 := -0.00436125434555340577
		c3 := 0.01391065882004640689
		c4 := -0.00409427227680839100
		c5 := 0.00275661310191541584
		c6 := -0.00124162645565305019
		c7 := 0.00065267976121802783
		c8 := -0.00032205261682710437
		c9 := 0.00016229131039545456
		g5 := c5 + eps*(c6+eps*(c7+eps*(c8+eps*c9)))
		g := eps * (c0 + eps*(c1+eps*(c2+eps*(c3+eps*(c4+eps*g5)))))

		/* calculate eps gamma(-1+eps), a negative quantity */
		gam_e := g - 1.0 - 0.5*eps*(1.0+3.0*eps)/(1.0-eps*eps)

		lng.val = math.Log(math.Abs(gam_e) / math.Abs(eps))
		lng.err = 2.0 * gsl.DBL_EPSILON * math.Abs(lng.val)
		if eps > 0.0 {
			*sgn = -1.0
		} else {
			*sgn = 1.0
		}
		return err.SUCCESS
	}

	var g float64

	/* series for sin(Pi(N+1-eps))/(Pi eps) modulo the sign
	* double-precision for |eps| < 0.02
	 */
	cs1 := -1.6449340668482264365
	cs2 := 0.8117424252833536436
	cs3 := -0.1907518241220842137
	cs4 := 0.0261478478176548005
	cs5 := -0.0023460810354558236
	e2 := eps * eps
	sin_ser := 1.0 + e2*(cs1+e2*(cs2+e2*(cs3+e2*(cs4+e2*cs5))))

	/* calculate series for ln(gamma(1+N-eps))
	* double-precision for |eps| < 0.02
	 */
	aeps := math.Abs(eps)
	var c1, c2, c3, c4, c5, c6, c7, lng_ser float64
	var c0, psi_0, psi_1, psi_2, psi_3, psi_4, psi_5, psi_6 *Result
	psi_2.val = 0.0
	psi_3.val = 0.0
	psi_4.val = 0.0
	psi_5.val = 0.0
	psi_6.val = 0.0
	Lnfact_e(N, c0)
	Psi_int_e(N+1, psi_0)
	Psi_1_int_e(N+1, psi_1)
	switch {
	case aeps > 0.00001:
		Psi_n_e(2, float64(N)+1.0, psi_2)
	case aeps > 0.0002:
		Psi_n_e(3, float64(N)+1.0, psi_3)
	case aeps > 0.001:
		Psi_n_e(4, float64(N)+1.0, psi_4)
	case aeps > 0.005:
		Psi_n_e(5, float64(N)+1.0, psi_5)
	case aeps > 0.01:
		Psi_n_e(6, float64(N)+1.0, psi_6)
	}
	c1 = psi_0.val
	c2 = psi_1.val / 2.0
	c3 = psi_2.val / 6.0
	c4 = psi_3.val / 24.0
	c5 = psi_4.val / 120.0
	c6 = psi_5.val / 720.0
	c7 = psi_6.val / 5040.0
	lng_ser = c0.val - eps*(c1-eps*(c2-eps*(c3-eps*(c4-eps*(c5-eps*(c6-eps*c7))))))

	/* calculate
	* g = ln(|eps gamma(-N+eps)|)
	*   = -ln(gamma(1+N-eps)) + ln(|eps Pi/sin(Pi(N+1+eps))|)
	 */
	g = -lng_ser - math.Log(sin_ser)

	lng.val = g - math.Log(math.Abs(eps))
	lng.err = c0.err + 2.0*gsl.DBL_EPSILON*(math.Abs(g)+math.Abs(lng.val))

	if eps > 0.0 {
		*sgn = 1.0
	} else {
		*sgn = -1.0
	}
	if gsl.IsOdd(N) {
		*sgn *= -1.0
	} else {
		*sgn *= 1.0
	}

	return err.SUCCESS
}

func Lngamma_1_pade(eps float64, result *Result) error {
	/* Use (2,2) Pade for Log[Gamma[1+eps]]/eps
	* plus a correction series.
	 */
	n1 := -1.0017419282349508699871138440
	n2 := 1.7364839209922879823280541733
	d1 := 1.2433006018858751556055436011
	d2 := 5.0456274100274010152489597514
	num := (eps + n1) * (eps + n2)
	den := (eps + d1) * (eps + d2)
	pade := 2.0816265188662692474880210318 * num / den
	c0 := 0.004785324257581753
	c1 := -0.01192457083645441
	c2 := 0.01931961413960498
	c3 := -0.02594027398725020
	c4 := 0.03141928755021455
	eps5 := eps * eps * eps * eps * eps
	corr := eps5 * (c0 + eps*(c1+eps*(c2+eps*(c3+c4*eps))))
	result.val = eps * (pade + corr)
	result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	return err.SUCCESS
}

func Lngamma_2_pade(eps float64, result *Result) error {
	/* Use (2,2) Pade for Log[Gamma[2+eps]]/eps
	* plus a correction series.
	 */
	n1 := 1.000895834786669227164446568
	n2 := 4.209376735287755081642901277
	d1 := 2.618851904903217274682578255
	d2 := 10.85766559900983515322922936
	num := (eps + n1) * (eps + n2)
	den := (eps + d1) * (eps + d2)
	pade := 2.85337998765781918463568869 * num / den
	c0 := 0.0001139406357036744
	c1 := -0.0001365435269792533
	c2 := 0.0001067287169183665
	c3 := -0.0000693271800931282
	c4 := 0.0000407220927867950
	eps5 := eps * eps * eps * eps * eps
	corr := eps5 * (c0 + eps*(c1+eps*(c2+eps*(c3+c4*eps))))
	result.val = eps * (pade + corr)
	result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	return err.SUCCESS
}

func Lngamma_e(x float64, result *Result) error {
	if math.Abs(x-1.0) < 0.01 {
		/* Note that we must amplify the errors
		* from the Pade evaluations because of
		* the way we must pass the argument, i.e.
		* writing (1-x) is a loss of precision
		* when x is near 1.
		 */
		stat := Lngamma_1_pade(x-1.0, result)
		result.err *= 1.0 / (gsl.DBL_EPSILON + math.Abs(x-1.0))
		return stat
	} else if math.Abs(x-2.0) < 0.01 {
		stat := Lngamma_2_pade(x-2.0, result)
		result.err *= 1.0 / (gsl.DBL_EPSILON + math.Abs(x-2.0))
		return stat
	} else if x >= 0.5 {
		return Lngamma_lanczos(x, result)
	} else if x == 0.0 {
		return DomainError(result)
	} else if math.Abs(x) < 0.02 {
		var sgn float64
		return Lngamma_sgn_0(x, result, &sgn)
	} else if x > -0.5/(gsl.DBL_EPSILON*gsl.M_PI) {
		/* Try to extract a fractional
		* part from x.
		 */
		z := 1.0 - x
		s := math.Sin(gsl.M_PI * z)
		as := math.Abs(s)
		if s == 0.0 {
			return DomainError(result)
		} else if as < gsl.M_PI*0.015 {
			/* x is near a negative integer, -N */
			if x < math.MinInt32+2.0 {
				result.val = 0.0
				result.err = 0.0
				return err.EROUND
			} else {
				N := -int(x - 0.5)
				eps := x + float64(N)
				var sgn float64
				return Lngamma_sgn_sing(N, eps, result, &sgn)
			}
		} else {
			var lg_z *Result
			Lngamma_lanczos(z, lg_z)
			result.val = gsl.M_LNPI - (math.Log(as) + lg_z.val)
			result.err = 2.0*gsl.DBL_EPSILON*math.Abs(result.val) + lg_z.err
			return err.SUCCESS
		}
	}

	/* |x| was too large to extract any fractional part */
	result.val = 0.0
	result.err = 0.0
	return err.EROUND
}

func Lnfact_e(n int, result *Result) error {
	if n <= FACT_NMAX {
		result.val = math.Log(fact_table[n].f)
		result.err = 2.0 * gsl.DBL_EPSILON * math.Abs(result.val)
	} else {
		Lngamma_e(float64(n)+1.0, result)
	}
	return err.SUCCESS
}