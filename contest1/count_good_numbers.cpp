#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    int cnt = 0;
    long long x;

    for (int i = 1; i <= n; i++) {
        cin >> x;

        if ((x != 0 && 18 % x == 0) || x % 45 == 0) {
            cnt++;
        }
    }
    
    cout << cnt;

    return 0;
}