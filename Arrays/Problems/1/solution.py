# iterative approach
# time complexity O(n)
# space complexity O(1)
def reverseList(arr, start, end):
    while start < end:
        arr[start], arr[end] = arr[end], arr[start]
        start += 1
        end -= 1

# recursive approach
# time complexity O(n)
# space complexity O(1)
def reverseListRecursive(arr, start, end):
    if start >= end:
        return
    arr[start], arr[end] = arr[end], arr[start]
    reverseListRecursive(arr, start+1, end-1)
 
if __name__ == "__main__":
    arr = [1, 2, 3, 4]
    print("original array: ", arr)
    reverseList(arr, 0, len(arr)-1)
    print("reversed array: ", arr)

    arr2 = [5, 6, 7, 8]
    print("original array: ", arr2)
    reverseList(arr2, 0, len(arr2)-1)
    print("reversed array: ", arr2)
