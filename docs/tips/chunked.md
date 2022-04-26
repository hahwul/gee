---
title: Chunked
parent: Tips
has_children: false
nav_order: 3
---
`-chunked` flag is Specify the maximum length of the file and save it in multiple files.

```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -chunked 500 output
```

## References
- https://twitter.com/hahwul/status/1360495565633540097