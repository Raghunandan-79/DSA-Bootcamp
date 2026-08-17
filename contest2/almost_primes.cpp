#include <bits/stdc++.h>
using namespace std;

int countPositiveDivisors(int n) {
    int count = 0;

    for (int i = 1; i <= n; i++) {
        if (n % i == 0) {
            count++;
        }
    }

    return count;
}

void printAlmostPrimes(int n) {
    for (int i = 1; i <= n; i++) {
        if (countPositiveDivisors(i) <= 4) {
            cout << i << " ";
        }
    }
}

int main() {
    int n;
    cin >> n;

    printAlmostPrimes(n);

    return 0;
}