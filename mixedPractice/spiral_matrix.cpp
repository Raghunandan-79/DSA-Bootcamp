#include <bits/stdc++.h>
using namespace std;

int main() {
    int m, n;
    cin >> m >> n;

    vector<vector<int>> matrix(m, vector<int>(n));
    for (auto& row : matrix) {
        for (int& value : row) {
            cin >> value;
        }
    }

    int top = 0, bottom = m - 1;
    int left = 0, right = n - 1;
    vector<int> spiral;

    while (top <= bottom && left <= right) {
        for (int col = left; col <= right; ++col) {
            spiral.push_back(matrix[top][col]);
        }
        ++top;

        for (int row = top; row <= bottom; ++row) {
            spiral.push_back(matrix[row][right]);
        }
        --right;

        if (top <= bottom) {
            for (int col = right; col >= left; --col) {
                spiral.push_back(matrix[bottom][col]);
            }
            --bottom;
        }

        if (left <= right) {
            for (int row = bottom; row >= top; --row) {
                spiral.push_back(matrix[row][left]);
            }
            ++left;
        }
    }

    for (size_t i = 0; i < spiral.size(); ++i) {
        if (i > 0) cout << ' ';
        cout << spiral[i];
    }
    cout << '\n';

    return 0;
}