#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    // upper part including middle row
    for (int i = 1; i <= n; i++) {
        int lead = n - i;
        for (int s = 0; s < lead; s++)
            cout << ' ';
        if (i == 1) {
            cout << "* ";
        }
        else {
            cout << "*";
            int inner = 2 * i - 3;
            for (int s = 0; s < inner; s++)
                cout << ' ';
            cout << " * ";
        }
        cout << "\n";
    }

    // lower part
    for (int i = n - 1; i >= 1; i--) {
        int lead = n - i;
        for (int s = 0; s < lead; s++)
            cout << ' ';
        if (i == 1) {
            cout << "* ";
        }
        else {
            cout << "*";
            int inner = 2 * i - 3;
            for (int s = 0; s < inner; s++)
                cout << ' ';
            cout << " * ";
        }
        cout << "\n";
    }

    return 0;
}