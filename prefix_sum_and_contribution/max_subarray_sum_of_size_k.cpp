#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, k;
    cin >> n >> k;

    vector<long long> arr(n);
    for (auto &x : arr) {
        cin >> x;
    }

    vector<long long> prefix(n + 1, 0);

    for (long long i = 0; i < n; i++) {
        prefix[i + 1] = prefix[i] + arr[i];
    }

    long long max_sum = LLONG_MIN;

    for (long long i = 0; i + k <= n; i++) {
        long long sum = prefix[i + k] - prefix[i];
        max_sum = max(max_sum, sum);
    }

    cout << max_sum << '\n';

    return 0;
}