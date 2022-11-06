package bitwiseManipulation

/*
Statement:
Given an image represented by an (n×n) matrix containing 0s and 1s,
your task is first to flip the image horizontally, then invert it, and return the resultant image.

Horizontally flipping an image means that the mirror image of the matrix should be returned.
Flipping of [1,0,0] horizontally results in [0,0,1].

Inverting an image means that every 0 is replaced by 1, and every 1 is replaced by 0.
Inverting [0,1,1] results in [1,0,0].

Constraints:
- Image should be a square matrix.
- 1 <= n <= 20
- image[i][j] is either 0 or 1.

How to solve:
1. Flip each row of the matrix.
2. Invert the rows of the flipped image with XOR
*/

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		length := len(row)
		for i := 0; i < length/2; i++ {
			row[i], row[length-1-i] = row[length-1-i]^1, row[i]^1
		}

		if (length/2)%2 != 0 {
			row[length/2] ^= 1
		}
	}

	return image
}
