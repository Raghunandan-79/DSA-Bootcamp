#include <bits/stdc++.h>
using namespace std;

void countPassedAndFailed(int arr[], int n, int marks) {
    int pass = 0, fail = 0;

    for (int i = 0; i < n; i++) {
        if (arr[i] >= marks) {
            pass++;
        }
        else {
            fail++;
        }
    }

    cout << "Pass: " << pass << endl;
    cout << "Fail: " << fail << endl;
}

int main() {
    int n;
    cin >> n;
    int arr[n];

    for (int i = 0; i < n; i++) {
        cin >> arr[i];
    }

    int marks;
    cin >> marks;

    countPassedAndFailed(arr, n, marks);

    return 0;
}