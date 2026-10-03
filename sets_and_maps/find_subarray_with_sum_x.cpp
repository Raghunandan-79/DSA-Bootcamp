#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;
    vector<long long> arr(n);
    for (auto &it: arr) {
        cin >> it;
    } 

    map<long long, long long> mpp;
    mpp[0] = 0;

    long long pref = 0;
    for (long long i = 1; i <= n; i++) {
        pref += arr[i - 1];

        if (mpp.count(pref - k)) {
            cout << mpp[pref - k] + 1 << " " << i << "\n";
            return 0;
        }

        if (!mpp.count(pref)) {
            mpp[pref] = i;
        }
    }

    cout << -1 << "\n";

    return 0;
}