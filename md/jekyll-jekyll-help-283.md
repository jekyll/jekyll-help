# [How do I use variables with Liquid in Jekyll?](https://github.com/jekyll/jekyll-help/issues/283)

> state: **closed** opened by: **** on: ****

I want to use a variable inside a &#x60;for&#x60; loop, like this:

&#x60;&#x60;&#x60;
{% capture sidebar %}
site.data.sidebar_fe.entries
{% endcapture %}

{% for entry in sidebar %}
&#x60;&#x60;&#x60;
Basically here I want to set up a variable called &#x60;sidebar&#x60; that really refers to &#x60;site.data.sidebar.fe.entries&#x60;.

However, this doesn&#x27;t work. Any ideas why? I thought &#x60;capture&#x60; was how I set up a variable. 

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/283#issuecomment-82429381) on: **Tuesday, March 17, 2015**

You might want to try using &#x60;assign&#x60; instead of &#x60;capture&#x60;
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/283#issuecomment-82433629) on: **Tuesday, March 17, 2015**

Yeah, &#x60;assign&#x60; is what you want. This is what I do:

&#x60;&#x60;&#x60;liquid
{% assign page-lang = site.data.locales[page.lang] %}
&#x60;&#x60;&#x60;

&#x60;page.lang&#x60; is either &#x60;de&#x60; or &#x60;en&#x60; and I have a YAML file in &#x60;_data&#x60; that has a structure like this:

&#x60;&#x60;&#x60;yaml
de:
  whatever: &#x27;Ding&#x27;

en:
  whatever: &#x27;thing&#x27;
&#x60;&#x60;&#x60;
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/283#issuecomment-82436933) on: **Tuesday, March 17, 2015**

Awesome! Thanks, that does work. I appreciate your quick response. 

Here&#x27;s the code I used. Basically I have an audience value defined in my config file. Depending on the audience, different sidebars are called:

&#x60;&#x60;&#x60;
{% if site.audience == &quot;fe&quot; %}
{% assign sidebar = site.data.sidebar_fe.entries %}

{% elsif site.audience == &quot;customer&quot; %}
{% assign sidebar = site.data.sidebar_customer.entries %}
{% endif %}

{% for entry in sidebar %} ...
&#x60;&#x60;&#x60;

Thanks for your help.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/283#issuecomment-82437145) on: **Tuesday, March 17, 2015**

:boom: :metal: 
