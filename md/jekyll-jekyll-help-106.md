# [textarea with linebreak?](https://github.com/jekyll/jekyll-help/issues/106)

> state: **closed** opened by: **** on: ****

In my markdown file, I have a simple HTML &#x60;textarea&#x60; code like this:

&#x60;&#x60;&#x60;&#x60;
&lt;textarea&gt;
hello

world
&lt;/textarea&gt;
&#x60;&#x60;&#x60;&#x60;

You can see that there are few line-breaks between *hello* and *world*. The problem is that &#x60;jekyll build&#x60; then removes the line-breaks between them.

How can I tell jekyll not to do this?

Thanks.

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50717211) on: **Wednesday, July 30, 2014**

We just pass it along to the markdown converter. Which one are you using? You&#x27;ll have to ask them.
---
> from: [**aquynh**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50722065) on: **Thursday, July 31, 2014**

I am not really sure. I am just using the default converter on jekyll. How can I know which converter it is using?

Thanks.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50728937) on: **Thursday, July 31, 2014**

If you use the current Markdown engine, that&#x27;s kramdown.
---
> from: [**aquynh**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50733231) on: **Thursday, July 31, 2014**

Great! But to be sure, how can I verify that I am indeed using kramdown?

Thanks.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50735299) on: **Thursday, July 31, 2014**

Go to your _config.yml and check if you have specified &#x60;markdown&#x60; anywhere. If not, the default valut for that is kramdown: http://jekyllrb.com/docs/configuration/#default-configuration
---
> from: [**aquynh**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50735746) on: **Thursday, July 31, 2014**

My _config.yml has no such information on markdown.

Oh and my Jekyll is 1.3.0. Is this version using kramdown, or other converter? I tried to find out, but cannot get this information anywhere.

Thanks.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50736498) on: **Thursday, July 31, 2014**

&#x60;&#x60;&#x60; jinja
&lt;p&gt;My current markdown engine is as follows:  {{ site.markdown }}.&lt;/p&gt;
&#x60;&#x60;&#x60;
---
> from: [**aquynh**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50737095) on: **Thursday, July 31, 2014**

thanks!

i will redirect my question to Kramdown guys.
---
> from: [**aquynh**](https://github.com/jekyll/jekyll-help/issues/106#issuecomment-50849373) on: **Thursday, July 31, 2014**

i asked Kramdown guys, and there is indeed a bug in Kramdown. they told me
how to fix that in Kramdown here:

    https://github.com/gettalong/kramdown/issues/147#issuecomment-50740471

the problem is that now Kramdown is working correctly (after the fix), but
Jekyll (2.2.0) still removes line-breaks in &lt;textarea&gt;.

could you please confirm this is indeed a bug in Jekyll? if so, any idea
how to fix that?

thanks.
