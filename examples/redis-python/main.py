import redis

# Connect to localhost on port 6379
# To receive decoded strings, set decode_responses=True
r = redis.Redis(host='localhost', port=6379, decode_responses=True)
r.set('foo', 'bar')
# True

# bar
print(r.get('foo'))
