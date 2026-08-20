#include <bits/stdc++.h>
using namespace std;

bool isPresent(vector<vector<int>> mat, int x) {
    int n = mat.size(), m = mat[0].size();

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            if (mat[i][j] == x) {
                return true;
            }
        }
    }

    return false;
}

int main() {
    int n, m, x;
    cin >> n >> m >> x;
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

    if (isPresent(mat, x)) {
        cout << "true" << endl;
    }
    else {
        cout << "false" << endl;
    }

    return 0;
}