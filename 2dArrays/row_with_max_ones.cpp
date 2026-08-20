#include <bits/stdc++.h>
using namespace std;

int maxOnesRow(vector<vector <int>> mat) {
    int n = mat.size(), m = mat[0].size();
    int onesRow = -1, maxOnes = -1;

    for (int i = 0; i < n; i++) {
        int cntOnes = 0;
        for (int j = 0; j < m; j++) {
            if (mat[i][j] == 1) {
                cntOnes++;
            }
        }

        if (cntOnes > maxOnes) {
            maxOnes = cntOnes;
            onesRow = i;
        }
        else if (cntOnes == 0) {
            onesRow = -1;
        }
    }

    return onesRow;
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

    cout << maxOnesRow(mat) << endl;

    return 0;
}