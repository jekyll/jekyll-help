# [Redcarpet not found on linux](https://github.com/jekyll/jekyll-help/issues/67)

> state: **closed** opened by: **** on: ****

&#x60;redcarpet&#x60; is installed and in the path.

&#x60;&#x60;&#x60;
$ redcarpet --version
Redcarpet 3.1.2

$ ruby --version
ruby 2.1.2p95 (2014-05-08 revision 45877) [i686-linux]

$ jekyll --version
jekyll 2.0.3
&#x60;&#x60;&#x60;
But when I try to build the site, I get:

&#x60;&#x60;&#x60;
$ jekyll build --trace
Configuration file: /home/fernando/develop/projs/fernandobasso.github.io/_config.yml
            Source: /home/fernando/develop/projs/fernandobasso.github.io
       Destination: /home/fernando/develop/projs/fernandobasso.github.io/_site
      Generating... 
You are missing a library required for Markdown. Please run:
  $ [sudo] gem install redcarpet
  Conversion error: There was an error converting &#x27;_posts/2014-03-10-welcome-to-jekyll.markdown/#excerpt&#x27;.
             ERROR: YOUR SITE COULD NOT BE BUILT:
                    ------------------------------------
                    Missing dependency: redcarpet
&#x60;&#x60;&#x60;

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/67#issuecomment-45121275) on: **Wednesday, June 04, 2014**

Are you using a Ruby version manager?
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/67#issuecomment-45134832) on: **Wednesday, June 04, 2014**

No. I am not. Anyway. I tried creating a new fresh site, and it worked. I saw that the new &#x60;_config.yml&#x60; had some more stuff than the old one. I added &#x60;_markdown: kramdown&#x60; to the config file on the &quot;old&quot; site and it worked.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/67#issuecomment-45135880) on: **Wednesday, June 04, 2014**

There&#x27;s also a bug presently where if you are using the rouge interpreter but the gem isn&#x27;t installed, Jekyll will tell you it&#x27;s because Redcarpet isn&#x27;t installed, when it&#x27;s actually just that rouge isn&#x27;t installed.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/67#issuecomment-45136249) on: **Wednesday, June 04, 2014**

Yeah, installing rouge solved. Thanks.


On Wed, Jun 4, 2014 at 3:58 PM, Parker Moore &lt;notifications@github.com&gt;
wrote:

&gt; There&#x27;s also a bug presently where if you are using the rouge interpreter
&gt; but the gem isn&#x27;t installed, Jekyll will tell you it&#x27;s because Redcarpet
&gt; isn&#x27;t installed, when it&#x27;s actually just that rouge isn&#x27;t installed.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/67#issuecomment-45135880&gt;.
&gt;
