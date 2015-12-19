# [Question about the &quot;where&quot; filter](https://github.com/jekyll/jekyll-help/issues/155)

> state: **open** opened by: **** on: ****

Am I crazy, or was it possible at one point to use the &quot;where&quot; filter to exclude items with a particular value, like this:

Documented usage:
&#x60;&#x60;&#x60;
{{ site.members | where:&quot;graduation_year&quot;,&quot;2014&quot; }}
&#x60;&#x60;&#x60;
and not-documented, but I thought once worked:
&#x60;&#x60;&#x60;
{{ site.members | where:&quot;graduation_year&quot;, -&quot;2014&quot; }}
&#x60;&#x60;&#x60;
to pull members who graduated in any year NOT 2014.

It seems to me like I used this once, but it doesn&#x27;t work now.

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/155#issuecomment-56181358) on: **Friday, September 19, 2014**

Just want to add that I know there are other ways to accomplish the same thing, I just think this is efficient and seemed to recall it working at some point.
