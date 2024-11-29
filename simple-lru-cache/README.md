# Simple lru-cache

Basi implementation of a lru-cache.
The cache have a fixed size. 

## Appending objects to the cache
Object are added to the head of the cache.
If the object already exist in the cache it is moved to the head of the cache.
If the cache is full, the Tail is removed object to make space.