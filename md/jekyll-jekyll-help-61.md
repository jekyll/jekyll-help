# [Tumblr to Jekyll Migration, Images not downloading](https://github.com/jekyll/jekyll-help/issues/61)

> state: **closed** opened by: **** on: ****

Using the migration code found here: http://import.jekyllrb.com/docs/tumblr/ 

&#x60;&#x60;&#x60;
$ ruby -rubygems -e &#x27;require &quot;jekyll-import&quot;;
    JekyllImport::Importers::Tumblr.run({
      &quot;url&quot;            =&gt; &quot;http://myblog.tumblr.com&quot;,
      &quot;format&quot;         =&gt; &quot;html&quot;, # or &quot;md&quot;
      &quot;grab_images&quot;    =&gt; true,  # whether to download images as well.
      &quot;add_highlights&quot; =&gt; false,  # whether to wrap code blocks (indented 4 spaces) in a Liquid &quot;highlight&quot; tag
      &quot;rewrite_urls&quot;   =&gt; false   # whether to write pages that redirect from the old Tumblr paths to the new Jekyll paths
    })&#x27;
&#x60;&#x60;&#x60;

This is only downloading images from text posts, not any images from photo posts. 

(There is a decent chance I&#x27;m missing something obvious, I&#x27;m new to jekyll :smile:)


### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/61#issuecomment-44763804) on: **Saturday, May 31, 2014**

Yes! The current implementation is simple and limited (https://github.com/jekyll/jekyll-import/pull/134). If you have the time to expand it, I&#x27;d love the help!
