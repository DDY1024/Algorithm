package bloom

// 分布式 bloom filter 实现
// 1. redis setbit/getbit 命令
// 2. lua 脚本实现原子操作
// 具体实现可以参考 go-zero: https://github.com/zeromicro/go-zero/blob/master/core/bloom/bloom.go

const (
	// 1. 设置脚本
	setScript = `
for _, offset in ipairs(ARGV) do
	redis.call("setbit", KEYS[1], offset, 1)
end
`
	// 2. 检测脚本
	testScript = `
for _, offset in ipairs(ARGV) do
	if tonumber(redis.call("getbit", KEYS[1], offset)) == 0 then
		return false
	end
end
return true
`
)

// Redis 相关操作
// 1. 设置操作
// redis.Eval(setScript, []string{key}, args)
//
// 2. 检测操作
// redis.Eval(testScript, []string{key}, args)
//
// 注意事项
// 1. redis 大 key 限制，单个 bloom 槽位上限为 10 * 1024 * 8 bit（10KB）
// 2. 如果需要更大的 bloom filter 支持，在 1 基础上，进一步进行分片
//

// pipeline 版本的 bloom filter
// 1. Set 操作
/*
func (r *RedisBitSet) Set(offsets []uint) error {
	pipe := NewPipeline("bloom_set")
	defer pipe.Close()

	for _, offset := range offsets {
		key, thisOffset := r.getKeyOffset(offset)
		err := pipe.SetBit(key, int64(thisOffset), 1).Err()  // pipe.SetBit 设置 bit 位操作
		if err != nil {
			_ = pipe.Close()
			return err
		}
	}
	_, err := pipe.Exec()
	return err
}
*/

// 2. Test 操作
/*
func (r *RedisBitSet) Test(offsets []uint) (bool, error) {
	pipe := r.conn.NewPipeline("bloom_test")
	defer pipe.Close()

	for _, offset := range offsets {
		key, thisOffset := r.getKeyOffset(offset)
		pipe.GetBit(key, int64(thisOffset))  // pipe.GetBit 读取 bit 位操作
	}
	checks, err := pipe.Exec()
	if err != nil {
		return false, err
	}

	for _, value := range checks {
		bitValue, err := value.(*redis.IntCmd).Result()
		if err != nil {
			return false, err
		}
		if bitValue == 0 {
			return false, nil
		}
	}
	return true, nil
}
*/
