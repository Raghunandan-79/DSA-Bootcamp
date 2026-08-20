#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    cin >> s;

    for (int i = 0; i < s.size(); i++) {
        if (isupper(s[i])) {
            s[i] = char(tolower(s[i]));
        }
        else {
            s[i] = char(toupper(s[i]));
        }
    }

    cout << s << endl;
    
    return 0;
}