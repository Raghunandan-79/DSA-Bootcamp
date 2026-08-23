#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    cin >> s;

    int aCount = 0, bCount = 0;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] == 'a') {
            aCount++;
        }
        else {
            bCount++;
        }
    }

    if (aCount == bCount) {
        cout << "YES" << endl;
    }
    else {
        cout << "NO" << endl;
    }

    return 0;
}