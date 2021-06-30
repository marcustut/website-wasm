const cacheName = "app-" + "2c97746469ee6a87e120ba47a8d98eb0a76dc235";

self.addEventListener("install", event => {
  console.log("installing app worker 2c97746469ee6a87e120ba47a8d98eb0a76dc235");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/website-wasm",
        "/website-wasm/app.css",
        "/website-wasm/app.js",
        "/website-wasm/manifest.webmanifest",
        "/website-wasm/wasm_exec.js",
        "/website-wasm/web/app.wasm",
        "/website-wasm/web/tailwind.css",
        "https://storage.googleapis.com/murlok-github/icon-192.png",
        "https://storage.googleapis.com/murlok-github/icon-512.png",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 2c97746469ee6a87e120ba47a8d98eb0a76dc235 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
