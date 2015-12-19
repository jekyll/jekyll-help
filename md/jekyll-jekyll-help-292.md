# [Using SASS files from gems installed with bundler](https://github.com/jekyll/jekyll-help/issues/292)

> state: **open** opened by: **** on: ****

I&#x27;m wondering if this is possible: I&#x27;m installing the [GitHub](http://github.com) gem by adding it to my &#x60;&#x60;&#x60;Gemfile&#x60;&#x60;&#x60;:

&#x60;&#x60;&#x60;
source &#x27;https://rubygems.org&#x27;
gem &#x27;github-pages&#x27;
gem &#x27;bootstrap-sass&#x27;, &#x27;~&gt; 3.3.4&#x27;
&#x60;&#x60;&#x60;

Then i call &#x60;&#x60;&#x60;bundle install&#x60;&#x60;&#x60;.

But I don&#x27;t get how I now can access the installed SASS files within Jekyll. How do I import the bootstrap SASS file from

&#x60;&#x60;&#x60;
/var/lib/gems/2.1.0/gems/bootstrap-sass-3.3.4.1/assets/stylesheets
&#x60;&#x60;&#x60;

in my local SASS files?

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/292#issuecomment-95733783) on: **Thursday, April 23, 2015**

I *think* you might have to install with bower to use it in Jekyll. Then just add to your config and Jekyll will build it for you http://jekyllrb.com/docs/assets/#sassscss I&#x27;d point you to a repo to look at, but I use bourbon.
---
> from: [**astehlik**](https://github.com/jekyll/jekyll-help/issues/292#issuecomment-95982547) on: **Friday, April 24, 2015**

Thank you for your reply!

Using bower is what I did actually to solve the problem :)

I was just curious if there is a possibility to use the files from the gem.
