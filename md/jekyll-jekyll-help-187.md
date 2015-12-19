# [Jekyll &amp; JavaScript Templates](https://github.com/jekyll/jekyll-help/issues/187)

> state: **closed** opened by: **** on: ****

How can I use JavaScript templating in a Jekyll site? (Angular.js, Handlebars.js, etc.) 

The {{ double brackets }} don&#x27;t work because the Liquid template uses them.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/187#issuecomment-62073474) on: **Thursday, November 06, 2014**

You can use the &#x60;{% raw %}&#x60; tag to prevent Liquid from parsing stuff. Like this:

&#x60;&#x60;&#x60;html
&lt;span&gt;You have {% raw %}{{ unread_count }}{% endraw %} unread messages&lt;/span&gt;
&#x60;&#x60;&#x60;
