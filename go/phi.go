package phi

import "math"

func Phi(x float64) float64 {
    // constants
    const a1 =  0.254829592
    const a2 = -0.284496736
    const a3 =  1.421413741
    const a4 = -1.453152027
    const a5 =  1.061405429
    const p  =  0.3275911

    // Save the sign of x
    sign := 1.0
    if x < 0 {
        sign = -1.0
    }
    x = math.Fabs(x)/math.Sqrt(2.0)

    // A&S formula 7.1.26
    t := 1.0/(1.0 + p*x)
    y := 1.0 - (((((a5*t + a4)*t) + a3)*t + a2)*t + a1)*t*math.Exp(-x*x)

    return 0.5*(1.0 + sign*y)
}
