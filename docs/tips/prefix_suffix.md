---
title: Prefix and Suffix
parent: Tips
has_children: false
nav_order: 1
---

```
▶ cat urls | gee -prefix "curl -i -k " -suffix " -H 'Auth: abcd'" curls.sh
```

```
curl -i -k https://www.hahwul.com/?q=123 -H 'Auth: abcd'
curl -i -k http://testphp.vulnweb.com/listproducts.php?cat=asdf&ff=1 -H 'Auth: abcd'
curl -i -k https://xss-game.appspot.com/level1/frame  -H 'Auth: abcd'
```