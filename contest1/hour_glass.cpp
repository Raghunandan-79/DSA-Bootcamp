#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    for (int i = 0; i < n; ++i) {
        int dots = n - i;
        for (int s = 0; s < i; ++s) cout << ' ';
        for (int d = 0; d < dots; ++d) {
            if (d) cout << ' ';
            cout << '.';
        }
        cout << '\n';
    }

    for (int i = n - 2; i >= 0; --i) {
        int dots = n - i;
        for (int s = 0; s < i; ++s) cout << ' ';
        for (int d = 0; d < dots; ++d) {
            if (d) cout << ' ';
            cout << '.';
        }
        if (i) cout << '\n';
    }

    return 0;
}