if (redis.call('EXISTS', KEYS[1]) == 0) then
    if redis.call('HMSET',KEYS[1],ARGV[1]) then
       redis.call('PUBLISH', KEYS[2],AEGV[2])
    end
    redis.call('HINCRBY', KEYS[1],KEYS[3],1)
    
end


eval "if (redis.call('EXISTS', KEYS[1]) == 0) then if redis.call('HMSET',KEYS[1],ARGV[1]) then redis.call('PUBLISH', KEYS[2],AEGV[2]) end end redis.call('HINCRBY', KEYS[1],KEYS[3],1)" 3 "routerCenter/default/test" "routerCenter/default/subscribeReg" life comtext aaa  test
