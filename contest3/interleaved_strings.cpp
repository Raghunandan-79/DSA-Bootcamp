#include <bits/stdc++.h>
using namespace std;

int main() {
    string s, t;
    cin >> s >> t;

    string ans;
    int i = 0, j = 0;
    while (i < s.size() && j < t.size()) {
        ans += s[i];
        ans += '-';
        ans += t[j];
        ans += '-';
        i++;
        j++;
    }

    while (i < s.size()) {
        ans += s[i];
        ans += '-';
        i++;
    }

    while (j < t.size()) {
        ans += t[j];
        ans += '-';
        j++;
    }

    ans.pop_back();
    cout << ans << endl;
    
    return 0;
}