# Manning.Algorithms.of.the.Intelligent.Web

notations:

    !  # important
    ?  # todo


## chapter 5: Classification: placing things where they belong
### ! 5.2 An overview of classifiers
![An overview of the classification algorithms based on their design](./assets/aiw_classifcation_overview.png)

#### 5.2.1 Structural classification algorithms

* Rule base: production rules(if-then clauses) & decision tree:  
    simle yet effective, not able to deal unseen data.

* distance-based algorithms: 
    >careful normalization and analysis of the attribute space is crucial to the success of distance-based algorithms.

    * Functional classifiers:  approximate the data by function, as the name suggests
        * linear regression & linear approximation
    * Nearest-neighbor algorithms attempt to find the nearest class for each data point. (using distance )

* Neural network (deep learning):
    > we don’t have a design methodology that would be applicable in a large number of problems

    这个问题 TensorFlow应该是解决掉了
    > , and it’s difficult to interpret the results of neural network classification

    这个应该会持续存在

#### 5.2.2 Statistical classification algorithms

* Regression algorithms: are usually employed when the data points are inherently numerical variables
    * logistic regression.

* Bayes rule or Bayes theorem:
    > The fascinating aspect of Bayesian algorithms is that they seem to work well even 
    when that independence assumption is clearly violated

#### 5.2.3 The lifecycle of a classifier
>There are three stages in the life- cycle of a classifier: training, testing, and production.

? using what metrics during validation phase to tell if the trained model is good enough 



# Links
* [! Machine Learning: Measuring Similarity and Distance](https://dzone.com/articles/machine-learning-measuring)


# todos:

    done:

    pending:

        linear regression & linear approximation & logistic regression.