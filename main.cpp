#include<iostream>
#include<vector>

using namespace std;

void bubbleSort(vector<int>& ret) {
    bool flag = false;
    for (int i = ret.size() - 1; i > 0; i--) {
        flag = false;
        for (int j = 0; j < i; j++) {
            if (ret[j] > ret[j + 1]) {
                swap(ret[j], ret[j + 1]);
                flag = true;
            }
        }
        if (!flag) return;
    }
}

int partition(vector<int>& ret, int l, int r) {
    int randIndex  = rand() % (r - l) + l;
    swap(ret[randIndex], ret[l]);
    int pivotNum = ret[l];
    while (l < r) {
        while (l < r && ret[r] >= pivotNum) r--;
        ret[l] = ret[r];
        while (l < r && ret[l] <= pivotNum) l++;
        ret[r] = ret[l];
    }
    ret[l] = pivotNum;
    return l;
}

void quickSort(vector<int>& ret, int l, int r) {
    if (l < r) {
        int randIndex  = rand() % (r - l) + l;
        swap(ret[randIndex], ret[l]);
        int index = partition(ret, l , r);
        quickSort(ret, l, index - 1);
        quickSort(ret, index + 1, r);
    }
    return ;
}

int selectK(vector<int>& ret, int l, int r, int k) {
    while (l <= r) {
        int index = partition(ret, l , r);
        if (index == k - 1) return ret[index];
        else if (index < k - 1){
            l = index + 1;
        }else {
            r = index - 1;
        }
    }
    return -1;
}

int main() {
    vector<int> ret = {42,321,453,2,123,3,67,4,2,23,5,6};
    cout << selectK(ret, 0, ret.size() - 1, 8) << endl;

    quickSort(ret, 0, ret.size() - 1);
    for(int i = 0; i < ret.size(); i++ ) {
        cout << ret[i] << endl;
    }
    cout << selectK(ret, 0, ret.size() - 1, 8) << endl;
}
