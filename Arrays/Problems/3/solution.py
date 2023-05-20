# Sorting the array and returning the element at (K-1)th index.
# Time complexity - O(n log n)
# Space complexity - O(1)

def getKthSmallestElement(arr, k):
    arr.sort()
    return arr[k-1]


if __name__ == "__main__":
    arr = [4, 5, 10, 30, 78, 97, 8]
    k = 3
    kthSmallest = getKthSmallestElement(arr, k)
    print("Kth smallest element in the given unsorted array: ", kthSmallest)
    