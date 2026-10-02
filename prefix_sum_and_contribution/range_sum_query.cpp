#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    
    vector<long long> arr(n);
    for (auto &it: arr) {
        cin >> it;
    }

    long long q;
    cin >> q;

    vector<long long> prefix(n);
    prefix[0] = arr[0];
    for (int i = 1; i < n; i++) {
        prefix[i] = prefix[i - 1] + arr[i];
    }

    while (q--) {
        long long l, r;
        cin >> l >> r;
        l--;
        r--;

        if (l == 0) {
            cout << prefix[r] << "\n";
        }
        else {
            cout << prefix[r] - prefix[l - 1] << "\n";
        }
    }

    return 0;
}