class pair:
    def __init__(self):
        self.min = 0
        self.max = 0

def findMinMax(arr) -> pair:
    minMax = pair()
    n = len(arr)
    if n==1:
        minMax.min = arr[0]
        minMax.max = arr[0]
        return minMax
    if arr[0] > arr[1]: 
        minMax.max = arr[0]
        minMax.min = arr[1]
    else:
        minMax.max = arr[1]
        minMax.min = arr[0]
    for i in range(2, n):
        if arr[i]>minMax.max:
            minMax.max = arr[i]
        elif arr[i]<minMax.min:
            minMax.min = arr[i]
    return minMax
	    


if __name__ == "__main__":
    arr = [2, 1, 5, 4, 10, 9, 6]
    minMax = findMinMax(arr)
    print("minimum element: ", minMax.min)
    print("maximum element: ", minMax.max)