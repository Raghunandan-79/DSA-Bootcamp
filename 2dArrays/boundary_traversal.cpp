#include <bits/stdc++.h>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;
    vector<vector<int>> mat(n, vector<int>(m));

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> mat[i][j];
        }
    }

    for (int j = 0; j < m; j++) {
        cout << mat[0][j] << ' ';
    }

    for (int i = 1; i < n; i++) {
        cout << mat[i][m - 1] << ' ';
    }

    if (n > 1) {
        for (int j = m - 2; j >= 0; j--) {
            cout << mat[n - 1][j] << ' ';
        }
    }

    if (m > 1) {
        for (int i = n - 2; i > 0; i--) {
            cout << mat[i][0] << ' ';
        }
    }

    return 0;
}