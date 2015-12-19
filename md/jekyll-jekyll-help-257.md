# [Liquid assign tag and filters](https://github.com/jekyll/jekyll-help/issues/257)

> state: **open** opened by: **** on: ****

I have a collection **classification**. In a template I want to read one item from the collection based on a front matter &#x60;id&#x60; value.

I thought this is straightforward. But then I noticed (using the Octopress debugger) that after executing

&#x60;&#x60;&#x60;
{% assign item = site.classification | where:&quot;id&quot;,&quot;xyz&quot; %}
{% debug %}
&#x60;&#x60;&#x60;

**item** contains a flattened String representation of the collection item and not the collection item itself (as I would have expected). I noticed the same for the **group_by** filter. It seems that the assign tag in conjunction with these filter flattens the result to a String. It does not seem to be the filter itself. Jekyll&#x27;s where filter seems fine when I look at the source code.

When I do

&#x60;&#x60;&#x60;
{% assign item = site.classification %}
{% debug %}
&#x60;&#x60;&#x60;

**without filter**, the debugger shows me that **item** contains the actual collection as it should. 

What is the expected behavior when applying the filters mentioned here: http://jekyllrb.com/docs/templates/?

### Comments

---
> from: [**mdotasia**](https://github.com/jekyll/jekyll-help/issues/257#issuecomment-71919198) on: **Wednesday, January 28, 2015**

Just read in Liquid docs that filters can be applied only to output tags &#x60;{{ }}&#x60;. Is there any way to select a subset from a collection other than iterating through the collection with a for loop?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/257#issuecomment-73107260) on: **Thursday, February 05, 2015**

Try the &#x60;where&#x60; filter.
