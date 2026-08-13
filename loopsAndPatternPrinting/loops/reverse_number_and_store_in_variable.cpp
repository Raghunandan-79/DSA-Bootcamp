#include <bits/stdc++.h>
using namespace std;

int main() {
    long long n;
    cin >> n;

    long long ans = 0;

    while (n > 0) {
        long long current_digit = n % 10;
        ans = (ans * 10) + current_digit;
        n /= 10;
    } 
    
    cout << ans;

    return 0;
}