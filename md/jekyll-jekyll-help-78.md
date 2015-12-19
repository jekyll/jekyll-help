# [Jekyll not outputting HTML](https://github.com/jekyll/jekyll-help/issues/78)

> state: **closed** opened by: **** on: ****

Since upgrading tot he latest version of Jekyll, I&#x27;m getting some really broken output. The output seems to be partial HTML - it uses tags like &#x60;&lt;h1&gt;&#x60;, &#x60;&lt;p&gt;&#x60; and &#x60;&lt;div&gt;&#x60; but doesn&#x27;t include &#x60;&lt;doctype&gt;&#x60; or &#x60;&lt;html&gt;&#x60; tags, so its not rendered correctly in the browser. It also seems to be ignoring supplied includes such as the sidebar and css files. The output index.html looks like this:

&#x60;&#x60;&#x60;html
&lt;h1 id=&quot;sdk-for-android-applications&quot;&gt;SDK for Android Applications&lt;/h1&gt;

&lt;hr /&gt;

&lt;p&gt;This guide will explain the functionality provided by the ForeSee cxReplay
for Mobile Software Development Kit (SDK) and provide some information on
the concepts it leverages.&lt;/p&gt;

&lt;ul&gt;
  &lt;li&gt;See the &lt;a href=&quot;getting-started/quick-start.html&quot;&gt;Quick Start&lt;/a&gt; to get going&lt;/li&gt;
  &lt;li&gt;See the &lt;a href=&quot;apidocs/index.html&quot;&gt;API&lt;/a&gt; docs for more detailed information about the methods
described in this guide.&lt;/li&gt;
  &lt;li&gt;See the examples provided with the SDK to help get you going.&lt;/li&gt;
&lt;/ul&gt;
&#x60;&#x60;&#x60;

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/78#issuecomment-47249210) on: **Thursday, June 26, 2014**

@sleeke This looks like HTML that has been processed by a Markdown renderer only.

You need to include a &#x60;layout&#x60; value in the [Front-matter](http://jekyllrb.com/docs/frontmatter/) that specifies which layout to use. Here&#x27;s an example of how to set this up:

*_layouts/post.html*

&#x60;&#x60;&#x60;html
&lt;!DOCTYPE HTML&gt;
&lt;html&gt;
&lt;head&gt;
  &lt;title&gt;Jekyll - {{ page.title }}&lt;/title&gt;
  {% include head.html %}
&lt;/head&gt;
&lt;body&gt;
  {{ content }}
&lt;/body&gt;
&lt;/html&gt;
&#x60;&#x60;&#x60;

*_includes/head.html*

&#x60;&#x60;&#x60;html
  &lt;meta charset=&quot;UTF-8&quot;&gt;
  &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1, maximum-scale=1&quot;&gt;
&#x60;&#x60;&#x60;

*_posts/a-cool-guy.md*

&#x60;&#x60;&#x60;
---
layout: post
---

# I&#x27;m a post!

This is a paragraph with some content.
&#x60;&#x60;&#x60;

---

There are three things in here that are worth knowing about, especially in the &#x60;_layouts/post.html&#x60; file

* &#x60;{{ page.title }}&#x60; - This is a Liquid tag that will render the variable that was defined in the Front-matter. &#x60;title&#x60; is special; Jekyll will try to automatically populate this variable based on the name of the post file. You can override that by explicitly setting it in the Front-matter.
* &#x60;{% include head.html %}&#x60; - This is another Liquid tag, but this one grabs the content from the file specified located in the &#x60;_includes&#x60; directory. This makes it easy to re-use HTML that should appear in multiple layouts.
* &#x60;{{ content }}&#x60; - This is the rendered Markdown from your post. This will be the HTML that you gave in your original post.

Hopefully this helps! If you have any other questions, feel free to hit us up!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/78#issuecomment-48134110) on: **Sunday, July 06, 2014**

I&#x27;m gonna go ahead and close this one. If you have any other questions about this topic, feel free to add a comment and @mention me.
