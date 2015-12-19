# [Invalid CSS when you include a page inside a page](https://github.com/jekyll/jekyll-help/issues/264)

> state: **closed** opened by: **** on: ****

When you include a page inside another page, the frontmatter gets included. This frontmatter is usually contained inside p tags, except when you have a tags or layout parameter in the frontmatter. In that case, this is what the HTML looks like: 

&#x60;&#x60;&#x60;
&lt;p&gt;layout: page
title: Endpoints
permalink: /sampleendpoint/&lt;/p&gt;

&lt;h2 id=&quot;tags:-implementation&quot;&gt;tags: Implementation&lt;/h2&gt;
&#x60;&#x60;&#x60;

In trying to hide this from the output, it&#x27;s not possible because the id &#x60;tags:-implementation&#x60; isn&#x27;t valid CSS. Colons aren&#x27;t allowed in CSS selectors. As a result, I can&#x27;t hide it programmatically from the page by using display none for that CSS element. (Right?)

Why am I including a page inside a page? I need a way to get my site to people offline. I&#x27;m using Jekyll for a user guide, and some people prefer PDFs. My approach is to create a page that includes all other pages, push this master compilation into a print layout, and then use a print stylesheet (and maybe Prince XML) to render it to PDF. But I&#x27;m getting hung-up with the h2 headings that contain information from the frontmatter about tags and layout. 

How can I work around this to hide the frontmatter?

As a side note, I don&#x27;t understand why Jekyll&#x27;s site output requires a server to view the files. If they&#x27;re all static files, why can&#x27;t you just zip the site file to someone so they can view it? I read through the techniques listed in this [StackOverflow thread](http://stackoverflow.com/questions/26778329/running-jekyll-generated-files-without-jekyll-local-server), and it confuses me why Jekyll was architected this way.


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/264#issuecomment-72374889) on: **Sunday, February 01, 2015**

On this:

&gt; As a side note, I don&#x27;t understand why Jekyll&#x27;s site output requires a server to view the files.

Technically, every HTML site is static on usable offline and online without links or ressource being not available. The culprit here is how you&#x27;re setting links (&#x60;href&#x60; or &#x60;src&#x60;). If you set absolute links or file relative links, you won&#x27;t have problems. It gets a bit more tricky when you introduce the principle of __root directory__. Links to &#x60;/index.html&#x60; are relative to the root, not the file itself. This usually leads to problems when try to view local files in a browser. It looks for its ressources at the wrong places, because it has no way to determine the correct root directoy (the project root that is in most cases).
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/264#issuecomment-72375214) on: **Sunday, February 01, 2015**

Thanks Kleinfreund. I appreciate your response. I tried creating a jekyll site using relative links (no baseurl), but it didn&#x27;t work when viewed outside a local server. If it&#x27;s possible, would you mind creating a sample jekyll site that demonstrates how it can be viewed without a local server?



---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/264#issuecomment-72375334) on: **Sunday, February 01, 2015**

I gave up on the includes and found a better approach for consolidating pages. I&#x27;m using a for loop to loop through pages, like this:

&#x60;&#x60;&#x60;
---
title: Printed manual
permalink: /print/
layout: none
---

{% for p in site.pages %}
{% if p.permalink == &quot;/troubleshooting_hmac/&quot; %}
&lt;h1&gt;{{p.title}}&lt;/h1&gt;
{{p.content}}
{% endif %}
{% endfor %}

{% for p in site.pages %}
{% if p.permalink == &quot;/release_notes/&quot; %}
&lt;h1&gt;{{p.title}}&lt;/h1&gt;
{{p.content}}
{% endif %}
{% endfor %}

{% for p in site.pages %}
{% if p.permalink == &quot;/overview/&quot; %}
&lt;h1&gt;{{p.title}}&lt;/h1&gt;
{{p.content}}
{% endif %}
{% endfor %}
&#x60;&#x60;&#x60;

The layout for this page is set to none, but in my config file, the default layout for pages is set to a print layout. If you include a page layout for this consolidation page, you end up with two instances of the content -- one as the content gets pushed into an existing layout, and another as it gets pushed into the layout specified by the configuration file.



---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/264#issuecomment-72383336) on: **Sunday, February 01, 2015**

I just realized something really simple. The whole reason I was exploring this thread is because I needed a way to deliver a Jekyll site to users who aren&#x27;t online. After figuring out a workflow to PDF, I realized I could just download the site using a utility called Sitesucker (there are others similar) to get all the resources offline. Then I can send a user the zip file. Duh!!!!! So much easier than trying to run loops to get posts and create a PDF out of them, etc. 
