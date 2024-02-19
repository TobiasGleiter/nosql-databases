# 🚀 Use Redis with Node.js

## 📋 Checklist

1. Run `npm init -y` to setup the node project (optional)
2. Install redis with `npm install redis`
3. Create a new `index.js` and connect to the client:

```javascript
import { createClient } from "redis";

const client = createClient();

client.on("error", (err) => console.log("Redis Client Error", err));

await client.connect();
```

4. Store and retrieve a simple string:

```javascript
await client.set("key", "value");
const value = await client.get("key");
console.log(value);
```

5. Start the server with `npm start` or `npm run start`
