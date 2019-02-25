#include <bits/stdc++.h>
using namespace std;


struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};


ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* p = l1, *q = l2;
    ListNode* head = new ListNode(-1);
    ListNode* cur = head;
    int carry = 0;
    while(p != NULL || q != NULL) {
        int x = (p != NULL) ? p->val : 0;
        int y = (q != NULL) ? q->val : 0;
        int sum = x + y + carry;
        carry = sum / 10;
        cur->next = new ListNode(sum % 10);
        cur = cur->next;
        if (p != NULL) p = p->next;
        if (q != NULL) q = q->next;
    }
    if (carry > 0) {
        cur->next = new ListNode(carry);
    }
    return head->next;
}

int main() {
	int n, m, t;
	cin>>n>>m;
	cin>>t;
	ListNode* l1 = new ListNode(t);
	ListNode* head = l1;
	for(int i = 1 ; i < n ; i++) {
		cin>>t;
		head->next = new ListNode(t);
		head = head->next;
	}
	cin>>t;
	ListNode* l2 = new ListNode(t);
	head = l2;
	for(int i = 1 ; i < m ; i++) {
		cin>>t;
		head->next = new ListNode(t);
		head = head->next;
	}

	addTwoNumbers(l1, l2);
}





