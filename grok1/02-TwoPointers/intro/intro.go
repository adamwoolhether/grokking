package intro

/*
Given an array of sorted numbers and a target sum,
find a pair in the array whose sum is equal to the given target.

Because the array is sorted we can start with a pointer at the end
and a pointer at the beginning.

If the sum > target sum, decrement the end-pointer.
If the sum < target sum, increment the start-pointer.
*/
