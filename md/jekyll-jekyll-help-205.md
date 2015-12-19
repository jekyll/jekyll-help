# [Post layout isn&#x27;t working](https://github.com/jekyll/jekyll-help/issues/205)

> state: **closed** opened by: **** on: ****

I followed all the directions on creating a .md post (e.g. 2011-12-31-new-years-eve-is-awesome.md), putting it in the _posts folder and included the YAML front matter and it still won&#x27;t recognize the layout. I am currently running it on my local server. 

When I look in my &quot;site&quot; folder and actually see the HTML code for the post, Jekyll is putting the front matter in the content &lt;p&gt; elements. 

Any suggestions? Thank you!

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65182438) on: **Monday, December 01, 2014**

Can you post a link to your repo?
---
> from: [**PaulTuraew**](https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65184126) on: **Monday, December 01, 2014**

Hi Bud,

https://github.com/PaulTuraew/turaew_io_jekyll_blog

Thanks in advance for the help!
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65230454) on: **Tuesday, December 02, 2014**

Hi Paul,
So, a couple of things. You have empty lines at the top of your .md file, so Jekyll is not properly interpreting your front matter. 

Also, have a look at your layout files. Your post layout doesn&#x27;t have everything it needs to render a page, yet isn&#x27;t referencing a default layout that has that either. You are calling a CSS file at the top of the page that isn&#x27;t doing anything because it&#x27;s not in the &lt;head&gt;.  And it looks like you&#x27;re using your header.html include the way you should be using a default template (because you have your closing body tags in their too).

Here are some examples from  the templates that Jekyll installs for you when you make a new instance:

This is the default layout that has a lot of the HTML for your page, but importantly a &#x60;&#x60;&#x60;{{ content }}&#x60;&#x60;&#x60; tag that allows other templates to inherit this template. 
&#x60;&#x60;&#x60;
&lt;!DOCTYPE html&gt;
&lt;html&gt;
  {% include head.html %}
  &lt;body&gt;
    &lt;div class=&quot;page-content&quot;&gt;
        {{ content }}
    &lt;/div&gt;
    {% include footer.html %}
  &lt;/body&gt;
&lt;/html&gt;
&#x60;&#x60;&#x60;

And this is a post template. Everything in this template gets inserted by Jekyll into the &#x60;&#x60;&#x60;{{content}}&#x60;&#x60;&#x60; tag in the default template. Note that at the top of this template, it references the default template&#x60;&#x60;&#x60;layout: default&#x60;&#x60;&#x60;
&#x60;&#x60;&#x60;
---
layout: default
---
&lt;div class=&quot;post&quot;&gt;
  &lt;header class=&quot;post-header&quot;&gt;
    &lt;h1 class=&quot;post-title&quot;&gt;{{ page.title }}&lt;/h1&gt;
  &lt;/header&gt;
  &lt;article class=&quot;post-content&quot;&gt;
    {{ content }}
  &lt;/article&gt;
&lt;/div&gt;

&#x60;&#x60;&#x60;
I&#x27;d suggest doing a fresh install of Jekyll, separate from your current site and having a look at the templates they use there. Hope that helps!

---
> from: [**PaulTuraew**](https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65439837) on: **Wednesday, December 03, 2014**

Hi Bud,

Thanks for the info, I looked into it, compared the templates and got the post layout to function correctly. Thank you again!

Paul
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65440260) on: **Wednesday, December 03, 2014**

Glad to hear that!

On Wed, Dec 3, 2014, 11:28 AM Paul &lt;notifications@github.com&gt; wrote:

&gt; Hi Bud,
&gt;
&gt; Thanks for the info, I looked into it, compared the templates and got the
&gt; post layout to function correctly. Thank you again!
&gt;
&gt; Paul
&gt;
&gt; --
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/205#issuecomment-65439837&gt;.
&gt;
