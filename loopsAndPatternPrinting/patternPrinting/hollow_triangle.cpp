#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    for (int i = 1; i <= n; i++) {
        for (int s = 0; s < n - i; s++) cout << ' ';

        if (i == n) {
            for (int k = 0; k < n; k++) cout << "* ";
        } 
        else if (i == 1) {
            cout << "* ";
        } 
        else {
            cout << "*";
            int inner = 2 * i - 3;
            for (int s = 0; s < inner; s++) cout << ' ';
            cout << "* ";
        }
        cout << "\n";
    }

    return 0;
}