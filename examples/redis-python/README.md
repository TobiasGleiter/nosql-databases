# 🚀 Use Redis with Python

## 📋 Checklist to setup redis yourself

1. Setup virtual environment (recomended): `python -m venv ./venv`
2. Activate the virtual environment:

- POSIX: `source <venv>/bin/activate` (bash/zsh)
- Windows: `<venv>\Scripts\activate.bat` (cmd.exe)

3. `pip install redis`
4. Create a main.py and add `import redis`
5. Connect to the Redis localhost: `r = redis.Redis(host='localhost', port=6379, decode_responses=True)`
6. Store and retrieve a simple string:

```python
r.set('foo', 'bar')
# True
print(r.get('foo'))
# bar
```
