package consts

const (
	// RedisUserUUID 用户 UUID 缓存用户数据的 Redis 键模板
	//
	//   - %s:  用户 UUID
	RedisUserUUID = "user:uuid:%s"

	// RedisUserUsername 用户名缓存用户数据的 Redis 键模板
	//
	//   - %s:  用户名
	RedisUserUsername = "user:username:%s"
)
