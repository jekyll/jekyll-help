# [Issue with Jekyll Mercenary](https://github.com/jekyll/jekyll-help/issues/189)

> state: **open** opened by: **** on: ****

Running Jekyll 2.5.1 with Ruby 2.0 . When I run &#x60;jekyll s&#x60; or &#x60;jekyll w&#x60;
&#x60;&#x60;&#x60;
Error output
/Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.4/lib/mercenary.rb:20:in &#x60;program&#x27;: cannot load such file -- mercenary/program (LoadError)
	from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.1/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

My Gemfile. I&#x27;ve commented out the other gem still getting error.
&#x60;&#x60;&#x60;
source &#x27;https://rubygems.org&#x27;
gem &#x27;jekyll&#x27;
# gem &#x27;jekyll-assets&#x27;
# gem &#x27;autoprefixer-rails&#x27;
# gem &#x27;jekyll-press&#x27;
# gem &#x27;jekyll-sitemap&#x27;
# gem &#x27;jekyll-mentions&#x27;
# gem &#x27;jekyll-tagging&#x27;
# gem &#x27;ruby-oembed&#x27;
# gem &#x27;mercenary&#x27;
&#x60;&#x60;&#x60;
bunder.rb file
&#x60;&#x60;&#x60;
require &#x27;rubygems&#x27;
require &#x27;bundler/setup&#x27;
require &#x27;jekyll&#x27;
# require &#x27;jekyll-assets&#x27;
# require &#x27;autoprefixer-rails&#x27;
# require &#x27;jekyll-press&#x27;
# require &#x27;jekyll-sitemap&#x27;
# require &#x27;jekyll-mentions&#x27;
# require &#x27;jekyll/tagging&#x27;
# require &#x27;oembed&#x27;
&#x60;&#x60;&#x60;

I&#x27;m not too experienced with Ruby, but I noticed in the Gemfile.lock under Jekyll in requires lists       mercenary (~&gt; 0.3.3) and then later on lists &quot; mercenary (0.3.4)&quot; separately.

If I run &#x60;bundle exec jekyll s&#x60; I receive no errors. Maybe it&#x27;s the way I&#x27;m calling / requiring these gems.





### Comments

---
> from: [**ericthor**](https://github.com/jekyll/jekyll-help/issues/189#issuecomment-62319752) on: **Sunday, November 09, 2014**

Witched to  Ruby to 2.1.2 with rvm. No longer receiving errors. Need to update back to latest Jekyll will see if it creates same errors on &#x60;jekyll s&#x60; or &#x60;jekyll serve&#x60; or &#x60;jekyll b&#x60; ...
---
> from: [**ericthor**](https://github.com/jekyll/jekyll-help/issues/189#issuecomment-62320091) on: **Sunday, November 09, 2014**

Running the same setup including other gems using Ruby 2.1.2 and Jekyll 2.51 no error.
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/189#issuecomment-62321438) on: **Sunday, November 09, 2014**

Should be a duplicate of jekyll/jekyll#3084
---
> from: [**sparanoid**](https://github.com/jekyll/jekyll-help/issues/189#issuecomment-62533644) on: **Tuesday, November 11, 2014**

I got the same error when running it with bundler:

https://travis-ci.org/sparanoid/sparanoid.com/builds/40634156

Still receiving errors using Ruby 2.1.2 with Jekyll 2.5.1.
