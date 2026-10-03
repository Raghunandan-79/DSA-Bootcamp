#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;
    vector<long long> arr(n);
    for (auto &it: arr) {
        cin >> it;
    }

    set<long long> seen;
    seen.insert(0);

    long long pref = 0;
    for (long long x: arr) {
        pref += x;

        if (seen.count(pref - k)) {
            cout << "YES" << "\n";
            return 0;
        }

        seen.insert(pref);
    }
    cout << "NO" << "\n";

    return 0;
}