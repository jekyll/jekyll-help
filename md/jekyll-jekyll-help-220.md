# [Installation problem](https://github.com/jekyll/jekyll-help/issues/220)

> state: **closed** opened by: **** on: ****

ERROR:  Error installing jekyll:
	classifier-reborn requires Ruby version &gt;= 1.9.3.

ruby -v returned this:
ruby 1.8.7 (2012-02-08 patchlevel 358) [universal-darwin10.0]



### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/220#issuecomment-68216671) on: **Sunday, December 28, 2014**

You&#x27;ll need to upgrade your Ruby to version 1.9.3 or above, though I&#x27;d recommend the latest stable, 2.2.

Since you&#x27;re on OS X, I would recommend doing this by using a Ruby manager. Some good choices are:
 - [RVM](http://rvm.io/)
 - [rbenv](https://github.com/sstephenson/rbenv) with the [ruby-build plugin](https://github.com/sstephenson/ruby-build)
 - [chruby](https://github.com/postmodern/chruby) with the [ruby-install tool](https://github.com/postmodern/ruby-install)

If you&#x27;re only going to use a single Ruby, [Homebrew](brew.sh) is also a good option.

You can also upgrade your OS to Mavericks (10.9) or Yosemite (10.10); these should include a more up-to-date version of Ruby.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/220#issuecomment-68660192) on: **Sunday, January 04, 2015**

Thanks for the thorough explanation, @alfredxing!
