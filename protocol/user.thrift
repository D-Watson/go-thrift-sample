namespace go user

// 定义用户角色的枚举类型
enum UserRole {
    ADMIN = 1,
    USER = 2,
    GUEST = 3
}

// 定义用户结构体
struct User {
    1: i32 id,              // 用户 ID
    2: string name,         // 用户名
    3: bool isActive,       // 用户是否活跃
    4: UserRole role,       // 用户角色
    5: list<string> tags     // 用户标签
}

// 定义用户服务
service UserService {
    User getUser(1: i32 userId),       // 获取用户信息
    void updateUser(1: User user),      // 更新用户信息
    list<User> getAllUsers(),           // 获取所有用户
}