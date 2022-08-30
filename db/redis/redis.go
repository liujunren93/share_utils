package redis

import (
	"context"
	"time"
	re "github.com/go-redis/redis/v8"
)

type Cmdable interface {
	Pipeline() re.Pipeliner
	Pipelined(ctx context.Context, fn func(re.Pipeliner) error) ([]re.Cmder, error)

	TxPipelined(ctx context.Context, fn func(re.Pipeliner) error) ([]re.Cmder, error)
	TxPipeline() re.Pipeliner

	Command(ctx context.Context) *re.CommandsInfoCmd
	ClientGetName(ctx context.Context) *re.StringCmd
	Echo(ctx context.Context, message interface{}) *re.StringCmd
	Ping(ctx context.Context) *re.StatusCmd
	Quit(ctx context.Context) *re.StatusCmd
	Del(ctx context.Context, keys ...string) *re.IntCmd
	Unlink(ctx context.Context, keys ...string) *re.IntCmd
	Dump(ctx context.Context, key string) *re.StringCmd
	Exists(ctx context.Context, keys ...string) *re.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *re.BoolCmd
	ExpireNX(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	ExpireXX(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	ExpireGT(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	ExpireLT(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	Keys(ctx context.Context, pattern string) *re.StringSliceCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *re.StatusCmd
	Move(ctx context.Context, key string, db int) *re.BoolCmd
	ObjectRefCount(ctx context.Context, key string) *re.IntCmd
	ObjectEncoding(ctx context.Context, key string) *re.StringCmd
	ObjectIdleTime(ctx context.Context, key string) *re.DurationCmd
	Persist(ctx context.Context, key string) *re.BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *re.BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *re.BoolCmd
	PTTL(ctx context.Context, key string) *re.DurationCmd
	RandomKey(ctx context.Context) *re.StringCmd
	Rename(ctx context.Context, key, newkey string) *re.StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *re.BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *re.StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *re.StatusCmd
	Sort(ctx context.Context, key string, sort *re.Sort) *re.StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort *re.Sort) *re.IntCmd
	SortInterfaces(ctx context.Context, key string, sort *re.Sort) *re.SliceCmd
	Touch(ctx context.Context, keys ...string) *re.IntCmd
	TTL(ctx context.Context, key string) *re.DurationCmd
	Type(ctx context.Context, key string) *re.StatusCmd
	// Append(ctx context.Context, key, value string) *re.IntCmd
	Decr(ctx context.Context, key string) *re.IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *re.IntCmd
	Get(ctx context.Context, key string) *re.StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *re.StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *re.StringCmd
	GetEx(ctx context.Context, key string, expiration time.Duration) *re.StringCmd
	GetDel(ctx context.Context, key string) *re.StringCmd
	Incr(ctx context.Context, key string) *re.IntCmd
	IncrBy(ctx context.Context, key string, value int64) *re.IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *re.FloatCmd
	MGet(ctx context.Context, keys ...string) *re.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *re.StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *re.BoolCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *re.StatusCmd
	SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *re.StatusCmd
	// TODO: rename to SetEx
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *re.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *re.BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *re.BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *re.IntCmd
	StrLen(ctx context.Context, key string) *re.IntCmd
	Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *re.IntCmd

	GetBit(ctx context.Context, key string, offset int64) *re.IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *re.IntCmd
	BitCount(ctx context.Context, key string, bitCount *BitCount) *re.IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *re.IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *re.IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *re.IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *re.IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *re.IntCmd
	BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd

	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd

	HDel(ctx context.Context, key string, fields ...string) *re.IntCmd
	HExists(ctx context.Context, key, field string) *re.BoolCmd
	HGet(ctx context.Context, key, field string) *re.StringCmd
	HGetAll(ctx context.Context, key string) *re.StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *re.IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *re.FloatCmd
	HKeys(ctx context.Context, key string) *re.StringSliceCmd
	HLen(ctx context.Context, key string) *re.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *re.SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *re.IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *re.BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *re.BoolCmd
	HVals(ctx context.Context, key string) *re.StringSliceCmd
	HRandField(ctx context.Context, key string, count int, withValues bool) *re.StringSliceCmd

	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *re.StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *re.StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *re.StringCmd
	LIndex(ctx context.Context, key string, index int64) *re.StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *re.IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *re.IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *re.IntCmd
	LLen(ctx context.Context, key string) *re.IntCmd
	LPop(ctx context.Context, key string) *re.StringCmd
	LPopCount(ctx context.Context, key string, count int) *re.StringSliceCmd
	LPos(ctx context.Context, key string, value string, args LPosArgs) *re.IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
	LPush(ctx context.Context, key string, values ...interface{}) *re.IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *re.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *re.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *re.IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *re.StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *re.StatusCmd
	RPop(ctx context.Context, key string) *re.StringCmd
	RPopCount(ctx context.Context, key string, count int) *re.StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *re.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *re.IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *re.IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *re.StringCmd
	// BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *re.StringCmd

	SAdd(ctx context.Context, key string, members ...interface{}) *re.IntCmd
	SCard(ctx context.Context, key string) *re.IntCmd
	SDiff(ctx context.Context, keys ...string) *re.StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *re.IntCmd
	SInter(ctx context.Context, keys ...string) *re.StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *re.IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *re.BoolCmd
	SMIsMember(ctx context.Context, key string, members ...interface{}) *re.BoolSliceCmd
	SMembers(ctx context.Context, key string) *re.StringSliceCmd
	SMembersMap(ctx context.Context, key string) *re.StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *re.BoolCmd
	SPop(ctx context.Context, key string) *re.StringCmd
	SPopN(ctx context.Context, key string, count int64) *re.StringSliceCmd
	SRandMember(ctx context.Context, key string) *re.StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *re.StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *re.IntCmd
	SUnion(ctx context.Context, keys ...string) *re.StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *re.IntCmd

	XAdd(ctx context.Context, a *XAddArgs) *re.StringCmd
	XDel(ctx context.Context, stream string, ids ...string) *re.IntCmd
	XLen(ctx context.Context, stream string) *re.IntCmd
	XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
	XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
	XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *re.StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *re.StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *re.StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *re.IntCmd
	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *re.IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *re.IntCmd
	XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *re.IntCmd
	XPending(ctx context.Context, stream, group string) *XPendingCmd
	XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
	XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
	XClaimJustID(ctx context.Context, a *XClaimArgs) *re.StringSliceCmd
	XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
	XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd

	// TODO: XTrim and XTrimApprox remove in v9.
	XTrim(ctx context.Context, key string, maxLen int64) *re.IntCmd
	XTrimApprox(ctx context.Context, key string, maxLen int64) *re.IntCmd
	XTrimMaxLen(ctx context.Context, key string, maxLen int64) *re.IntCmd
	XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *re.IntCmd
	XTrimMinID(ctx context.Context, key string, minID string) *re.IntCmd
	XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *re.IntCmd
	XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
	XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
	XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
	XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd

	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd

	// TODO: remove
	//		ZAddCh
	//		ZIncr
	//		ZAddNXCh
	//		ZAddXXCh
	//		ZIncrNX
	//		ZIncrXX
	// 	in v9.
	// 	use ZAddArgs and ZAddArgsIncr.

	ZAdd(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddNX(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddXX(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddCh(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddNXCh(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddXXCh(ctx context.Context, key string, members ...*Z) *re.IntCmd
	ZAddArgs(ctx context.Context, key string, args ZAddArgs) *re.IntCmd
	ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *re.FloatCmd
	ZIncr(ctx context.Context, key string, member *Z) *re.FloatCmd
	ZIncrNX(ctx context.Context, key string, member *Z) *re.FloatCmd
	ZIncrXX(ctx context.Context, key string, member *Z) *re.FloatCmd
	ZCard(ctx context.Context, key string) *re.IntCmd
	ZCount(ctx context.Context, key, min, max string) *re.IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *re.IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *re.FloatCmd
	ZInter(ctx context.Context, store *re.ZStore) *re.StringSliceCmd
	ZInterWithScores(ctx context.Context, store *re.ZStore) *re.ZSliceCmd
	ZInterStore(ctx context.Context, destination string, store *re.ZStore) *re.IntCmd
	ZMScore(ctx context.Context, key string, members ...string) *re.FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *re.ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *re.ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *re.StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *re.ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *re.ZRangeBy) *re.StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *re.ZRangeBy) *re.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *re.ZRangeBy) *re.ZSliceCmd
	ZRangeArgs(ctx context.Context, z re.ZRangeArgs) *re.StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z re.ZRangeArgs) *re.ZSliceCmd
	ZRangeStore(ctx context.Context, dst string, z re.ZRangeArgs) *re.IntCmd
	ZRank(ctx context.Context, key, member string) *re.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *re.IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *re.IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *re.IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *re.IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *re.StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *re.ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *re.ZRangeBy) *re.StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *re.ZRangeBy) *re.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *re.ZRangeBy) *re.ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *re.IntCmd
	ZScore(ctx context.Context, key, member string) *re.FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *re.ZStore) *re.IntCmd
	ZUnion(ctx context.Context, store re.ZStore) *re.StringSliceCmd
	ZUnionWithScores(ctx context.Context, store re,ZStore) *re.ZSliceCmd
	ZRandMember(ctx context.Context, key string, count int, withScores bool) *re.StringSliceCmd
	ZDiff(ctx context.Context, keys ...string) *re.StringSliceCmd
	ZDiffWithScores(ctx context.Context, keys ...string) *re.ZSliceCmd
	ZDiffStore(ctx context.Context, destination string, keys ...string) *re.IntCmd

	PFAdd(ctx context.Context, key string, els ...interface{}) *re.IntCmd
	PFCount(ctx context.Context, keys ...string) *re.IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *re.StatusCmd

	BgRewriteAOF(ctx context.Context) *re.StatusCmd
	BgSave(ctx context.Context) *re.StatusCmd
	ClientKill(ctx context.Context, ipPort string) *re.StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *re.IntCmd
	ClientList(ctx context.Context) *re.StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *re.BoolCmd
	ClientID(ctx context.Context) *re.IntCmd
	ConfigGet(ctx context.Context, parameter string) *re.SliceCmd
	ConfigResetStat(ctx context.Context) *re.StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *re.StatusCmd
	ConfigRewrite(ctx context.Context) *re.StatusCmd
	DBSize(ctx context.Context) *re.IntCmd
	FlushAll(ctx context.Context) *re.StatusCmd
	FlushAllAsync(ctx context.Context) *re.StatusCmd
	FlushDB(ctx context.Context) *re.StatusCmd
	FlushDBAsync(ctx context.Context) *re.StatusCmd
	Info(ctx context.Context, section ...string) *re.StringCmd
	LastSave(ctx context.Context) *re.IntCmd
	Save(ctx context.Context) *re.StatusCmd
	Shutdown(ctx context.Context) *re.StatusCmd
	ShutdownSave(ctx context.Context) *re.StatusCmd
	ShutdownNoSave(ctx context.Context) *re.StatusCmd
	SlaveOf(ctx context.Context, host, port string) *re.StatusCmd
	Time(ctx context.Context) *TimeCmd
	DebugObject(ctx context.Context, key string) *re.StringCmd
	ReadOnly(ctx context.Context) *re.StatusCmd
	ReadWrite(ctx context.Context) *re.StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *re.IntCmd

	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *re.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *re.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *re.BoolSliceCmd
	ScriptFlush(ctx context.Context) *re.StatusCmd
	ScriptKill(ctx context.Context) *re.StatusCmd
	ScriptLoad(ctx context.Context, script string) *re.StringCmd

	Publish(ctx context.Context, channel string, message interface{}) *re.IntCmd
	PubSubChannels(ctx context.Context, pattern string) *re.StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *re.StringIntMapCmd
	PubSubNumPat(ctx context.Context) *re.IntCmd

	ClusterSlots(ctx context.Context) *re.ClusterSlotsCmd
	ClusterNodes(ctx context.Context) *re.StringCmd
	ClusterMeet(ctx context.Context, host, port string) *re.StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *re.StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *re.StatusCmd
	ClusterResetSoft(ctx context.Context) *re.StatusCmd
	ClusterResetHard(ctx context.Context) *re.StatusCmd
	ClusterInfo(ctx context.Context) *re.StringCmd
	ClusterKeySlot(ctx context.Context, key string) *re.IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *re.StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *re.IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *re.IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *re.StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *re.StatusCmd
	ClusterSaveConfig(ctx context.Context) *re.StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *re.StringSliceCmd
	ClusterFailover(ctx context.Context) *re.StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int) *re.StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *re.StatusCmd

	GeoAdd(ctx context.Context, key string, geoLocation ...*re.GeoLocation) *re.IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *re.GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *re.GeoRadiusQuery) *re.GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *re.GeoRadiusQuery) *re.IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *re.GeoRadiusQuery) *re.GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *re.GeoRadiusQuery) *re.IntCmd
	GeoSearch(ctx context.Context, key string, q *re.GeoSearchQuery) *re.StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q *re.GeoSearchLocationQuery) *re.GeoSearchLocationCmd
	GeoSearchStore(ctx context.Context, key, store string, q *re.GeoSearchStoreQuery) *re.IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *re.FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *re.StringSliceCmd



	///configCenter

	Subscribe(ctx context.Context, channels ...string)*re.PubSub
	PSubscribe(ctx context.Context, channels ...string) *re.PubSub 
}
