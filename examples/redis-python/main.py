import redis
# Connect to localhost on port 6379
# To receive decoded strings, set decode_responses=True
r = redis.Redis(host='localhost', port=6379, decode_responses=True)
r.set('foo', 'bar')
r.get('foo')
r.setnx('foo', 'bar')
r.getset('foo', 'rab')
r.mset({'fo': 'ba', 'f': 'b'})
r.mget({'foo', 'fo', 'f'})
