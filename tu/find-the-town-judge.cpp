#include <iostream>
#include <vector>
#include <map>
#include <string> 
#include <vector>
using namespace std;
int findJudge(int N, vector<vector<int>>& trust) {
	int res[N + 2][2];
	memset(res, 0, sizeof(res));
	int sz = trust.size();
	for(int i = 0 ; i < sz ; i++) {
		res[trust[i][1]][0]++;
		res[trust[i][0]][1]++;
	}
	int ans = -1;
	for(int i = 1 ; i <= N ; i++) {
		if(res[i][0] == (N - 1) && res[i][1] == 0) {
			ans = i;
			break;
		}
	}
	return ans;
}
int main() {
	
}