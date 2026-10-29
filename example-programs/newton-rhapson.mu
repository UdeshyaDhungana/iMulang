
$ Program to demonstrate newton rhapson's method in Mulang $

thoos nrm = kaam_gar(a, b, c, initialGuess, tolerance, maxIterations) {
    thoos x = initialGuess;
    thoos fx = 0.0;
    thoos dfx = 0.0;
    thoos xNext = 0.0;

    ghuma(thoos i = 0; i < maxIterations; i = i + 1) {
		fx = a * x * x + b * x + c;
        dfx = 2 * a * x + b;

        yedi(abs(dfx) < 0.000000000000001) {
            patha "division by zero risk";
        }

        xNext = x - fx / dfx;

        yedi (abs(xNext - x) < tolerance) {
            patha xNext;
        }
        x = xNext;
	}

    patha "Failed to converge";
};

$ x² - 3x + 2 = 0 → roots are x=1 and x=2 $
thoos a = 1.0;
thoos b = -3.0;
thoos c = 2.0;

thoos res = nrm(a, b, c, 0, 0.0000001, 100);

bhan(res)