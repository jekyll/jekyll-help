# [Working, official-unofficial Atom feed code](https://github.com/jekyll/jekyll-help/issues/7)

> state: **open** opened by: **** on: ****

I currently use [this](https://github.com/ndarville/ndarville.github.io/blob/master/feed/index.xml) feed code for [my blog](http://ndarville.github.io/). Unfortunately, it doesn’t seem to work in readers (ndarville/ndarville.github.io#11).

Finding the right feed code is a terrible experience I wouldn’t wish upon my worst enemy, and I recall how long it took me to get it to work for the static Blogofile CMS.

Can we establish a working official-unofficial Jekyll feed code that people can use, and which can be improved upon in one central place, should any new bugs be found?

## Decisions

- [x] &#x60;{{ site.baseurl }}&#x60; vs. &#x60;{{ site.url }}&#x60; w/r/t SSL
- [x] &#x60;xml_escape&#x60; vs. &#x60;CDATA&#x60;
- [x] Filename: &#x60;/atom/index.xml&#x60;, &#x60;atom.xml&#x60;, &#x60;feed.atom&#x60;, etc.

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39351056) on: **Wednesday, April 02, 2014**

This is a fantastic idea! @benbalter or @imathis might have an idea. I&#x27;d be down to write a plugin which handles this for us.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39354594) on: **Wednesday, April 02, 2014**

I think @mdo has a good example of a feed: https://github.com/poole/poole/blob/master/atom.xml

this works for me, realizing that much of this would vary from site to site:
&#x60;&#x60;&#x60;

&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;feed xmlns=&quot;http://www.w3.org/2005/Atom&quot;&gt;

 &lt;title&gt;{{ site.sitename }}&lt;/title&gt;
 &lt;link href=&quot;{{ site.baseurl }}/atom.xml&quot; rel=&quot;self&quot;/&gt;
 &lt;link href=&quot;{{ site.baseurl }}/&quot;/&gt;
 &lt;updated&gt;{{ site.time | date_to_xmlschema }}&lt;/updated&gt;
 &lt;id&gt;{{ site.baseurl }}/&lt;/id&gt;
 &lt;author&gt;
   &lt;name&gt;{{ site.author }}&lt;/name&gt;
   {% if  site.author.email  %}&lt;email&gt;{{ site.author.email }}&lt;/email&gt;{% endif %}
 &lt;/author&gt;

 {% for post in site.posts %}
 &lt;entry&gt;
   &lt;title&gt;{{ post.title }}&lt;/title&gt;
   &lt;link href=&quot;{{ site.baseurl }}{{ post.url }}&quot;/&gt;
   &lt;updated&gt;{{ post.date | date_to_xmlschema }}&lt;/updated&gt;
   &lt;id&gt;{{ site.baseurl }}{{ post.id }}&lt;/id&gt;
   &lt;content type=&quot;html&quot;&gt;{{ post.content | xml_escape }}&lt;/content&gt;
 &lt;/entry&gt;
 {% endfor %}

&lt;/feed&gt;
&#x60;&#x60;&#x60;


---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39355017) on: **Wednesday, April 02, 2014**

Is there some kind of feed validation available? My method of seeing whether a handful of feed readers will eat my XML probably isn’t the smartest way of testing, especially if someone breaks farther down the road.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39355137) on: **Wednesday, April 02, 2014**

This: http://feedvalidator.org/
But beware - it will throw warnings based on content.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39678321) on: **Sunday, April 06, 2014**

I implemented mdo’s script, but I still get some errors: &lt;http://feedvalidator.org/check.cgi?url=https%3A%2F%2Fndarville.github.io%2Ffeed%2Findex.xml&gt;.

It looks like feed readers still won’t eat it as well, so maybe we need something beyond mdo’s script. Maybe wrapping it in &#x60;CDATA&#x60; will help. I’ll try to look into it tomorrow.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39678672) on: **Sunday, April 06, 2014**

Adding [this](https://github.com/ndarville/ndarville.github.io/commit/ba4dad36d2eed113e3cfea68532a656f87eb2151) change fixed it.

So if we add that part, we can canonize mdo’s as the unofficial-official Jekyll feed.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39681234) on: **Sunday, April 06, 2014**

except I think it should be {{site.baseurl}} rather than just {{site.url}} http://jekyllrb.com/docs/upgrading/#baseurl
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-39701445) on: **Monday, April 07, 2014**

Hmm. I am using a [hardcoded, absolute URL](https://github.com/ndarville/ndarville.github.io/commit/168acc112ad418ef956697c02cf42d26b9604ab9) to ensure I direct users to the SSL version of my website—although I understand this is something specific to my set-up. Won’t &#x60;{{ site.baseurl }}&#x60; just redirect to &#x60;/feed/index.xml&#x60;?

Either way, the change is to add an &#x60;xml_escape&#x60; filter; it’s also important to figure out which URL is better for SSL nuts, though, albeit a separate discussion. We should still figure it out here, though.

&#x60;- [ ] Decide on &#x60;{{ site.baseurl }}&#x60; vs. &#x60;{{ site.url }}&#x60; w/r/t SSL&#x60;
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40368472) on: **Monday, April 14, 2014**

For what it&#x27;s worth, here&#x27;s my implementation of an atom feed with only one warning using the validator:

&#x60;&#x60;&#x60;
---
---
&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;feed xmlns=&quot;http://www.w3.org/2005/Atom&quot;&gt;
  &lt;title&gt;{{ site.title }}&lt;/title&gt;
  &lt;link href=&quot;http://{{ site.domain }}/&quot;/&gt;
  &lt;link href=&quot;http://{{ site.domain }}/atom.xml&quot; rel=&quot;self&quot;/&gt;
  &lt;updated&gt;{{ site.time | date_to_xmlschema }}&lt;/updated&gt;
  &lt;id&gt;http://{{ site.domain }}/&lt;/id&gt;
  &lt;author&gt;
    &lt;name&gt;{{ site.author.name }}&lt;/name&gt;
    &lt;email&gt;{{ site.author.email }}&lt;/email&gt;
  &lt;/author&gt;

  {% for post in site.posts %}
  &lt;entry&gt;
    &lt;title&gt;{{ post.title }}&lt;/title&gt;
    &lt;link href=&quot;http://{{ site.domain }}{{ post.url }}?utm_source=atom&amp;amp;utm_medium=rss&amp;amp;utm_campaign=atom&quot;/&gt;
    &lt;updated&gt;{{ post.date | date_to_xmlschema }}&lt;/updated&gt;
    &lt;id&gt;http://{{ site.domain }}{{ post.id }}&lt;/id&gt;
    &lt;content type=&quot;html&quot;&gt;{{ post.content | xml_escape }}&lt;/content&gt;
  &lt;/entry&gt;
  {% endfor %}
&lt;/feed&gt;
&#x60;&#x60;&#x60;

This is combined with a few fairly straightforward settings in &#x60;config.yml&#x60;:

&#x60;&#x60;&#x60;
title: Matthew Scharley
domain: matt.scharley.me
author:
  name: Matthew Scharley
  email: matt.scharley@gmail.com
&#x60;&#x60;&#x60;

http://feedvalidator.org/check.cgi?url=http%3A%2F%2Fmatt.scharley.me%2Fatom.xml
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40505897) on: **Tuesday, April 15, 2014**

After doing some testing, it seems CDATA is equivalent to &#x60;xml_escape&#x60;, which makes sense now that I think about it. That said, I&#x27;ve taken to using xml_escape for the &#x60;&lt;title&gt;&#x60; and CDATA for the &#x60;&lt;content&gt;&#x60; as it will negate the need to escape to XML entities, making the content both shorter (and hence smaller filesize) as well as easier to read the XML directly.

No opinion on baseurl vs url. baseurl seems useless at the moment as you can&#x27;t include it if it is &#x60;/&#x60;,  which is the default and the case in 90% of cases.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40594745) on: **Wednesday, April 16, 2014**

I don&#x27;t find baseurl useless. I set my production baseurl in config and then flag my local builds with &#x60;&#x60;&#x60;-baseurl &quot;&quot;&#x60;&#x60;&#x60; so to set url and baseurl would be redundant being that baseurl has meaning and url, to my knowledge does not.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40598717) on: **Wednesday, April 16, 2014**

An empty baseurl is actually a good idea, but it&#x27;s not what&#x27;s reflected in the documentation which recommends doing what you do and setting it to &#x60;/&#x60; locally. In that case, it&#x27;s probably a good idea.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40598764) on: **Wednesday, April 16, 2014**

That also still doesn&#x27;t take care of the default being &#x60;/&#x60; though, so it would break on default setups.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40608374) on: **Wednesday, April 16, 2014**

I&#x27;ve transitioned to using [GitHub Pages Metadata](https://help.github.com/articles/repository-metadata-on-github-pages) instead of using something like &#x60;baseurl&#x60; in the configuration file.

Here&#x27;s an example:

&#x60;&#x60;&#x60;html
&lt;link rel=&quot;icon&quot; href=&quot;{{ site.github.url }}/img/favicon-16x16.png&quot; sizes=&quot;16x16&quot; type=&quot;image/png&quot;&gt;
&lt;link rel=&quot;icon&quot; href=&quot;{{ site.github.url }}/img/favicon-32x32.png&quot; sizes=&quot;32x32&quot; type=&quot;image/png&quot;&gt;
&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ site.github.url }}/css/cornerstone.css&quot;&gt;
&#x60;&#x60;&#x60;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40608786) on: **Wednesday, April 16, 2014**

Very cool, @troyswanson 
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40988478) on: **Monday, April 21, 2014**

According to [RFC4287](http://tools.ietf.org/html/rfc4287) you can use the file extension &#x60;atom&#x60; for atom feeds:

&#x60;&#x60;&#x60;xml
&lt;link rel=&quot;self&quot;
  type=&quot;application/atom+xml&quot;
  href=&quot;http://example.org/feed.atom&quot;/&gt;
&#x60;&#x60;&#x60;
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-40988821) on: **Monday, April 21, 2014**

Yep, it&#x27;s an official mime type
On 22/04/2014 9:27 am, &quot;Anatol Broder&quot; &lt;notifications@github.com&gt; wrote:

&gt; According to RFC4287 &lt;http://tools.ietf.org/html/rfc4287&gt; you can use the
&gt; file extension atom for atom feeds:
&gt;
&gt; &lt;link rel=&quot;self&quot;
&gt;   type=&quot;application/atom+xml&quot;
&gt;   href=&quot;http://example.org/feed.atom&quot;/&gt;
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/7#issuecomment-40988478&gt;
&gt; .
&gt;
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41010137) on: **Tuesday, April 22, 2014**

Just to get an idea of where we are, I added some tasks to the top post for the remaining things to hash out, before we decide on a canonical Atom file.

Feel free to add to the list.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41014006) on: **Tuesday, April 22, 2014**

My 2c.

atom.xml.

xml_escape. CDATA will yield slightly smaller file sizes due to less
escaping, but xml_escape will yield more reliability.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41032484) on: **Tuesday, April 22, 2014**

@ndarville I think the first issue is solved by @troyswanson&#x27;s comment about using Github Pages Metadata (site.github.url). I think the use of one of the two (site.url, etc.) may be user preferences, whereas using GH metadata ensures it won&#x27;t break.  SSL might be an issue, but I that&#x27;s an edge case.

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41032601) on: **Tuesday, April 22, 2014**

The issue with Github Pages Metadata is that it relies on people deploying
to Github Pages, which is far from the only way that people are deploying
jekyll websites.


On 22 April 2014 22:17, Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; @ndarville &lt;https://github.com/ndarville&gt; I think the first issue is
&gt; solved by @troyswanson &lt;https://github.com/troyswanson&gt;&#x27;s comment about
&gt; using Github Pages Metadata (site.github.url). I think the use of one of
&gt; the two (site.url, etc.) may be user preferences, whereas using GH metadata
&gt; ensures it won&#x27;t break. SSL might be an issue, but I that&#x27;s an edge case.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/7#issuecomment-41032484&gt;
&gt; .
&gt;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41032663) on: **Tuesday, April 22, 2014**

Of course, good point.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41033266) on: **Tuesday, April 22, 2014**

Out of curiousity, what&#x27;s the argument for using HTTPS with jekyll anyway?
There&#x27;s nothing private in a static site, unless the entire site is, but
that would be the minority of uses I think. Perhaps we&#x27;re overthinking it?

Of course, there&#x27;s always another option that I just thought of. Scheme
relative URL&#x27;s, eg. &#x60;//github.com/jekyll/jekyll&#x60;, which means use either
HTTP or HTTPS depending on how the original document was fetched.


On 22 April 2014 22:19, Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; Of course, good point.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/7#issuecomment-41032663&gt;
&gt; .
&gt;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41081557) on: **Tuesday, April 22, 2014**

From what I see here &#x60;site.url&#x60; doesn’t exist by default. Jekyll expose &#x60;site.baseurl&#x60; only. The [documentation](http://jekyllrb.com/docs/configuration/) lies about it. We can use conditions to find the base:

&#x60;&#x60;&#x60;xml
&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;feed xmlns=&quot;http://www.w3.org/2005/Atom&quot;&gt;

{% if site.url %}
  {% assign base = site.url %}
{% elsif site.github.url %}
  {% assign base = site.github.url %}
{% else %}
  {% assign base = &#x27;&#x27; %}
{% endif %}

&lt;id&gt;{{ base }}{{ page.url }}&lt;/id&gt;

&lt;/feed&gt;
&#x60;&#x60;&#x60;

If one doesn’t use Github, she sets the url in the config file to &#x60;http://example.org&#x60; (no trailing slash).


:unlock: 
My votes:
* &#x60;site.url&#x60;
* &#x60;xml_escape&#x60;
*  &#x60;feed.atom&#x60; (Github uses the atom extension too)


:lock: 
Correction: 

* &#x60;site.github.url&#x60; if &#x60;site.url&#x60; isn&#x27;t set
* &#x60;xml_escape&#x60;
* &#x60;feed.atom&#x60;
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41082900) on: **Tuesday, April 22, 2014**

I like your idea of using conditionals. It means people also don’t have to alter their Atom code according to their set-up for most use cases.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41102198) on: **Tuesday, April 22, 2014**

Why not use &#x60;site.github.url&#x60; if &#x60;site.url&#x60; isn&#x27;t set and &#x60;site.github.url&#x60; is available?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41103151) on: **Tuesday, April 22, 2014**

I vote for:

- &#x60;site.github.url&#x60; if &#x60;site.url&#x60; isn&#x27;t set.
- &#x60;xml_escape&#x60;
- &#x60;feed.atom&#x60;

Same as @penibelst but with @parkr conditional
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41126901) on: **Tuesday, April 22, 2014**

@parkr This is exactly what I want. Thank you @albertogg for putting everything together.

I second @albertogg’s selection.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41152962) on: **Wednesday, April 23, 2014**

I&#x27;m still not clear on why anyone would use site.url instead of site.baseurl. Can anyone enlighten me?
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/7#issuecomment-41168033) on: **Wednesday, April 23, 2014**

I’m down with @albertogg and @parkr as well.
