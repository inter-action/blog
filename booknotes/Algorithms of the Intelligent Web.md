# Book: Manning.Algorithms.of.the.Intelligent.Web.2nd.Edition.2016.8
由于看过第一版(只读了几章), 书的第一版干货很多。所以对第二版自然很有兴趣。书的一二版区别还是
挺大的，

第一版:
* 代码用java写的
* 代码很多，好多算法都是自己用java实现的, 能从源码中了解好多算法底层的知识
* 内容相对较旧, 毕竟09年出的, 但是好多算法还是很有价值
* 没有直观的可运行图表,来直观的查看数据

第二版：
* 代码python
* 代码很少，主要是讲解。利用了很多高级的工具库, eg. scikit-learn 来演示对应的实现
* 内容较新
* 有图表可以直观查看数据


书中有好多统计学相关的知识, 如果你不懂的话会影响你的理解。我的建议就是可以试着先
读下*OReilly.Think.Stats.2nd.Edition.2014.10*一书。这也是我自己的打算。
当初没有下毅力读完。

## chapter 2: Extracting structure from data: clustering and transforming your data
 
### 2.3: K-means
* desc:  
  用图解的方式介绍了k-means的算法实现, 具体的代码是实现第一版书有源码。

* k-means的工作方式:   
  kmeans基本的方式就是你预估要cluster数据的个数, 然后提供对应个数的centroid(中心点)
  的数据, 然后算法会自动在迭代中优化cluster中心点。直到converge(数据不再变化)算法停止。


### 2.4: The Gaussian mixture model

* desc: 
书先讲了什么是正太分布(Gaussian model), 用欧洲人的身高。
然后又讲了什么是Gaussian mixture model, 用男性和女性的两种正太分布组合在一张图上。
然后就直接应用scikit-learn的load_iris()数据作图了(尴尬脸, 还没有看懂呢好吗？)。
示例代码是3个cluster, 4个维度的向量。而男女的是2个cluster(男或者女)的1个维度向量(身高)。
而下文说的Gaussian mixture model不仅计算距离的相似性，还计算不同cluster的shape。
我只能理解是正太分布的mean和variation了, :(.  
至于这个算法怎么工作的, 我的猜测(根据作者举的男女身高的正太分布的例子)就是通过训练数据找出
不同cluster的不同维度的正太分布的mean和variation值。获得正太分布的数学定义, 然后带入
需要分类的数据, 计算在男性和女性不同的概率的值, 比较取概率最大的那个去分是属于哪个cluster.
在feature多维度下, 可能会综合各个维度的概率值产生一个最终的结果。具体公式是啥就不知道啦~

* terms:
  *  Gaussian distribution - akka (normal distribution) 正太分布
  * PDF(probability density function) - p(x), 给定x值返回对应的概率
  * CDF(Cumulative distribution function) - p(x1 <= x), 给定x值返回所有x1 <= x值得概率总和

### 2.5 The relationship between k-means and GMM
>K-means can be expressed as a special case of the Gaussian mixture model. In general, the Gaussian mixture is more expressive because membership of a data item to a clus- ter is dependent on the shape of that cluster, not just its proximity.

什么时候GMM变成k-means
>In particular, if covariance matrices for each cluster are tied together (that is, they all must be the same), and covariances across the diagonal are restricted to being equal, with all other entries set to zero, then you’ll obtain spherical clusters of the same size and shape for each cluster. Under such circumstances, each point will always belong to the cluster with the closest mean. Try it and see!

todo: 这里说的 `covariance matrices` 是什么

