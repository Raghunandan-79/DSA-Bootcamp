#include <bits/stdc++.h>
using namespace std;

int main() {int n, m;
    cin >> n >> m;
    vector<vector<int>> grid(n, vector<int>(m));
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < m; ++j) {
            cin >> grid[i][j];
        }
    }

    vector<int> ans;

    for (int r = n - 1; r >= 0; --r) {
        if (grid[r][0] == -1) break;
        ans.push_back(grid[r][0]);
    }

    for (int c = 1; c < m; ++c) {
        if (grid[0][c] == -1) break;
        ans.push_back(grid[0][c]);
    }

    for (int r = 1; r < n; ++r) {
        if (grid[r][m - 1] == -1) break;
        ans.push_back(grid[r][m - 1]);
    }

    for (int c = m - 2; c >= 0; --c) {
        if (grid[n - 1][c] == -1) break;
        ans.push_back(grid[n - 1][c]);
    }

    for (int i = 0; i < (int)ans.size(); ++i) {
        if (i) cout << ' ';
        cout << ans[i];
    }
    cout << endl;

    return 0;
}