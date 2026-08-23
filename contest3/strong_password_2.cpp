#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);

    bool lowercase = false, uppercase = false, digit = false, special = false;
    for (unsigned char c : s) {
        if (c >= 'a' && c <= 'z') lowercase = true;
        else if (c >= 'A' && c <= 'Z') uppercase = true;
        else if (c >= '0' && c <= '9') digit = true;
        else special = true;
    }

    int categories = lowercase + uppercase + digit + special;
    if (categories == 4) cout << "Strong\n";
    else if (categories >= 3) cout << "Moderate\n";
    else cout << "Weak\n";

    return 0;
}