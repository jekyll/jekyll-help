# [After moving from GitHub Pages to S3 &amp; Cloudfront, Google Analytics No Longer Firing](https://github.com/jekyll/jekyll-help/issues/164)

> state: **closed** opened by: **** on: ****

Google Analytics tag is no longer showing up in source code on vincentbarr.com. Config file: https://gist.github.com/vincentbarr/032012e103eee4208f98

Can you please help?*

*Please let me know if there is a better way to demonstrate issues going forward. 

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/164#issuecomment-59591841) on: **Friday, October 17, 2014**

Do you manage to solve the issue?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/164#issuecomment-61438896) on: **Sunday, November 02, 2014**

It&#x27;s a pretty convoluted way render this, but &#x60;{{ site.JB.analytics.google.tracking_id }}&#x60; should work fine.

&#x60;&#x60;&#x60;html
&lt;!-- GOOGLE ANALYTICS --&gt;
  &lt;script&gt;
    (function(i,s,o,g,r,a,m){i[&#x27;GoogleAnalyticsObject&#x27;]=r;i[r]=i[r]||function(){
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,&#x27;script&#x27;,&#x27;//www.google-analytics.com/analytics.js&#x27;,&#x27;ga&#x27;);
    ga(&#x27;create&#x27;, &#x27;{{ site.JB.analytics.google.tracking_id }}&#x27;, &#x27;auto&#x27;);
    ga(&#x27;send&#x27;, &#x27;pageview&#x27;);
  &lt;/script&gt;
&#x60;&#x60;&#x60;
