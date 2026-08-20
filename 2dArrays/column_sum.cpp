#include <bits/stdc++.h>
using namespace std;

vector<int> columnSum(vector<vector<int>> mat) {
    int n = mat.size(), m = mat[0].size();
    vector<int> ans(m, 0);

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            ans[j] += mat[i][j];
        }
    }

    return ans;
}

int main() {
    int n, m;
    cin >> n >> m;
    vector<vector<int>> mat(n, vector<int>(m));

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> mat[i][j];
        }
    }

    for (auto num : columnSum(mat)) {
        cout << num << " ";
    }

    return 0;
}