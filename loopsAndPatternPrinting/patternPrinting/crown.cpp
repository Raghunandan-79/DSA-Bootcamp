#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    for (int i = 1; i <= n; i++){
        if (i == n) {
            for (int k = 0; k < 2 * n; k++) cout << "*";
            cout << "\n";
            break;
        }

        for (int k = 0; k < i; k++) cout << "*";
        int spaces = 2 * (n - i);
        for (int s = 0; s < spaces; s++) cout << " ";
        for (int k = 0; k < i; k++) cout << "*";
        cout << "\n";
    }

    return 0;
}