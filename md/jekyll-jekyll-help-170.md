# [Any idea of how get Jekyll work under c9.io?](https://github.com/jekyll/jekyll-help/issues/170)

> state: **closed** opened by: **** on: ****

Hi there, I&#x27;m a big fan of c9.io ide, as I&#x27;m working with a Chromebook, but there seems I&#x27;m not doing something well, as I can&#x27;t get Jekyll (or Poole) working under my c9 project.

It&#x27;s there anyone who managed to achieve it?

Regards

### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/170#issuecomment-58924309) on: **Monday, October 13, 2014**

Hi!
I just set up a custom workspace on c9, and managed to get Jekyll working from there:

1. In Terminal, install Jekyll with &#x60;gem install jekyll&#x60;
2. &#x60;cd&#x60; to the directory that contains your Jekyll site
3. Run &#x60;jekyll serve --port $PORT&#x60;
4. Visit your site at &#x60;http://&lt;workspacename&gt;-c9-&lt;username&gt;.c9.io&#x60;
---
> from: [**isorna**](https://github.com/jekyll/jekyll-help/issues/170#issuecomment-59167174) on: **Wednesday, October 15, 2014**

Thank you very much Alfred!
