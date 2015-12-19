# [ Error:  undefined method &#x60;[]&#x27; for nil:NilClass](https://github.com/jekyll/jekyll-help/issues/144)

> state: **closed** opened by: **** on: ****

(First time submitting this here so not sure if I&#x27;m doing this right or need more details/etc.)

If I have any post in my _posts folder, I get the error when I try to run &#x60;jekyll serve&#x60;. If I delete the post (and the _posts folder is then empty), the site builds fine and I can view it locally.

&#x60;&#x60;&#x60;
Conversion error: There was an error converting &#x27;_posts/2014-09-01-Making-me-mad.md&#x27;.
jekyll 2.2.0 | Error:  undefined method &#x60;[]&#x27; for nil:NilClass
&#x60;&#x60;&#x60;

Content in the **post** is simply:

&#x60;&#x60;&#x60;
---
Title: &quot;Making Me Mad&quot;
Layout: post
date: September 1, 2014
---

Here&#x27;s some copy for the post that&#x27;s making me mad.
&#x60;&#x60;&#x60;

Content in my **_config.yml file**:

&#x60;&#x60;&#x60;
# Site settings
title: &quot;pdcst blg&quot;

# Handling Reading
# include:      [&quot;.htaccess&quot;]
exclude:      []
markdown_ext: &quot;markdown,mkdown,mkdn,mkd,md&quot;

# Build settings
highlighter: rouge
markdown: rdiscount
rdiscount:
extensions: [smart]
permalink: /articles/:title
excerpt_separator: &quot;&quot;
# paginate: 3
&#x60;&#x60;&#x60;


### Comments

---
> from: [**iChris**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54649657) on: **Friday, September 05, 2014**

Suggested [via Twitter](https://twitter.com/davatron5000/status/507929751307042816) to lowercase my YAML but still gives the same error.

&#x60;&#x60;&#x60;
---
title: &quot;making me mad&quot;
layout: post
date: september 1, 2014
---

Here&#x27;s some copy for the post that&#x27;s making me mad.
&#x60;&#x60;&#x60;

Same error:

&#x60;&#x60;&#x60;
Conversion error: There was an error converting &#x27;_posts/2014-09-01-Making-me-mad.md&#x27;.
jekyll 2.2.0 | Error:  undefined method &#x60;[]&#x27; for nil:NilClass
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54654043) on: **Friday, September 05, 2014**

What&#x27;s the output if you run jekyll with &#x60;--trace&#x60;?
---
> from: [**iChris**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54654705) on: **Friday, September 05, 2014**

**Update** if I comment out:

&#x60;&#x60;&#x60;
# rdiscount:
# extensions: [smart]
&#x60;&#x60;&#x60;

then it builds no problem. I was following another tutorial and put that on there but don&#x27;t really know why? :) 
---
> from: [**iChris**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54657675) on: **Friday, September 05, 2014**

I think I can close this now? I&#x27;ll hit the close and comment button and hopefully nothing explodes. Thanks for your attempts to help me @parkr. I knew it was likely a PEBKAC issue of some sort.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54662213) on: **Friday, September 05, 2014**

&gt; &#x60;&#x60;&#x60;yaml
&gt; # rdiscount:
&gt; # extensions: [smart]
&gt; &#x60;&#x60;&#x60;

I think you wanted:

&#x60;&#x60;&#x60;yaml
rdiscount:
  extensions: [smart] # i&#x27;m indented
&#x60;&#x60;&#x60;

Not sure if your commenting reflects the indention or not. :+1: I also haven&#x27;t tested that so maybe it won&#x27;t work.  
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54679803) on: **Friday, September 05, 2014**

@parkr Is on to something there. You might also try removing the &#x60;exclude: []&#x60; line from your &#x60;_config.yml&#x60; file.

PS: You don&#x27;t have to specify a date string in your post&#x27;s front-matter like that. Liquid has all kinds of [fancy date filters](http://docs.shopify.com/themes/liquid-documentation/filters/additional-filters#date) that allow you to format the date that is automatically generated via the file name.

For example, if you have a post that has a file name of &#x60;2014-09-01-Making-me-mad.md&#x60;, Jekyll will automatically send a date value into Liquid for you to format. The format that you indicated in your example would look like this in your layout: &#x60;{{ page.date | date: &quot;%B %d, %Y&quot; }}&#x60;
---
> from: [**iChris**](https://github.com/jekyll/jekyll-help/issues/144#issuecomment-54690017) on: **Friday, September 05, 2014**

Ah yes. Thanks @troyswanson. Slowly figuring this stuff out.

I&#x27;ve got CNAME issues (or just patience issues) otherwise you could [see the site here](http://blog.pdcst.ca). I&#x27;m pretty sure it&#x27;s just taking time to propagate. The site worked fine on ichris.github.io before I changed the CNAME file.
