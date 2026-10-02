#include <bits/stdc++.h>
using namespace std;

long long digit_sum(long long n) {
    long long sum = 0;

    while (n != 0) {
        sum += n % 10;
        n /= 10;
    }

    return sum;
}

int main() {
    long long n, q, k;
    cin >> n >> q >> k;

    vector<long long> arr(n);
    for (auto &it: arr) cin >> it;

    vector<long long> prefix(n);
    if (digit_sum(arr[0]) == k) {
        prefix[0] = 1;
    }
    else {
        prefix[0] = 0;
    }

    for (long long i = 1; i < n; i++) {
        if (digit_sum(arr[i]) == k) {
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