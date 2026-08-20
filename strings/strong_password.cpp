#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    cin >> s;

    int lengthOfString = s.size();
    int lowercaseCount = 0;
    int uppercaseCount = 0;
    int digitCount = 0;
    int specialCount = 0;

    for (int i = 0; i < lengthOfString; i++) {
        if (islower(s[i])) {
            lowercaseCount++;
        }
        else if (isupper(s[i])) {
            uppercaseCount++;
        }
        else if (isdigit(s[i])) {
            digitCount++;
        }
        else {
            specialCount++;
        }
    }

    if (lengthOfString == 10 && lowercaseCount >= 1 && uppercaseCount >= 1 && digitCount >= 1 && specialCount >= 1) {
        cout << "Strong" << endl;
    }
    else {
        cout << "Weak" << endl;
    }

    return 0;
}