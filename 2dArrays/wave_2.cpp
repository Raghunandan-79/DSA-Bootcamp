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
        if (j % 2 == 0) {
            for (int i = 0; i < n; i++) {
                cout << mat[i][j] << ' ';
            }
        } 
        else {
            for (int i = n - 1; i >= 0; i--) {
                cout << mat[i][j] << ' ';
            }
        }
    }    

    return 0;
}