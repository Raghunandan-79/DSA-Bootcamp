#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    vector<long long> arr(n);
    for (auto &x : arr) {
        cin >> x;
    }

    long long sum = 0;

    for (long long i = 0; i < n; i++) {
        sum += arr[i] * (i + 1) * (n - i);
    }

    cout << sum << '\n';

    return 0;
}