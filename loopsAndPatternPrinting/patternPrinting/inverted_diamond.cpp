#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    for (int i = 1; i <= 2 * n - 1; i++) {
        int k = (i <= n) ? i : (2 * n - i);
        int stars = n - k + 1;
        int spaces = 2 * (k - 1) + 1;

        for (int j = 0; j < stars; j++) cout << '*';
        for (int j = 0; j < spaces; j++) cout << ' ';
        for (int j = 0; j < stars; j++) cout << '*';
        cout << '\n';
    }

    return 0;
}