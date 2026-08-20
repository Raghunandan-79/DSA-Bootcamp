#include <bits/stdc++.h>
using namespace std;

int main() {
    char ch;
    cin >> ch;

    if (islower(ch)) {
        cout << "Lowercase" << endl;
    }
    else if (isupper(ch)) {
        cout << "Uppercase" << endl;
    }
    else if (isdigit(ch)) {
        cout << "Digit" << endl;
    }
    else {
        cout << "Special" << endl;
    }

    return 0;
}