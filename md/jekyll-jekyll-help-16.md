# [Variables in include variables](https://github.com/jekyll/jekyll-help/issues/16)

> state: **closed** opened by: **** on: ****

Is there any way to make something like this work?

    {% include navigation/social_listing.html url=&quot;mailto:{{ site.author.email }}&quot; shortname=&quot;mail&quot; label=&quot;Email&quot; %}

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40381504) on: **Monday, April 14, 2014**

**_includes/navigation/social_listing.html**

&#x60;&#x60;&#x60;
&lt;li&gt;&lt;a href=&quot;{{ include.url }}&quot; class=&quot;{{ include.shortname }}&quot; onClick=&quot;_gaq.push([&#x27;_trackEvent&#x27;, &#x27;Social Links&#x27;, &#x27;Click&#x27;, &#x27;{{ include.label }}&#x27;]);&quot;&gt;
  &lt;i class=&quot;fi-{% if include.shortname != &#x27;mail&#x27; %}social-{% endif %}{{ include.shortname }}&quot;&gt;&lt;/i&gt;
  &lt;span&gt;{{ include.label }}&lt;/span&gt;
&lt;/a&gt;&lt;/li&gt;
&#x60;&#x60;&#x60;

Using the above &#x60;{% include %}&#x60; causes the href for the link to be the literal value in the string instead of interpolating the email address in.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40398875) on: **Monday, April 14, 2014**

Try it without directly passing email address

&#x60;&#x60;&#x60;
{% assign email_address = {{site.author.email}} %}
{% include social_listing.html shortname=&quot;mail&quot; label=&quot;Email&quot; %}
&#x60;&#x60;&#x60;

&#x60;&#x60;&#x60;
&lt;li&gt;&lt;a href=&quot;{{email_address}}&quot; class=&quot;{{ include.shortname }}&quot; onClick=&quot;_gaq.push([&#x27;_trackEvent&#x27;, &#x27;Social Links&#x27;, &#x27;Click&#x27;, &#x27;{{ include.label }}&#x27;]);&quot;&gt;
  &lt;i class=&quot;fi-{% if include.shortname != &#x27;mail&#x27; %}social-{% endif %}{{ include.shortname }}&quot;&gt;&lt;/i&gt;
  &lt;span&gt;{{ include.label }}&lt;/span&gt;
&lt;/a&gt;&lt;/li&gt;
&#x60;&#x60;&#x60;



---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40399478) on: **Monday, April 14, 2014**

Hi Bud, that may work, but it won&#x27;t help the issue. The reason it&#x27;s an
include with a variable in the first place is that the include is called a
dozen times with different values.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40400088) on: **Monday, April 14, 2014**

I just meant that whenever you call the include, you just put the value in a separate assign tag rather than in the include. Maybe it would have been more clear if I made it
&#x60;&#x60;&#x60;
instance 1
 {% assign url = {{site.author.email}} %}
{% include social_listing.html shortname=&quot;mail&quot; label=&quot;Email&quot; %}
&#x60;&#x60;&#x60;
&#x60;&#x60;&#x60;
instance2
 {% assign url = someOtherURL %}
{% include social_listing.html shortname=&quot;somethingelse&quot; label=&quot;twitter&quot; %}
[edited]
&#x60;&#x60;&#x60;

&#x60;&#x60;&#x60;
&lt;li&gt;&lt;a href=&quot;{{url}}&quot; class=&quot;{{ include.shortname }}&quot; onClick=&quot;_gaq.push([&#x27;_trackEvent&#x27;, &#x27;Social Links&#x27;, &#x27;Click&#x27;, &#x27;{{ include.label }}&#x27;]);&quot;&gt;
  &lt;i class=&quot;fi-{% if include.shortname != &#x27;mail&#x27; %}social-{% endif %}{{ include.shortname }}&quot;&gt;&lt;/i&gt;
  &lt;span&gt;{{ include.label }}&lt;/span&gt;
&lt;/a&gt;&lt;/li&gt;
&#x60;&#x60;&#x60;

or, maybe I&#x27;m just reading your intent wrong.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40407388) on: **Monday, April 14, 2014**

Found an even better solution. I forgot that you actually *can* pass variables directly through &#x60;{% include %}&#x60;. Final solution:

&#x60;&#x60;&#x60;
{% capture author_email_link %}mailto:{{ site.author.email }}{% endcapture %}
{% include navigation/social_listing.html url=author_email_link shortname=&quot;mail&quot; label=&quot;Email&quot; %}
&#x60;&#x60;&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40407548) on: **Monday, April 14, 2014**

You might be able to use &#x60;{% assign %}&#x60;, too. It would be cleaner, but I&#x27;m not 100% sure it would work.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40407703) on: **Monday, April 14, 2014**

From what I read in the Liquid docs, assign suffers the same problem as the
original issue, namely it&#x27;s for assigning static values, it won&#x27;t do the
interpolation. That&#x27;s the reason capture exists.

http://docs.shopify.com/themes/liquid-basics/logic#assign
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40408376) on: **Monday, April 14, 2014**

I just tested something that would look like this for you:

&#x60;&#x60;&#x60;
{% assign author_email_link = &#x27;mailto:&#x27; | append: site.author.email %}
&#x60;&#x60;&#x60;

And it worked! That might be cleaner and more maintainable for you. :smile: 
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/16#issuecomment-40408929) on: **Monday, April 14, 2014**

Sure, why not. https://bitbucket.org/matthew-scharley/matt.scharley.me/src/24535323068d66d48f20a7f0966a6472573eaba0/src/_includes/navigation/social.html
