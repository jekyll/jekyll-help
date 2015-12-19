# [site.tags list omits page title](https://github.com/jekyll/jekyll-help/issues/57)

> state: **closed** opened by: **** on: ****

If you iterate through site.tags in a for loop the first value of the array is the tag, and the second is the content of all the pages, but not including the headings or any separators between that content. Effectively it just looks like a big mess. 

The functionality would be much more useful if the first item was the tag, and the second was a list of page objects with the tag - so we could get title and any other values defined in the page itself.

This can of course be emulated by iterating the pages, but it is a bit more painful that it needs to be.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/57#issuecomment-48136825) on: **Sunday, July 06, 2014**

@hamishwillee Hey there. Sorry for taking so long to respond. This is a tricky one, since you have to deal with iterating over a Liquid hash. There really isn&#x27;t any good documentation on this, unfortunately, but [this Stack Overflow answer](http://stackoverflow.com/questions/8206869/iterate-over-hashes-in-liquid-templates) is probably the best at explaining those pesky punks:

&gt; When you iterate over a hash using a variable called &#x60;hash&#x60;, &#x60;hash[0]&#x60; contains the key and &#x60;hash[1]&#x60; contains the value on each iteration.

Here&#x27;s an implementation that might work for you:

&#x60;&#x60;&#x60;html
&lt;ul&gt;
{% for tag in site.tags %}
  &lt;li&gt;&lt;span class=&quot;tag&quot;&gt;{{ tag[0] }}&lt;/span&gt;
    &lt;ul&gt;
      {% for p in tag[1] %}
      &lt;li&gt;&lt;a href=&quot;{{ p.url }}&quot;&gt;{{ p.title }}&lt;/a&gt;&lt;/li&gt;
      {% endfor %}
    &lt;/ul&gt;
  &lt;/li&gt;
{% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;

Closing this issue for now, but if you have any questions, feel free to add a comment!
---
> from: [**hamishwillee**](https://github.com/jekyll/jekyll-help/issues/57#issuecomment-48139723) on: **Sunday, July 06, 2014**

Hi Troy, 

Thanks for that. Are you saying that tag[1] is an iterable list of the pages - not just the page dump as I thought? Interesting! That certainly should work, though I didn&#x27;t see the title in the page dump. 

The way I solved this problem was to iterate through the tags then iterate the pages. Your solution is better, and I will try it at some point!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/57#issuecomment-48143231) on: **Sunday, July 06, 2014**

&gt; Are you saying that tag[1] is an iterable list of the pages - not just the page dump as I thought?

Yes, exactly! If you have two tags (&quot;food&quot; and &quot;drink&quot;) and a few posts that use those tags, the &#x60;site.tags&#x60; hash would look something like this:

&#x60;&#x60;&#x60;
[&#x27;food&#x27;] =&gt; list(
  Post,
  Post
),
[&#x27;drink&#x27;] =&gt; list(
  Post
)
&#x60;&#x60;&#x60;

In this example, when you iterate on &#x60;site.tags&#x60;, the first iteration will receive the &#x60;food&#x60; item. If your iterator variable is &#x60;tag&#x60; (e.g. &#x60;for tag in site.tags&#x60;), then &#x60;tag[0]&#x60; would be &quot;food&quot; and &#x60;tag[1]&#x60; would be a list of Posts.

Iterating on that list of Posts allows you to do the same kind of things you would normally do inside of a post&#x27;s layout. I used &#x60;for p in tag[1]&#x60;, so doing stuff like &#x60;{{ p.url }}&#x60; outputs the URL for that post, &#x60;{{ p.title }}&#x60; outputs the title, and so on.

---

The way Liquid handles outputting hash objects is a little wonky, imo. Experiment with it a little bit by outputting the &#x60;site.tags&#x60; hash directly onto a page somewhere (using &#x60;{{ site.tags }}&#x60;).

You&#x27;ll see something that looks like: &#x60;{&quot;food&quot;=&gt;[, ], &quot;drink&quot;=&gt;[]}&#x60;. Unfortunately, Jekyll doesn&#x27;t have a way to represent a Post in a list like this, so the values in the list are a vasty nothingness. That said, you know that there is at least something inside that list, otherwise the key for that tag wouldn&#x27;t be there.

---

Experimentation goes a long way when trying to understand hashes in Liquid. Try putting this in your front-matter:

&#x60;&#x60;&#x60;yaml
hash:
  hello: [&#x27;world&#x27;, &#x27;earth&#x27;, &#x27;sunshine&#x27;]
  ahoy: mateys
&#x60;&#x60;&#x60;

If you just output the &#x60;hash&#x60; value on that page with &#x60;{{ hash }}&#x60;, it returns with:

&gt; {&quot;hello&quot;=&gt;[&quot;world&quot;, &quot;earth&quot;, &quot;sunshine&quot;], &quot;ahoy&quot;=&gt;&quot;mateys&quot;}

Reasonable enough, right? But when you try to loop on that hash using Liquid, like so...

&#x60;&#x60;&#x60;html
{% for h in page.hash %}
&lt;div&gt;{{ h }}&lt;/div&gt;
{% endfor %}
&#x60;&#x60;&#x60;

...you wind up with this:

&gt; helloworldearthsunshine
&gt; ahoymateys

As you can see, the hash key and the value are all jumbled together, and even the values of the list in the first item get squished. In order to make sense of these, you need to use &#x60;val[0]&#x60; and &#x60;val[1]&#x60;.

Here&#x27;s some example code of how to get this looking pretty (I added some additional spaces to show loop nesting):

&#x60;&#x60;&#x60;html
&lt;dl&gt;
  {% for h in page.hash %}
    &lt;dt&gt;{{ h[0] }}&lt;/dt&gt;
    {% for i in h[1] %}
      &lt;dd&gt;{{ i }}&lt;/dd&gt;
    {% endfor %}
  {% endfor %}
&lt;/dl&gt;
&#x60;&#x60;&#x60;

When rendered by a browser, here&#x27;s what it looks like:

![screen shot 2014-07-07 at 12 40 07 am](https://cloud.githubusercontent.com/assets/1420926/3491677/80e09430-0599-11e4-9025-8e103eb891d9.png)

---

I&#x27;ll admit, I may have gone a little overkill answering this question. However, dealing with hashes was one of the biggest pains in my ass when I started using Jekyll. Hopefully this will serve some good for folks who are getting into it and having the same trouble I did.

If you have any other questions, feel free to respond here! I&#x27;ll do my best to help.
