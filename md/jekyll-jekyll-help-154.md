# [Mangled quotation marks](https://github.com/jekyll/jekyll-help/issues/154)

> state: **open** opened by: **** on: ****

Hi,

I&#x27;m just getting started with Jekyll. I&#x27;ve run into an issue where my quotation marks are being mistransformed. Presumably this is an encoding issue, but I&#x27;ve explicitly set my encoding to UTF-8 in my *_config.yml*:
&#x60;&#x60;&#x60;
exclude: [ Gemfile ]
encoding: &quot;UTF-8&quot;

defaults:
  -
    scope:
      path: &quot;&quot;          # all files
      type: &quot;posts&quot;     # only posts though
    values:
      layout: &quot;default&quot;
&#x60;&#x60;&#x60;
I&#x27;ve also double-checked that the file in question is being saved as UTF-8. I tried two different text editors to be sure.

Here&#x27;s the input file (markdown):
&#x60;&#x60;&#x60;
---
title: A test
---

Here are some &#x27;single quotes&#x27;. Here are some &quot;double quotes&quot;.
&#x60;&#x60;&#x60;
And here is what I see when I browse to the generated site:

![screen shot 2014-09-19 at 12 15 54 pm](https://cloud.githubusercontent.com/assets/1901832/4330107/2de6e124-3fa7-11e4-9f79-c57889ad3dc0.png)

Can anyone tell me how to rectify this?

Thanks


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/154#issuecomment-56143250) on: **Friday, September 19, 2014**

Can you show us what you have in your &#x60;_layouts/default.html&#x60;?

There has to be a line like this:

&#x60;&#x60;&#x60;html
&lt;meta charset=&quot;utf-8&quot;&gt;
&#x60;&#x60;&#x60;

If not, add it between the &#x60;&lt;head&gt;&lt;/head&gt;&#x60; tags in the HTML.
