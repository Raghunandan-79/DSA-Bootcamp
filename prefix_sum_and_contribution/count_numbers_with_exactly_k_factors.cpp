#include <bits/stdc++.h>
using namespace std;

long long count_factors(long long n) {
    long long count = 0;

    for (int i = 1; i <= sqrt(n); i++) {
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
    long long n, q, k;
    cin >> n >> q >> k;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    vector<long long> prefix(n);
    if (count_factors(arr[0]) == k) {
        prefix[0] = 1;
    }
    else {
        prefix[0] = 0;
    }

    for (int i = 1; i < n; i++) {
        if (count_factors(arr[i]) == k) {
            prefix[i] = prefix[i - 1] + 1;
        }
        else {
            prefix[i] = prefix[i - 1];
        }
    }

    while (q--) {
        long long l, r;
        cin >> l >> r;
        l--, r--;

        if (l == 0) {
            cout << prefix[r] << "\n";
        }
        else {
            cout << prefix[r] - prefix[l - 1] << "\n";
        }
    }

    return 0;
}