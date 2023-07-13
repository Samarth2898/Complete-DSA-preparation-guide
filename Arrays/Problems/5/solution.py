# In place sorting function using two pointers.
# Time complexity - O(n)
# Space complexity - O(1)

def sortArr(arr):
    s = 0
    e = 0
    n = len(arr) - 1
    while e <= n:
        if arr[e] < 0:
            arr[s], arr[e] = arr[e], arr[s]
            e += 1
            s += 1
        else:
            e += 1
    return arr

if __name__ == '__main__':
    arr = [-1, -3, 2, -7, 4, 5, -5, -9, 10]
    sortedArr = sortArr(arr)
    print("sorted array: ", sortedArr)
