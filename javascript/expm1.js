function expm1(x) {
    if (Math.abs(x) < 1e-5) {
        return x + 0.5*x*x;
    } else {
        return Math.exp(x) - 1.0;
    }
}
