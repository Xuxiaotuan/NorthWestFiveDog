#include <bits/stdc++.h>
using namespace std;
int ispan(int x, int y) {
	if (x >= 0 && x < 8 && y >= 0 && y < 8) 
		return 1;
	return 0;
}
int numRookCaptures(vector<vector<char>>& board) {
	int flagx, flagy;
	for (int i = 0 ; i < 8 ; i++) {
		for (int j = 0 ; j < 8 ; j++) {
			if (board[i][j] == 'R') {
				flagx = i;
				flagy = j;
			}
		}
	} 
	int dx[] = {0, 1, -1, 0};
	int dy[] = {1, 0, 0, -1};
	int ans = 0;
	for(int i = 0 ; i < 4 ; i++) {
		int x = flagx + dx[i];
		int y = flagy + dy[i];
		while (ispan(x, y)) {
			if (board[x][y] == 'p') {
				ans++;
				break;
			} else if (board[x][y] == 'B') {
				break;
			} else {
				x = x + dx[i];
				y = y + dy[i];
			}
		}
	}
	return ans;   
}
int main() {
	char t;
	vector<vector<char>> board;
	for (int i = 0 ; i < 8 ; i++) {
		vector<char> tmp;
		for (int j = 0 ; j < 8 ; j++) {
			cin>>t;
			tmp.push_back(t);
		}
		board.push_back(tmp);
	}
	cout<<numRookCaptures(board)<<endl;
}