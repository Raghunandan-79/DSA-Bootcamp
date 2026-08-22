#include <bits/stdc++.h>
using namespace std;

int main() {
    int n;
    cin >> n;
    
    long long diagonalSum = 0;
    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < n; ++j) {
            int value;
            cin >> value;
            if (i == j) {
                diagonalSum += value;
            }
        }
    }

    cout << diagonalSum << '\n';
    return 0;
}