#include <bits/stdc++.h>
using namespace std;

long long factorial(int n) {
    long long fact = 1;

    for (int i = 1; i <= n; i++) {
        fact *= i;
    }

    return fact;
}

int main() {
    int n, r;
    cin >> n >> r;

    long long nFactorial = factorial(n);
    long long rFactorial = factorial(r);
    long long nMinusRFactorial = factorial(n - r);

    cout << nFactorial / (rFactorial * nMinusRFactorial) << endl;

    return 0;
}