# In place sorting function using three pointers.
# Time complexity - O(n)
# Space complexity - O(1)

def sort(arr):
    l = 0
    m = 0
    h = len(arr)-1
    while (m<=h):
        if arr[m]==0:
            arr[m], arr[l] = arr[l], arr[m]
            m+=1
            l+=1
        elif arr[m]==1:
            m+=1
        else:
            arr[m], arr[h] = arr[h], arr[m]
            h-=1
    return arr
    

if __name__ == "__main__":
    arr = [1, 0, 2, 1, 1, 2, 0, 0]
    arrSorted = sort(arr)
    print("sorted array: ", arrSorted)
    