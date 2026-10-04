#include <bits/stdc++.h>
using namespace std;

long long count_factors(long long n) {
    long long count = 0;

    for (long long i = 1; i <= sqrt(n); i++) {
        if (n % i == 0) {
            count++;

            if (i != n / i) {
                count++;
            }
        }
    }

    return count;
}

int main() {
    long long a, b;
    cin >> a >> b;

    if (count_factors(a) > count_factors(b)) {
        cout << "A" << "\n";
    }
    else if (count_factors(a) < count_factors(b)) {
        cout << "B" << "\n";
    }
    else {
        cout << "DRAW" << "\n";
    }

    return 0;
}