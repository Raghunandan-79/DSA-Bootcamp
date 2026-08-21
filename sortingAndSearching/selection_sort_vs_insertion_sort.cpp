#include <bits/stdc++.h>
using namespace std;

long long countInsertionShifts(vector<int> a) {
    long long shifts = 0;
    int n = a.size();

    for (int i = 1; i < n; ++i) {
        int key = a[i];
        int j = i - 1;

        while (j >= 0 && a[j] > key) {
            a[j + 1] = a[j];
            ++shifts;
            --j;
        }
        a[j + 1] = key;
    }

    return shifts;
}

long long countSelectionSwaps(vector<int> a) {
    long long swaps = 0;
    int n = a.size();

    for (int i = 0; i < n - 1; ++i) {
        int minIndex = i;

        for (int j = i + 1; j < n; ++j) {
            if (a[j] < a[minIndex]) {
                minIndex = j;
            }
        }

        if (minIndex != i) {
            swap(a[i], a[minIndex]);
            ++swaps;
        }
    }

    return swaps;
}

int main() {
    int t;
    cin >> t;

    while (t--) {
        int n;
        cin >> n;

        vector<int> a(n);
        for (int i = 0; i < n; ++i) cin >> a[i];

        long long insertionShifts = countInsertionShifts(a);

        vector<int> b = a;
        long long selectionSwaps = countSelectionSwaps(b);

        if (insertionShifts < selectionSwaps) {
            cout << "Insertion Sort\n";
        } 
        else if (selectionSwaps < insertionShifts) {
            cout << "Selection Sort\n";
        } 
        else {
            cout << "Tie\n";
        }
    }

    return 0;
}