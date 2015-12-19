# [Jekyll refuses to install (copied from main repo)](https://github.com/jekyll/jekyll-help/issues/190)

> state: **closed** opened by: **** on: ****

I believe I have all the dependencies installed:
* Bundler
* Ruby
* RubyGems
* Linux system
* NodeJS

But when I run the command, &#x60;sudo gem install jekyll&#x60;, it fails and gives me:
&#x60;&#x60;&#x60;
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

        /usr/bin/ruby1.9.1 extconf.rb
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;: cannot load such file -- mkmf (LoadError)
	from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from extconf.rb:1:in &#x60;&lt;main&gt;&#x27;


Gem files will remain installed in /var/lib/gems/1.9.1/gems/yajl-ruby-1.1.0 for inspection.
Results logged to /var/lib/gems/1.9.1/gems/yajl-ruby-1.1.0/ext/yajl/gem_make.out
&#x60;&#x60;&#x60;
And the file at &#x60;/var/lib/gems/1.9.1/gems/yajl-ruby-1.1.0/ext/yajl/gem_make.out&#x60; contains:
&#x60;&#x60;&#x60;
/usr/bin/ruby1.9.1 extconf.rb
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;: cannot load such file -- mkmf (LoadError)
        from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
        from extconf.rb:1:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
What do I need to do to fix this? Did I miss a dependency? I am not good with Ruby or RubyGems.

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/190#issuecomment-63156139) on: **Friday, November 14, 2014**

If I recall correctly jekyll won&#x27;t run in that version of ruby, you&#x27;ll need ruby 1.9.3+. Apart from that you may be also missing ruby headers for compiling extensions... the package name is &#x60;ruby-dev&#x60; in debian/ubuntu.
---
> from: [**Plenglin**](https://github.com/jekyll/jekyll-help/issues/190#issuecomment-63410877) on: **Monday, November 17, 2014**

Updating to ruby-dev worked, thanks!
