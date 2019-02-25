#include <iostream>
#include <vector>
#include <map>
#include <string> 
using namespace std;

vector<int> twoSum(vector<int>& nums, int target) {
    vector<int> res;
    int len = nums.size();
    map<int, int>  mp;
    for(int i = 0 ; i < len ; i++) {
        mp.insert(pair<int, int>(nums[i], i));
    }

    for(int i = 0 ; i < len ; i++) {
        int flag = target - nums[i];
        // cout<<mp[flag]<<" "<<flag<<" baikun"<<endl;
        if(mp.find(flag) != mp.end() && mp[flag] != i) {
            res.push_back(i);
            res.push_back(mp[flag]);
            break;
        }
    }
    return res;
}



int main() {
    int m, x, t;
    cin >> m;
    vector<int> nums;
    for(int i = 0 ; i < m ; i++) {
        cin>>x;
        nums.push_back(x);
    } 
    cin>>t; 
    vector<int> res = twoSum(nums, t);
    for(int i = 0 ; i < res.size() ; i++) {
        cout<<res[i]<<" ";
    }
}