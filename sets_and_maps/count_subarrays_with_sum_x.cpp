#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;

    unordered_map<long long, long long> mpp;
    mpp[0] = 1;

    long long pref = 0;
    long long ans = 0;

    for (long long i = 0; i < n; i++) {
        long long x;
        cin >> x;

        pref += x;
        ans += mpp[pref - k];
        mpp[pref]++;
    }

    cout << ans << "\n";

    return 0;
}