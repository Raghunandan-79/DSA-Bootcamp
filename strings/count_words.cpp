#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);

    int countWords = 1;
    for (int i = 0; i < s.size(); i++) {
        if (s[i] == ' ') {
            countWords++;
        }
    }
    cout << countWords << endl;

    return 0;
}