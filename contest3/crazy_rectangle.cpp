#include <bits/stdc++.h>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;

    int number = 1;
    char letter = 'a';
    bool useNumber = true;

    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            if (useNumber) {
                cout << number;
                number = number % 9 + 1;
            } 
            else {
                cout << letter;
                letter = letter == 'z' ? 'a' : letter + 1;
            }
            useNumber = !useNumber;
        }
        cout << '\n';
    }

    return 0;
}