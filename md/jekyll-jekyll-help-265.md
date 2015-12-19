# [jekyll-3.0.0.pre.beta](https://github.com/jekyll/jekyll-help/issues/265)

> state: **closed** opened by: **** on: ****

/Library/Ruby/Gems/2.0.0/gems/jekyll-3.0.0.pre.beta1/lib/jekyll.rb:33:in &#x60;&lt;top (required)&gt;&#x27;: undefined method &#x60;error_mode=&#x27; for Liquid::Template:Class (NoMethodError)
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/jekyll-3.0.0.pre.beta1/bin/jekyll:6:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-73399051) on: **Saturday, February 07, 2015**

It&#x27;s possible that Jekyll loaded Liquid 2 instead of Liquid 3. Make sure you&#x27;ve uninstalled Liquid 3 from your system or are using Jekyll with Bundler.
---
> from: [**crearc**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-74764305) on: **Tuesday, February 17, 2015**

I&#x27;m receiving the same error. I use Bundler and I made sure all the packages were up to date. @parkr can you re-open?
---
> from: [**crearc**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-74774151) on: **Tuesday, February 17, 2015**

@frall8 did you find a solution?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-74926911) on: **Wednesday, February 18, 2015**

@crearc What liquid version is in your bundle?
---
> from: [**crearc**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-75078943) on: **Thursday, February 19, 2015**

liquid (3.0.1) straight from my Gemfile.lock
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/265#issuecomment-75108723) on: **Thursday, February 19, 2015**

@crearc What did you run? &#x60;bundle exec jekyll build&#x60;?
