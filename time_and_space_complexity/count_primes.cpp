#include <bits/stdc++.h>
using namespace std;

bool is_prime(long long n) {
    if (n <= 1) {
        return false;
    }

    for (long long i = 2; i <= sqrt(n); i++) {
        if (n % i == 0) {
            return false;
        }
    }

    return true;
}

int main() {
    long long n;
    cin >> n;

    long long count = 0;
    for (long long i = 2; i <= n; i++) {
        if (is_prime(i)) {
            count++;
        }
    }
    cout << count << "\n";

    return 0;
}