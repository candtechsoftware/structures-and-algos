#include <iostream> 
#include <vector> 

using namespace std; 
// Need to comilei in std=c++11

bool isValid(vector<int> arr, vector<int> sequence){
    //Going with same approach as go
    int index = 0;

    for (int val : arr){
        if (index == sequence.size())
            break;

        if (sequence[index] == val)
            index++;
    }
    return index == sequence.size();

}

int main(){
    vector<int> arr1 = {5, 1, 22, 25, 6, -1, 8, 10};
    vector<int> seq1 = {1, 6, -1, 10};
    vector<int> arr2 = {5, 21, 2, 25, 3, -4, 8, 40};

    bool answer = isValid(arr1, seq1);
    bool answer_two = isValid(arr2, seq1);
    
    cout << answer << " " << answer_two << endl;

    return 0;

}
