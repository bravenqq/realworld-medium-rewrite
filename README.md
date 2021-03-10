# go-realworld-medium

一个实现用户创做文章和浏览文章的平台

## 为什么需要使用 DDD

1. 从业务理解和分析方面

- 能更好的与领域专家沟通理解他们的需求

  1.  传统的数据库设计不考虑真正应用场景总会设计出耦合度非常高的架构

```
   example
    实现用户登录，由两个步骤构成：校验用户输入的用户名、密码-----》生成token返回给客户端,在用户模块实现
    func Login(name,password string)(string,error){
        VerifyUser(name,password)
        id := GenerateToken()
        Save(id)
    }

    上面的实现登录的问题是将token和用户模块耦合在一起，考虑通过手机号登录，这种实现一定要全部重构

```

2. 从架构设计的合理性和可移植性

   1. 分离了业务层和基础服务（如何实现）

      ddd 将业务与基础设施（比如 JSON/XML 的，数据库存储，通信机制等）分离开。这样一来可以很容易用 mock 的方式来解耦模型和基础设施，从而更容易测试和修改，二来我们的领域模型也更独立，更精简，在适应新的需求时修改也会更容易。

3. 从开发的效率和分工

4. 从微服务的合理拆分

5. 对需求变化的适应性

- 使用 DDD 思考了需求的实现场景，能实现出真正有价值的需求

```
   example
    1. 在设计用户模块时总会把登录中token的生成放到用户模块实现，实际上token的生成只是调用用户模块校验一下用户信息是否正确，后面跟用户并没有关系
```

## 如何使用 DDD

1. 使用 DDD 分析和理解当前的商业逻辑，找出我们需要实现的需求和领域词条

2. 分析出领域实体和实体的关系

3. 分析具体的应用场景，设计用例

4. 运用测试驱动开发测试需求的使用，判断是否需求是否正确合理

## 如何使用 DDD 开始分析简单的需求

## 使用 DDD 对商业需求进行模块拆分
