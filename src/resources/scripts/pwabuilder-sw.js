/* eslint no-unused-vars: 0 no-restricted-globals: 0 */
// This is the "Offline copy of pages" wervice worker

// Install stage sets up the index page (home page) in the cache
// and opens a new cache
self.addEventListener('install', (event) => {
  const indexPage = new Request('/')
  event.waitUntil(fetch(indexPage).then(response =>
    caches.open('pwabuilder-offline').then(cache =>
      cache.put(indexPage, response))))
})

// If any fetch fails, it will look for the request in the cache
// and serve it from there first
self.addEventListener('fetch', (event) => {
  const updateCache = request =>
    caches.open('pwabuilder-offline').then(cache =>
      fetch(request).then(response => cache.put(request, response)))

  event.waitUntil(updateCache(event.request))

  event.respondWith(fetch(event.request).catch(error =>
    // Check to see if you have it in the cache
    // Return response
    // If not in the cache, then return error page
    caches.open('pwabuilder-offline').then(cache =>
      cache.match(event.request).then(matching =>
        (!matching || matching.status === 404 ? Promise.reject(new Error('no-match')) : matching)))))
})
