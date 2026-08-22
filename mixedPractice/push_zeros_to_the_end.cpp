#include <bits/stdc++.h>
using namespace std;

void solve() {
    int n;
    cin >> n;
    vector<long long> arr(n);

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    int nextNonZero = 0;
    for (int i = 0; i < n; i++) {
        if (arr[i] != 0) {
            arr[nextNonZero++] = arr[i];
        }
    }

    while (nextNonZero < n) {
        arr[nextNonZero++] = 0;
    }

    for (int i = 0; i < n; i++) {
        cout << arr[i] << " ";
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        solve();
    }

    return 0;
}