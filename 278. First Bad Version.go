package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// binary search BUT take care of the beginning (left elem)
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left <= right {
		middle := (left + right) / 2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
func isBadVersion(n int) bool { return true }
