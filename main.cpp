#include<iostream>
#include<vector>

using namespace std;

int partition(vector<int>& nums, int l, int r) {
    int index = l + rand()%(r - l + 1);
    swap(nums[index], nums[l]);
    int pivot = nums[l];
    while (l < r) {
        while (l < r && nums[r] >= pivot) r--;
        nums[l] = nums[r];
        while (l < r && nums[l] <= pivot) l++;
        nums[r] = nums[l];
    }
    nums[l] = pivot;
    return l;
}

void quickSort(vector<int> &nums, int l, int r) {
    if (l <= r) {
        int index = partition(nums, l, r);
        quickSort(nums, l, index - 1);
        quickSort(nums, index + 1, r);
    }
}

int main() {
    vector<int> nums = {3,43,523,56,3,23,1,-342,4234};
    int len = nums.size();
    quickSort(nums, 0, len - 1);
    for (int i : nums) {
        cout << i << " ";
    }
    cout << endl;
    return 0;
}