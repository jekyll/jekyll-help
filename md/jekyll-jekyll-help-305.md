# [Can&#x27;t install Jekyll on Mac OS X 10.9.5](https://github.com/jekyll/jekyll-help/issues/305)

> state: **closed** opened by: **** on: ****

Guidance/suggestions would be much appreciated.  Thanks.
&#x60;&#x60;&#x60;
&gt; sudo gem install jekyll
Password:
Fetching: liquid-2.6.2.gem (100%)
Successfully installed liquid-2.6.2
Fetching: kramdown-1.7.0.gem (100%)
Successfully installed kramdown-1.7.0
Fetching: mercenary-0.3.5.gem (100%)
Successfully installed mercenary-0.3.5
Fetching: safe_yaml-1.0.4.gem (100%)
Successfully installed safe_yaml-1.0.4
Fetching: colorator-0.1.gem (100%)
Successfully installed colorator-0.1
Fetching: yajl-ruby-1.2.1.gem (100%)
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
mkmf.rb can&#x27;t find header files for ruby at /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/include/ruby.h


Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.2.1 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.2.1/ext/yajl/gem_make.out
&#x60;&#x60;&#x60;
Oh well.  Let&#x27;s follow the troubleshooting guide here: http://jekyllrb.com/docs/troubleshooting/#installation-problems

&#x60;&#x60;&#x60;
&gt; sudo gem update --system
Password:
Updating rubygems-update
Fetching: rubygems-update-2.4.6.gem (100%)
Successfully installed rubygems-update-2.4.6
Parsing documentation for rubygems-update-2.4.6
Installing ri documentation for rubygems-update-2.4.6
Installing darkfish documentation for rubygems-update-2.4.6
Installing RubyGems 2.4.6
RubyGems 2.4.6 installed
Parsing documentation for rubygems-2.4.6
Installing ri documentation for rubygems-2.4.6

=== 2.4.6 / 2014-02-05

Bug fixes:
[...]
RubyGems system software updated
&#x60;&#x60;&#x60;
But no ...
&#x60;&#x60;&#x60;
&gt; sudo gem install jekyll
Fetching: fast-stemmer-1.0.2.gem (100%)
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20150510-83523-6gdy6c.rb extconf.rb
mkmf.rb can&#x27;t find header files for ruby at /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/include/ruby.h

extconf failed, exit code 1

Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-13/2.0.0/fast-stemmer-1.0.2/gem_make.out
&#x60;&#x60;&#x60;

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/305#issuecomment-100771247) on: **Sunday, May 10, 2015**

Migrated to Jekyll Talk!
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/305#issuecomment-100771271) on: **Sunday, May 10, 2015**

https://talk.jekyllrb.com/t/cant-install-jekyll-on-mac-os-x-10-9-5/428
