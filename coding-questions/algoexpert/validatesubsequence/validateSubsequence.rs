
fn is_valid_subsequence(arr: Vec<i32>, seq: Vec<i32>) -> bool {
    let mut index = 0;
    
    for val in arr.iter() {
        if index == arr.len() {
            break; 
        }
        if seq[index] == *val {
            index+=1; 
        }
    }
    return index == seq.len();
}

fn main() {
    let arr = vec![5, 1, 22, 25, 6, -1, 8, 10];
    let seq = vec![1, 6, -1, 10];

    let ans = is_valid_subsequence(arr, seq); 
    println!("{}", ans);
}
