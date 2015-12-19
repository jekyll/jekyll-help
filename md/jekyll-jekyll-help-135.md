# [Deploying to Heroku](https://github.com/jekyll/jekyll-help/issues/135)

> state: **open** opened by: **** on: ****

Hello,
I&#x27;ve had a hell of a time trying to get Jekyll to run on Heroku. I followed the link for getting it to run with Rack 
http://blog.crowdint.com/2010/08/02/instant-blog-using-jekyll-and-heroku.html but it didn&#x27;t work and the article is over four years old. Right now I&#x27;m getting &quot;Push rejected, no Cedar-supported app detected&quot; but at one point I could push to it but the logs said it crashed without giving a useful reason.

I don&#x27;t know Rack or Ruby so this has been a very frustrating experience. I&#x27;ve tried several different tutorials on getting Jekyll to run on heroku but nothing has worked. Eventually I need to get this running on EC2 as well.

Even the static site files by themselves would be fine really.

My Procfile:
&#x60;&#x60;&#x60;
web: jekyll serve -p $PORT
&#x60;&#x60;&#x60;
My config.ru
&#x60;&#x60;&#x60;
require &quot;rack/jekyll&quot;

run Rack::Jekyll.new
&#x60;&#x60;&#x60;
My .gems
&#x60;&#x60;&#x60;
rack-jekyll
&#x60;&#x60;&#x60;
Full repo is here https://github.com/AndyCErnst/beat-the-banker-jekyll

### Comments

---
> from: [**AndyCErnst**](https://github.com/jekyll/jekyll-help/issues/135#issuecomment-52878899) on: **Wednesday, August 20, 2014**

Okay, more info. Adding a Gemfile allowed it to deploy to Heroku (but it doesn&#x27;t run). They really shouldn&#x27;t assume there is a Gemfile in Jekyll because it doesn&#x27;t need one to run locally. According to the heroku logs the problem now is: &quot;Error reading file /app/vendor/bundle/ruby/2.1.0/gems/jenkyll-1.5.1/test/fixtures/broken_front_matter3.erb: (&lt;unknown&gt;): did not find expected node content while parsing a flow node at line 3 column 1.&quot; 

---
> from: [**AndyCErnst**](https://github.com/jekyll/jekyll-help/issues/135#issuecomment-53171690) on: **Saturday, August 23, 2014**

Ah, finally figured it out. The parts I feel should really be mentioned on the site/article are:
_config.yml file must have:
&#x60;&#x60;&#x60;
excludes: [&quot;vendor&quot;]
&#x60;&#x60;&#x60;
The Procfile must have:
&#x60;&#x60;&#x60;
web: jekyll serve -P $PORT
&#x60;&#x60;&#x60;
There has to be a Gemfile with at least:
&#x60;&#x60;&#x60;
source &quot;https://rubygems.org&quot;
ruby &quot;2.1.1&quot;
gem &quot;rack-jekyll&quot;, &quot;~&gt; 0.4.1&quot;
&#x60;&#x60;&#x60;
