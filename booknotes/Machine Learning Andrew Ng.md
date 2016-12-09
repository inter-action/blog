# video: [Stanford: Machine Learning by Andrew Ng](https://www.youtube.com/playlist?list=PLJ1-ciQ35nuiyL1PX6O4NdF5CjjaDdnVC)

## Blah
this is not a booknote, but a video notes, but whatever.

## notations

under link context:

    ?              # todos



## chapter 2:

###2.6 Gradient Descent Intuition:
Gradient Descent Properties:
* if learning rate is too high, it may not converge even deverge
* if learning rate is too low, it just may take many steps(or iterations) to finally get converged
* if learning rate if fixed, the descent rate still get smaller with each iteration towards bottom (local optimal)
* most the case: theta0 or theta1 often initialized with 0

###2.7 Gradient Descent For Linear Regression
includes:
* linear regression's cost functional always be a convex function (or a bowl shaped function), which has only one local optimal value.
* what is `Batch Gradient Descent` a type of descent use all training example to iterate.

###2.8 What's Next

## Chapter 3:

###3.3 Matrix Vector Multiplication, 3.4 Matrix Matrix Multiplication
includes:
* ! a matrix application, right half of the video: `prediction = data_matrix * parameters`
### 3.5 Matrix Multiplication Properties
* associative
* not commutative

### 3.6 Inverse and Transpose
* the production of matrix with its Inverse matrix is equal to Identity matrix.
* matrices that dont have inverse are "singular" or "degenerate"

### 3.7 Gradient Descent in Practice I Feature Scaling
* scale all your feature to a reasonable range, roughly :
    * -1 < x < 1
    * mean normalization: (x-mean)/range(x)

### 4.4 Gradient Descent in Practice II Learning Rate
* plot cost function J(theta) as y, iteration count as x, to see whether gradient descent is working or not

### ! 4.5 Features and Polynomial Regression
* look you model see if features can merge
* choose right polynomial function, quantic or cubic 

### 4.6 Normal Equation
* normal equation: use mathematical equation to directly solve thetas, works better with small features
    and you dont need to choose learning rate

### 4.7 Normail Equation Noninvertibility
* when normal equation could break


# Links
* [List of mathematical symbols](https://en.wikipedia.org/wiki/List_of_mathematical_symbols)
* [Derivative](https://en.wikipedia.org/wiki/Derivative): notice the `partial derivative` section & its notation.
