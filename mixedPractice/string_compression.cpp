#include <bits/stdc++.h>
using namespace std;

void solve() {
    string s;
    cin >> s;

    string compressed;
    for (int i = 0; i < (int)s.size(); ) {
        int j = i;
        while (j < (int)s.size() && s[j] == s[i]) j++;

        int cnt = j - i;
        compressed += s[i];
        if (cnt > 1) compressed += char('0' + cnt);
        i = j;
    }
    cout << compressed << '\n';
}

int main() {
    int T;
    cin >> T;

    while (T--) {
        solve();
    }

    return 0;
}