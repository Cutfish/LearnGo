#include<iostream>

using namespace std;

bool isLetter(char c) {
    if ((c <= 'Z' && c >= 'A') || (c <= 'z' && c >= 'a')) return true;
    return false;
}
bool isMatch(char c1, char c2) {
    if (c1 == c2 || c1 - 32 == c2 || c1 + 32 == c2) {
        return true;
    }else {
        return false;
    }
}
bool isPalindrome(string s) {
    int l = 0, r = s.size() - 1;
    while (l < r) {
        while(l < r && !isLetter(s[l])) l++; 
        while(l < r && !isLetter(s[r])) r--; 
        if (l < r) {
            if (isMatch(s[l], s[r])) {
                l++;
                r--;
                cout << "flag" << endl;
            }else {
                return false;
            }
        }
    }
    return true;
}


int main() {
    string str = "OP";
    if (isPalindrome(str)) cout << "true" << endl;
    return 0;
}