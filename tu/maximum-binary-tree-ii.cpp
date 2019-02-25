#include <bits/stdc++.h>
using namespace std;
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

TreeNode * link(TreeNode *pre, TreeNode *root, TreeNode *node) {
    if(root == NULL || node->val > root->val) {
        node->left = root;
        if(pre) {
            pre->right = node;
            return NULL;
        }
        return node;
    }
    link(root, root->right, node);
    return root;
}
TreeNode* insertIntoMaxTree(TreeNode* root, int val) {
    TreeNode *node = new TreeNode(val);
    return link(NULL, root, node);
}

int main() {

}