#include <bits/stdc++.h>
using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    long long n;
    cin >> n;

    vector<long long> arr(n);
    for (auto &x : arr) {
        cin >> x;
    }

    vector<long long> prefix(n + 1, 0);

    for (long long i = 0; i < n; i++) {
        prefix[i + 1] = prefix[i] + arr[i];
    }

    for (long long i = 0; i < n; i++) {
        for (long long j = i; j < n; j++) {
            long long sum = prefix[j + 1] - prefix[i];
            cout << sum << "\n";
        }
    }

    return 0;
}