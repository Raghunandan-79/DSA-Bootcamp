#include <bits/stdc++.h>
using namespace std;

int main() {
    string s;
    getline(cin, s);

    int longest = 0;
    int current = 0;

    for (char c : s) {
        if (c == ' ') {
            longest = max(longest, current);
            current = 0;
        } 
        else {
            current++;
        }
    }

    longest = max(longest, current);
    cout << longest << '\n';

    return 0;
}