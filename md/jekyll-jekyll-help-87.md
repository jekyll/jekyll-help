# [Changing class of div created by Pygments?](https://github.com/jekyll/jekyll-help/issues/87)

> state: **closed** opened by: **** on: ****

Howdy all!

As you know Pygments creates a &#x60;&lt;div&gt;&#x60; and assigns a class called &#x60;highlight&#x60; to it.

Is it possible to change that class to a custom defined class?

I&#x27;ve looked at Pygments&#x27; documentation and it is possible to pass in an option on the command line, called &#x60;cssclass&#x60;, to change the class of the &#x60;&lt;div&gt;&#x60;. Jekyll documentation doesn&#x27;t have mention of this, so I was wondering if anyone knows how to do it with Jekyll?

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503132) on: **Monday, June 30, 2014**

You could *try*:

&#x60;&#x60;&#x60;text
{% highlight ruby cssclass=&quot;myclass&quot; %}
my_ruby.call()
{% endhighlight %}
&#x60;&#x60;&#x60;

Let me know if that works.
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503170) on: **Monday, June 30, 2014**

I actually tried that. Nope, didn&#x27;t work. :frowning: 
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503338) on: **Monday, June 30, 2014**

Tbh, I don&#x27;t think there is an option for this. The error message returns a valid syntax example and it does not suggest what I&#x27;m trying to do is possible:

&#x60;&#x60;&#x60;
Valid syntax: highlight &lt;lang&gt; [linenos]
&#x60;&#x60;&#x60;
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503423) on: **Monday, June 30, 2014**

Maybe this could be a feature request?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503507) on: **Monday, June 30, 2014**

It looks like it should be working with Jekyll 2.1.0: https://github.com/jekyll/jekyll/pull/2532/files#diff-4ef4410ae4c188981d2f90c33c569487R82
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47503642) on: **Monday, June 30, 2014**

Hmmm, indeed.

I need to upgrade.
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47504573) on: **Monday, June 30, 2014**

Updated to 2.1.0 and it still doesn&#x27;t work. Same error message.

Removing the invalid syntax returns a message saying Pygments 0.6.0 needs to be installed.

But upgrading Pygments past v0.5.0 will cause problems on Windows. :frowning: However you&#x27;re already aware of this in https://github.com/jekyll/jekyll/issues/2556

Seems like I have to wait for a fix or change OS :laughing: 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47504673) on: **Monday, June 30, 2014**

@winghouchan &lt;s&gt;Is pygments 0.6.0 still causing issues? Oh. :(&lt;/s&gt; Looks like @parkr is ahead of time: https://github.com/tmm1/pygments.rb/issues/129
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47504861) on: **Monday, June 30, 2014**

@kleinfreund thanks for chiming in again! Ya, Pygments 0.6.0 doesn&#x27;t work on Windows. I&#x27;m still on Pygments 0.5.0. However I needed Pygments 0.6.0 so Jekyll 2.1.0 would work. I need Jekyll 2.1.0 as there is some new functionality where you can change the class assigned to the &#x60;&lt;div&gt;&#x60; wrapping the highlighted code.
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47505106) on: **Monday, June 30, 2014**

@parkr and @kleinfreund

Thanks for your help with this again!

I think there is no way for me to change the class of the &#x60;&lt;div&gt;&#x60; until I change OS (which I don&#x27;t think I will) or until there is a fix.

So I think I&#x27;ll have to wait for a fix. In the meantime, the inability to change the class isn&#x27;t really a big problem. I could just use some dirty JS to change it client-side :laughing: or not change it all.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47511371) on: **Monday, June 30, 2014**

This should be solved by the new release of posix-spawn! In the meantime, you can bundle the master branch of rtomayko/posix-spawn to get it working now. Let me know if you run into more problems!
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/87#issuecomment-47512338) on: **Monday, June 30, 2014**

@parkr Okey dokey, thanks!
