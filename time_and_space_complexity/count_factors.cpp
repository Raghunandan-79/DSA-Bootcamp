#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    long long count = 0;
    for (long long i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            count++;

            if (i != n / i) {
                count++;
            }
        }
    }
    cout << count << "\n";

    return 0;
}