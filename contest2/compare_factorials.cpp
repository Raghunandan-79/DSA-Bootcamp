#include <bits/stdc++.h>
using namespace std;

long long factorial(long long n) {
    if (n == 0 || n == 1) return 1;
    if (n < 0) return -1;
    long long result = 1;
    for (long long i = 2; i <= n; i++) {
        if (result > LLONG_MAX / i) return -1;
        result *= i;
    }
    return result;
}

bool compare(long long a, long long b) {
    if (a == b) return true;
    if ((a == 0 && b == 1) || (a == 1 && b == 0)) return true;
    return false;
}

int main() {
    long long a, b;
    cin >> a >> b;

    if (compare(a, b)) {
        cout << "Yes" << endl;
    }
    else {
        cout << "No" << endl;
    }

    return 0;
}