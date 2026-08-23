#include <bits/stdc++.h>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;

    vector<vector<int>> grid(n, vector<int>(m));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> grid[i][j];
        }
    }

    vector<int> boundary;
    for (int r = n - 1; r >= 0; r--) boundary.push_back(grid[r][0]);
    for (int c = 1; c < m; c++) boundary.push_back(grid[0][c]);
    for (int r = 1; r < n; r++) boundary.push_back(grid[r][m - 1]);
    for (int c = m - 2; c >= 1; c--) boundary.push_back(grid[n - 1][c]);

    bool first = true;
    for (int x : boundary) {
        if (x == -1) break;

        if (!first) cout << ' ';
        cout << x;
        first = false;
    }

    cout << '\n';
    return 0;
}