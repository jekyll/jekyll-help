# [Having issues running &#x27;gem install jekyll&#x27;](https://github.com/jekyll/jekyll-help/issues/28)

> state: **closed** opened by: **** on: ****

I have all the requirements to run this command, however this keeps popping up.

ubuntu:~$ gem install jekyll
ERROR:  While executing gem ... (Gem::FilePermissionError)
    You don&#x27;t have write permissions into the /var/lib/gems/1.8 directory.

Do I run &#x27;sudo gem install jekyll&#x27; or will I screw something up doing that?

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/28#issuecomment-42144262) on: **Sunday, May 04, 2014**

If you&#x27;re using the system default ruby (which is looks like in your case)
then using &#x60;sudo gem install&#x60; is fine.


On 5 May 2014 05:59, Max &lt;notifications@github.com&gt; wrote:

&gt; I have all the requirements to run this command, however this keeps
&gt; popping up.
&gt;
&gt; ubuntu:~$ gem install jekyll
&gt; ERROR: While executing gem ... (Gem::FilePermissionError)
&gt; You don&#x27;t have write permissions into the /var/lib/gems/1.8 directory.
&gt;
&gt; Do I run &#x27;sudo gem install jekyll&#x27; or will I screw something up doing that?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/28&gt;
&gt; .
&gt;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/28#issuecomment-42151219) on: **Sunday, May 04, 2014**

It seems that conversation moved to #29, I will close this issue to keep the conversation in one thread. BTW if you were using ruby 1.8, Jekyll 1.5.1+ will not work, updating ruby is recommended. Check [issue #8](https://github.com/jekyll/jekyll-help/issues/8).
