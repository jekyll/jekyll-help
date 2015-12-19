# [Leading whitespace is stripped from Markdown highlight block](https://github.com/jekyll/jekyll-help/issues/229)

> state: **closed** opened by: **** on: ****

Following from the [lanyon](https://github.com/poole/lanyon) base Jekyll theme, I have Markdown in a file of the form (this being output from &#x60;R&#x60;):

&#x60;&#x60;&#x60;
{% highlight text %}
      [,1] [,2]
[1,] FALSE TRUE
[2,] FALSE TRUE
{% endhighlight %}
&#x60;&#x60;&#x60;

which, after render (through whatever machinery runs in &#x60;jekyll serve&#x60;), gets generated HTML as:

&#x60;&#x60;&#x60;
&lt;div class=&quot;highlight&quot;&gt;&lt;pre&gt;&lt;code class=&quot;language-text&quot; data-lang=&quot;text&quot;&gt;[,1]
[1,]    0
[2,]    1
[3,]    2
[4,]    2&lt;/code&gt;&lt;/pre&gt;&lt;/div&gt;
&#x60;&#x60;&#x60;

with the leading whitespace in front of &#x60;[,1]&#x60; removed. 

Is there a way to ensure that leading whitespace is not stripped on render -- perhaps something I can specify in &#x60;_config.yml&#x60;, or otherwise? FWIW, I have:

&#x60;&#x60;&#x60;
markdown:            kramdown
highlighter:         pygments
&#x60;&#x60;&#x60;

in my &#x60;_config.yml&#x60;, alongside other (obstensibly) unrelated options.

My understanding is that pygments is handling this translation, and [the documentation](http://pygments.org/docs/quickstart/) here suggests that this might (?) be controlled through the &#x60;stripall&#x60; argument, but I am not sure how I would get access to that (if that is indeed what I need).

Finally -- in the end, my Jekyll website is going to be served through GitHub pages, so the solution would preferably not involve post-processing the generated HTML.

Thanks!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/229#issuecomment-69059843) on: **Wednesday, January 07, 2015**

From the sounds of it, this may be a bug that you want to file in the actual Jekyll repo.

For your particular issue, you might want to try using tables.

 | [,1] | [,2]
--- | --- | ---
**[1,]** | FALSE | TRUE
**[2,]** | FALSE | TRUE

&#x60;&#x60;&#x60;
 | [,1] | [,2]
--- | --- | ---
**[1,]** | FALSE | TRUE
**[2,]** | FALSE | TRUE
&#x60;&#x60;&#x60;
---
> from: [**kevinushey**](https://github.com/jekyll/jekyll-help/issues/229#issuecomment-69061790) on: **Wednesday, January 07, 2015**

Thanks -- I&#x27;ll do that, although in general I do want to print generic, verbatim code output without any whitespace stripping.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/229#issuecomment-69063102) on: **Wednesday, January 07, 2015**

Fair enough. Open an issue in the actual Jekyll repo and see if they can look at it.
