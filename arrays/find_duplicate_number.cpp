#include <bits/stdc++.h>
using namespace std;

void findDuplicate() {
    int n;
    cin >> n;

    long long sum = 0;
    for (int i = 0; i < n; i++) {
        int x;
        cin >> x;
        sum += x;
    }

    long long expected = 1LL * (n - 2) * (n - 1) / 2;
    cout << sum - expected << endl;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        findDuplicate();
    }
    
    return 0;
}