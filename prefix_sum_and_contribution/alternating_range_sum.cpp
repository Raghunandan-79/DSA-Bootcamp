#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, q;
    cin >> n >> q;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    vector<long long> prefix(n);
    prefix[0] = arr[0];
    for (long long i = 1; i < n; i++) {
        long long val;
        if (i % 2 == 0) {
            val = arr[i];
        }
        else {
            val = -arr[i];
        }

        prefix[i] = prefix[i - 1] + val;
    }


    while (q--) {
        long long l, r;
        cin >> l >> r;
        l--, r--;

        long long ans = prefix[r];
        if (l > 0) {
            ans -= prefix[l - 1];
        }
        
        if (l % 2 == 1) {
            ans = -ans;
        }

        cout << ans << "\n";
    }

    return 0;
}