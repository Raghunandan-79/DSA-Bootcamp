#include <bits/stdc++.h>
using namespace std;

void sortArray(int arr[], int n) {
    int countZeros = 0, countOnes = 0, countTwos = 0;

    for (int i = 0; i < n; i++) {
        if (arr[i] == 0) {
            countZeros++;
        }
        else if (arr[i] == 1) {
            countOnes++;
        }
        else if (arr[i] == 2) {
            countTwos++;
        }
    }

    for (int i = 1; i <= countTwos; i++) {
        cout << "2 ";
    }

    for (int i = 1; i <= countOnes; i++) {
        cout << "1 ";
    }

    for (int i = 1; i <= countZeros; i++) {
        cout << "0 ";
    }
}

int main() {
    int n;
    cin >> n;
    int arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    sortArray(arr, n);

    return 0;
}