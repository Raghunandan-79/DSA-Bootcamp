#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;
    string s;
    cin >> s;

    vector<long long> prefix(n);
    if (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u') {
        prefix[0] = 1;
    }
    else {
        prefix[0] = 0;
    }

    for (long long i = 1; i < n; i++) {
        if (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u') {
            prefix[i] = prefix[i - 1] + 1;
        }
        else {
            prefix[i] = prefix[i - 1];
        }
    }

    long long q;
    cin >> q;
    while (q--) {
        long long l, r;
        cin >> l >> r;
        l--, r--;

        if (l == 0) {
            cout << prefix[r] << "\n";
        }
        else {
            cout << prefix[r] - prefix[l - 1] << "\n";
        }
    }

    return 0;
}