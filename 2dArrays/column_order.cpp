#include <bits/stdc++.h>
using namespace std;

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

    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            cout << mat[j][i] << " ";
        }
    }

    return 0;
}