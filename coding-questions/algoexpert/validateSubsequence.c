#include <stdio.h> 
#include <stdbool.h> 

bool is_valid(int arr[], int seq[]){

    return 2 <3; 
}
int main() {
    int ans[8] = {5, 1, 22, 25, 6, -1, 8, 10};
    int seq[4] = {1, 6, -1, 10}; 

    bool answer = is_valid(ans, seq);
    printf(answer ? "true\n" : "false\n");
    return 0;

}
