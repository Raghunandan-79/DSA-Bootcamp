#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;

    vector<vector<long long>> matrix(n, vector<long long>(n));
    for (auto& row : matrix) {
        for (long long& value : row) {
            cin >> value;
        }
    }

    for (int column = 0; column < n; ++column) {
        for (int row = n - 1; row >= 0; --row) {
            cout << matrix[row][column] << (row == 0 ? '\n' : ' ');
        }
    }

    return 0;
}