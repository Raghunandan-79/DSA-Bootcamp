#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, q;
    cin >> n >> q;
    vector<long long> arr(n);
    for (auto &it: arr) cin >> it;

    unordered_map<long long, long long> mpp;
    for (long long i = 0; i < n; i++) {
        mpp[arr[i]] = i + 1;
    } 

    while (q--) {
        long long x;
        cin >> x;

        if (mpp.count(x)) {
            cout << mpp[x] << "\n";
        }
        else {
            cout << -1 << "\n";
        }
    }

    return 0;
}