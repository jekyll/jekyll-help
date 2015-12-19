# [Front matter defaults for collection&#x27;s documents in subfolders](https://github.com/jekyll/jekyll-help/issues/237)

> state: **open** opened by: **** on: ****

Hi there,

I would like to organize collection&#x27;s documents in subfolders and assign them different categories through front matter defaults. I have the following structure:

&#x60;&#x60;&#x60;&#x60;
_kb
    - category1
        - article1.md
    - category2
        - article2.md
&#x60;&#x60;&#x60;&#x60;

 In my &#x60;_config.yml&#x60; file I have:

&#x60;&#x60;&#x60;&#x60;
# Collections
collections:
  kb:
    output: true
    permalink: /kb/:name/

# Defaults
defaults:
  -
    scope:
      path: &quot;_kb/category1/&quot;
      type: &quot;kb&quot;
    values:
      category: &quot;category1&quot;
  -
    scope:
      path: &quot;_kb/category2/&quot;
      type: &quot;kb&quot;
    values:
      category: &quot;category2&quot;
&#x60;&#x60;&#x60;&#x60;

But this doesn&#x27;t work. Any idea how I can assign different categories to documents in different subfolders through front-matter defaults?

### Comments

---
> from: [**pearsonca**](https://github.com/jekyll/jekyll-help/issues/237#issuecomment-95775533) on: **Thursday, April 23, 2015**

I have a similar problem.

I&#x27;ve tried &#x60;path: &quot;category&quot;&#x60; (with &#x60;type: &quot;collection&quot;&#x60;), &#x60;collection/category&#x60;, &#x60;_collection/category&#x60;, and assorted other combinations.  No luck yet!
