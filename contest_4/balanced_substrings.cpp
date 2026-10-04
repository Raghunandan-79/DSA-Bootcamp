#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n, q;
    cin >> n >> q;
    string s;
    cin >> s;

    vector<long long> pref_vowels(n);
    vector<long long> pref_consonents(n);

    if (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0] == 'A' || s[0] == 'E' || s[0] == 'I' || s[0] == 'O' || s[0] == 'U') {
        pref_vowels[0] = 1;
        pref_consonents[0] = 0;
    }
    else {
        pref_vowels[0] = 0;
        pref_consonents[0] = 1;
    }

    for (int i = 1; i < n; i++) {
        if (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U') {
            pref_vowels[i] = pref_vowels[i - 1] + 1;
            pref_consonents[i] = pref_consonents[i - 1];
        }
        else {
            pref_vowels[i] = pref_vowels[i - 1];
            pref_consonents[i] = pref_consonents[i - 1] + 1;
        }
    }

    while (q--) {
        long long l, r;
        cin >> l >> r;
        l--, r--;

        if (l == 0) {
            if (pref_vowels[r] == pref_consonents[r]) {
                cout << "YES" << "\n";
            }
            else {
                cout << "NO" << "\n";
            }
        }
        else {
            if ((pref_vowels[r] - pref_vowels[l - 1]) == (pref_consonents[r] - pref_consonents[l - 1])) {
                cout << "YES" << "\n";
            }
            else {
                cout << "NO" << "\n";
            }
        }
    }


    return 0;
}