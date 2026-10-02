#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    vector<long long> ans;
    for (long long i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            ans.push_back(i);

            if (i != n / i) {
                ans.push_back(n / i);
            }
        }
    }
    sort(ans.begin(), ans.end());

    for (auto &num: ans) {
        cout << num << " ";
    }

    return 0;
}