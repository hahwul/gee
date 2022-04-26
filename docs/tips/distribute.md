---
title: Distribute
parent: Tips
has_children: false
nav_order: 4
---
`-distribute` flag is each line sequentially to multiple files.

```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -distribute alice.txt bob.txt charlie.txt
```

## References
- https://twitter.com/hahwul/status/1360495570922704897