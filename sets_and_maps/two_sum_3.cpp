#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;

    unordered_map<long long, long long> mpp;
    long long ans = 0;

    for (long long i = 0; i < n; i++) {
        long long x;
        cin >> x;
        ans += mpp[k - x];
        mpp[x]++;
    }
    cout << ans << "\n";

    return 0;
}