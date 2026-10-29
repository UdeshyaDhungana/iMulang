
$ recursion wala program $
thoos recursion = kaam_gar(x) {
    yedi (x == 0) {
        patha 1;
    } natra {
        patha x * recursion(x - 1);
    }
};

bhan(recursion(4));