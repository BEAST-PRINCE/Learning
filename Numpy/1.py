import numpy as np
# Numpy - Numeric Python
l = [1,2,3,4]
print(l)

arr1 = np.array([1,2,3,4,5,6,7,8,9,10])
# print(arr1)
# print(arr1.shape)

arr2 = np.arange(arr1.shape[0])
# print(arr2)

arr2 = np.arange(0, arr1.shape[0], 3)
# print(arr2)


zero_arr = np.zeros(5)
# print(zero_arr)

# Multidimensional arrays
arr3 = np.zeros((3, 5)) # dimensions, no. of elements
# print(arr3)


# numpy.full(shape, fill_value, dtype=None, order='C', *, like=None)
arr4 = np.full((3, 5), 5, dtype=int)
# print(arr4)


arr5 = np.array(l)
print(arr5)
l2 = [1, "I", "LOVE", "INDIA", 1947, True]
arr5 = np.array(l2)
print(arr5)