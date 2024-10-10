import numpy as np
# Numpy - Numeric Python
l = [1,2,3,4]
# print(l)

arr1 = np.array([-2,-1,1,2,3,4,5,6,7,8,9,10])
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
# print(arr5)
l2 = ["I", "LOVE", "INDIA", 1947, True]
arr5 = np.array(l2)
# print(arr5)


# print(arr1[2:6])
# print(arr1[-5:8])
# print(arr1[5:-8])


# arr6 = np.array([[10,11,12,13,14,15,16], [0,1,2,3,4,5,6,7,8,9]], dtype=object)
arr6 = np.array([[10,11,12,13,14,15,16], [0,1,2,3,4,5,6]])
# print(arr6[1,5])  # print(arr6[1][5])
# print(arr6[0:1, 2:5])



#####################################################

# # Square root of all elem
# print(np.sqrt(arr6))

# # Absolute value
# print(np.absolute(arr1))

# # Exponentials
# print(np.exp(arr1))

# # Max Min
# print(np.max(arr6))
# print(np.min(arr6))

# # Negative or positive
# print(np.sign(arr1))

# # Trigonometry
# print(np.sin(arr1))
# print(np.cos(arr1))
# print(np.tan(arr1))

# # Log
# print(np.log(arr1))



########################################################

# view1 = arr1.view()
# print("Original: ", arr1, "View: ", view1)
# arr1[0] = 100
# print("Original: ", arr1, "View: ", view1)
# view1[0] = -200
# print("Original: ", arr1, "View: ", view1)
# print(type(view1), id(view1), id(arr1))


# copy1 = arr1.copy()
# print("Original: ", arr1, "Copy: ", copy1)
# arr1[0] = 100
# print("Original: ", arr1, "Copy: ", copy1)
# copy1[0] = -200
# print("Original: ", arr1, "Copy: ", copy1)
# print(type(copy1), id(copy1), id(arr1))

#######################################################

''' 
Shape  Reshape
'''

# print(arr1.shape)
# print(arr6.shape)

# copy1 = arr1.reshape(4,3)
# print(copy1, copy1.shape)

# copy2 = arr1.reshape(3,2,2)
# print(copy2, copy2.shape)

# copy3 = copy1.reshape(-1) # Reshape to 1D
# print(copy3, copy3.shape)

#######################################################

# copy2 = arr1.reshape(3,2,2)
# print(copy2)

# # Using np.nditer()

# for i in np.nditer(copy2):
#     print(i)


#######################################################

# arr7 = np.array([5,8,6,4,2,1,1])
# print(arr7)
# print(np.sort(arr7))

# Original is not changed
# A copy is returned

# print(arr5)
# print(np.sort(arr5))

# arr7 = np.array([[8,6,1,3,5], [3,1,9,5,6]])
# print(arr7)
# print(np.sort(arr7))


######################################################

# Searching

# x = np.where(arr1 == 6)
# print(arr1, x)
# print(x[0])
# print(arr1[x[0]])


# x = np.where(arr1 > 0)
# print(arr1, x[0])

# x = np.where(arr1 % 3 == 0)
# print(arr1, x[0])
# print(set(arr1))


######################################################

# Filtering with Boolean Index Lists

arr8 = np.array([1, 2, 3, 4, 5, 6])
x = [True, False] * 3
print(arr8)
print(arr8[x])

filter = arr8 %2 != 0
print(arr8)
print(filter)
print(arr8[filter])