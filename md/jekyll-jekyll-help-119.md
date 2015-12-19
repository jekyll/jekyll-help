# [How do I prevent {{ page.url }} returning a .html extension on GHP?](https://github.com/jekyll/jekyll-help/issues/119)

> state: **closed** opened by: **** on: ****

When I use {{ page.url }} locally it returns /page-name/. When I use it {{ page.url }} on a GitHub Pages hosted site it returns /page-name.html. 

How can I get the /page-name/ format on a GitHub Pages hosted site?

### Comments

---
> from: [**barryclark**](https://github.com/jekyll/jekyll-help/issues/119#issuecomment-51858618) on: **Monday, August 11, 2014**

Looks like I should probably use {{ page.id }} instead?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/119#issuecomment-55644551) on: **Monday, September 15, 2014**

Try including &#x60;permalink: pretty&#x60; in your &#x60;_config.yml&#x60; file. That is what I have on my site and it is doing exactly what you&#x27;re asking for.

There&#x27;s a fancy doc page about [permalinks](http://jekyllrb.com/docs/permalinks/) that might be relevant to your interests, too.
