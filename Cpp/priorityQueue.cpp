#include <iostream>
#include <ostream>
#include <queue>
#include <vector>

using namespace std;

struct node {
    int data;
    int test;
    bool operator<(const node &b) const { return data < b.data; }
    node(int d) { data = d; }
};

struct cmp {
    bool operator()(node *a, node *b) { return *a < *b; }
};

int main() {
    priority_queue<node *, vector<node *>, cmp> data;

    node *new_node;

    int input, count;
    cin >> count;

    for (int i = 0; i < count; i++) {
        cin >> input;
        new_node = new node(input);
        data.push(new_node);
    }

    node *input1;
    node *input2;

    while (data.size() > 1) {
        input1 = data.top();
        data.pop();
        input2 = data.top();
        data.pop();
        input1->data = input1->data + input2->data;
        delete input2;
        data.push(input1);
    }

    /* count << data.top()->data << endl; */
    new_node = data.top();
    data.pop();

    delete new_node;
    return 0;
}
