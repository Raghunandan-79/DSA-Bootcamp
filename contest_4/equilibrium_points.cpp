#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    long long total = accumulate(arr.begin(), arr.end(), 0LL);
    long long prefix = 0;
    long long equilibriumPoints = 0;

    for (long long value : arr) {
        prefix += value;
        if (prefix == total - prefix + value) {
            ++equilibriumPoints;
        }
    }

    cout << equilibriumPoints << '\n';

    return 0;
}