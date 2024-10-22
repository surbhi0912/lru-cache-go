#### LRU-Cache

Aim is to build an in memory cache.

LRU cache is one that when low on memory,
evicts item that was used least recently.

Hashtable for ```Get``` and ```Set``` operations being ```O(1)```

```container/list``` is a doubly linked list, where whenever we make
use of a an item of the cache, whether by ```Get``` or ```Set```,
we promote the object to the front of the linkedlist.
So whenever we need to remove an item from the cache, we trim from
the tail of the list.