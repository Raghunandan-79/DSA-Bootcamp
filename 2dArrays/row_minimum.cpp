#include <bits/stdc++.h>
using namespace std;

vector<int> rowMinimums(vector<vector<int>> mat) {
    int n = mat.size(), m = mat[0].size();
    vector<int> ans;

    for (int i = 0; i < n; i++) {
        int minimum = INT_MAX;

        for (int j = 0; j < m; j++) {
            if (mat[i][j] < minimum) {
                minimum = mat[i][j];
            }
        }

        ans.push_back(minimum);
    }

    return ans;
}

int main() {
    int n, m;
    cin >> n >> m;
    vector<vector<int>> mat;

    for (int i = 0; i < n; i++) {
        vector<int> row;
        for (int i = 0; i < m; i++) {
            int num;
            cin >> num;
            row.push_back(num);
        }
        mat.push_back(row);
    }

    for (auto num : rowMinimums(mat)) {
        cout << num << " ";
    }

    return 0;
}