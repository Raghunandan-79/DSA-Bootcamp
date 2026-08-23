#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;
    long long totalSum = 0, diagonalSum = 0;
    
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            long long value;
            cin >> value;
            totalSum += value;

            if (i == j || i + j == n - 1) {
                diagonalSum += value;
            }
        }
    }

    cout << diagonalSum << " " << totalSum - diagonalSum << endl;

    return 0;
}