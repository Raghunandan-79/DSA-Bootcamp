#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    vector<long long> arr(n);
    for (auto &x : arr) cin >> x;

    vector<long long> prefix(n);

    prefix[0] = 0;

    for (int i = 1; i < n; i++) {
        prefix[i] = prefix[i - 1];

        if ((i + 1) % 2 == 0) {
            prefix[i] += arr[i];
        }
    }

    long long q;
    cin >> q;

    while (q--) {
        long long l, r;
        cin >> l >> r;

        l--;
        r--;

        long long ans = prefix[r];

        if (l > 0) {
            ans -= prefix[l - 1];
        }

        cout << ans << "\n";
    }

    return 0;
}