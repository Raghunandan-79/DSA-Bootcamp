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

    int maximum = INT_MIN;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (mat[i][j] > maximum) {
                maximum = mat[i][j];
            }
        }
    }

    cout << maximum << endl;

    return 0;
}