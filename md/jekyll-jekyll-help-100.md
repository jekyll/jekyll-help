# [For Loops seem to break](https://github.com/jekyll/jekyll-help/issues/100)

> state: **closed** opened by: **** on: ****

I upgraded Jekyll and Ruby, and my for loops seem to have broken.

Liquid Exception: Syntax Error in &#x27;for loop&#x27; - Valid syntax: for [item] in [collection] in blog.html

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49777844) on: **Tuesday, July 22, 2014**

Would you please post your code or a not working example? The error is reported as a __syntax__ error.
---
> from: [**jcm2226**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49778616) on: **Tuesday, July 22, 2014**

This is the only for loop in the layout, and it still throws an error.

&lt;pre&gt;&lt;code&gt;{% for post in site.posts limit:4 %}

{{ post.title }}

{{ post.date | date: &#x27;%B&#x27; }} {{ post.date | date: &#x27;%e&#x27; }}, {{ post.date | date: &#x27;%Y&#x27; }}

{{ post.excerpt }}

{% endfor %}&lt;/code&gt;&lt;/pre&gt;
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49779258) on: **Tuesday, July 22, 2014**

Strange. Can you post some surrounding code? Because from what we have here this should work.
---
> from: [**jcm2226**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49780550) on: **Tuesday, July 22, 2014**

This is the entire body element:
&#x60;&#x60;&#x60;
&lt;body&gt;
  &lt;div class=&quot;row-fluid&quot;&gt;
    &lt;div class=&quot;span1 visible-desktop&quot;&gt;&lt;/div&gt;
 { % include menu.html % }

      &lt;div class=&quot;span12&quot; style=&quot;margin:0px;margin-top:10px&quot;&gt;
	  
	  &lt;div class=&quot;span3 well&quot; height=&quot;100%&quot;&gt;&lt;p class=&quot;lead&quot;&gt;Posts&lt;/p&gt;
{% for post in site.posts limit:4 %}
&lt;h2&gt;
    &lt;a href=&quot;/&quot; rel=&quot;bookmark&quot; title=&quot;Permanent link to &quot;&gt;{{ post.title }}&lt;/a&gt;
&lt;/h2&gt;
&lt;span&gt;{{ post.date | date: &#x27;%B&#x27; }} {{ post.date | date: &#x27;%e&#x27; }}, {{ post.date | date: &#x27;%Y&#x27; }}&lt;/span&gt;
&lt;p&gt;
    {{ post.excerpt }}
&lt;/p&gt;
{% endfor %}

        &lt;div class=&quot;span9&quot;&gt;
         &lt;p class=&quot;lead&quot;&gt; The Blog&lt;/p&gt;
		 &lt;div class=&quot;hero-unit&quot;&gt; {{content}} &lt;/div&gt;
		 &lt;div class=&quot;hero-unit&quot;&gt;&lt;/div&gt;
&lt;/div&gt;
    &lt;div class=&quot;span1 visible-desktop&quot;&gt;&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;&lt;div class=&quot;span12 navabar navbar-inner navbar-fixed-bottom&quot; style=&quot;margin:0px; padding-bottom:length;&quot;70px&quot;&gt;
     &lt;h6 class=&quot;pull-left&quot;&gt;&lt;a href=&quot;picturecredit.html&quot;&gt;Picture credits&lt;/a&gt;&lt;/h6&gt;
	 
	   &lt;h6 class=&quot;pull-right&quot; style=&quot;color: white&quot;&gt;Site &amp;copy; 2013&lt;/h6&gt;
    &lt;/div&gt;
  
&lt;/body&gt;
&#x60;&#x60;&#x60;
---
> from: [**jcm2226**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49780604) on: **Tuesday, July 22, 2014**

Wait a second. How do I turn HTML rendering off?
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49781251) on: **Tuesday, July 22, 2014**

Putting it in a code block with three back ticks: https://help.github.com/articles/github-flavored-markdown#fenced-code-blocks
---
> from: [**jcm2226**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49782216) on: **Tuesday, July 22, 2014**

Ah much better thanks.


On Tue, Jul 22, 2014 at 1:34 PM, Philipp Rudloff &lt;notifications@github.com&gt;
wrote:

&gt; Putting it in a code block with three back ticks:
&gt; https://help.github.com/articles/github-flavored-markdown#fenced-code-blocks
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49781251&gt;.
&gt;
---
> from: [**jcm2226**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49783440) on: **Tuesday, July 22, 2014**

It worked before I upgraded Jekyll. I updated my Ruby environment to 2.1.1. I read the changelog for the latest release of jekyll. I checked my post markdown files, and they seem ok. I read the liquid bug reports. I&#x27;m just stumped. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49953335) on: **Wednesday, July 23, 2014**

I don&#x27;t have time to test it myself, but I&#x27;d be willing to bet it is breaking on the &#x60;limit:4&#x60; business. I&#x27;d recommend checking out [Jekyll&#x27;s Pagination docs](http://jekyllrb.com/docs/pagination/) and use that instead of trying to drop a filter into your &#x60;for&#x60; loop.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-49967316) on: **Wednesday, July 23, 2014**

I&#x27;m using &#x60;post in site.posts&#x60; to generate my blog posts as well. &#x60;limit:n&#x60; should not be the problem: https://github.com/kleinfreund/kleinfreund.github.io/blob/dev/index.html#L7-L36
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-50049959) on: **Thursday, July 24, 2014**

@jcm2226 Your include syntax is broken:

&#x60;&#x60;&#x60;
{ % include menu.html % }
&#x60;&#x60;&#x60;

It must be:


&#x60;&#x60;&#x60;
{% include menu.html %}
&#x60;&#x60;&#x60;


---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/100#issuecomment-55628698) on: **Monday, September 15, 2014**

@jcm2226 What&#x27;s the status of this? A lot of us have given some solutions, and I know at least I am curious to know if you got it working. I&#x27;m going to close this issue for now, but if you&#x27;re still having trouble, feel free to re-open it and let us know.
