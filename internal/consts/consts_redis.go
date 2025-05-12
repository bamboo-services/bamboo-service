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

	// RedisUserToken 用户登录令牌缓存用户数据的 Redis 键模板
	//
	//   - %s:  用户 UUID
	//   - %s:  令牌
	RedisUserToken = "user:token:%s:%s"

	// RedisUserEmail 邮箱缓存用户数据的 Redis 键模板
	//
	//   - %s:  邮箱
	RedisUserEmail = "user:email:%s"

	// RedisUserPhone 手机号缓存用户数据的 Redis 键模板
	//
	//   - %s:  手机号
	RedisUserPhone = "user:phone:%s"

	// RedisMailCodeSendTime 邮件验证码发送时间的 Redis 键模板。
	//
	//   - %s: MD5 加密的邮箱地址
	RedisMailCodeSendTime = "mail:code:send:%s"
)
