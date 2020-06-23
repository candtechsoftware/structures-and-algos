// Pretty much same approach as others but using a while loop. This one to me 
// isn't as intuitive as the other solutions (check go code).
//

let isValiSubsequence = (arr, seq) => {
    let index = 0;
    let seqIndex = 0;

    while(index < arr.length && seqIndex < seq.length) {
        if (arr[index] === seq[seqIndex]) seqIndex++;
        index++;
    }

    return seqIndex === seq.length; 

}


console.log(isValiSubsequence([5, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1, 10]))
