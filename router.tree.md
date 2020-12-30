# Radix Tree

在计算机科学中，基数树，或称Patricia trie/tree，或crit bit tree，压缩前缀树，是一种更节省空间的Trie（前缀树）。对于基数树的每个节点，如果该节点是唯一的子树的话，就和父节点合并。

`radix tree`可以被认为是一棵简洁版的前缀树。拥有共同前缀的节点也共享同一个父节点。下面是一个GET方法对应的路由树的结构：

```text
Priority   Path             Handle
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>

```

`*<num>`是方法（handler）对应的指针，从根节点遍历到叶子节点我们就能得到完整的路由表，图中的示例实现了以下路由：

```text
GET("/", func1)
GET("/search/", func2)
GET("/support/", func3)
GET("/blog/", func4)
GET("/blog/:post/", func5)
GET("/about-us/", func6)
GET("/about-us/team/", func7)
GET("/contact/", func8)
```

`:post`是真实的post name的一个占位符（就是一个参数）。这里体现了radix tree相较于hash-map的一个优点，树结构允许我们的路径中存在动态的部分（参数）,因为我们匹配的是路由的模式而不是hash值

为了更具扩展性，每一层的节点按照priority排序，priority是节点的子节点（儿子节点，孙子节点等）注册的handler的数量，这样做有两个好处：

1. 被最多路径包含的节点会被最先评估。这样可以让尽量多的路由快速被定位。

2. 有点像成本补偿。最长的路径可以被最先评估，补偿体现在最长的路径需要花费更长的时间来定位，如果最长路径的节点能被优先评估（即每次拿子节点都命中），那么所花时间不一定比短路径的路由长。

## insertChild

```text
给定父结点 n， 给定路径 path

如果 n 存在子节点
    找到 path 的共同前缀
        
    如果前缀只是部分共同
        分裂当前节点， 按照共同前缀创建新中间节点作为父节点
        
    用减去公共前缀后的 path，创建新节点，插入 n
否则
    直接新增节点

（对于通配符 :xxxx或者*，会作为一个单独的节点出现）
```

## search

```text
给定结点 n， 给定路径 path

循环
    如果当前节点 n 的路径 等于 path
        返回结果
    否则
        如果当前节点有通配符子节点
            从 URL 取出匹配值
            
            取出后，如果 URL 已经结尾
                返回结果
            否则
                继续寻找
                取出通配符节点 （含有通配符的父节点只有一个子节点，即通配符节点）
                赋值该节点给 n，进入第二次循环
        否则
            从 indices 找到对应子节点
            赋值该节点给 n，进入第二次循环
```