#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    vector<long long> arr(n);

    for (auto &it: arr) {
        cin >> it;
    }

    long long sum = 0;
    for (long long i = 0; i < n; i++) {
        for (long long j = i; j < n; j++) {
            for (long long k = i; k <= j; k++) {
                sum += arr[k];
            }
        }
    }
    cout << sum;

    return 0;
}