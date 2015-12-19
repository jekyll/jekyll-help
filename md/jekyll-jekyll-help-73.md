# [Install Jekyll Error](https://github.com/jekyll/jekyll-help/issues/73)

> state: **closed** opened by: **** on: ****

I am installing jekyll using &#x27;sudo gem install jekyll&#x27; on Yosemite 10.10 DP1. I get the following error.

&#x60;&#x60;&#x60;text
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
creating Makefile

make &quot;DESTDIR=&quot; clean

make &quot;DESTDIR=&quot;
make: *** No rule to make target &#x60;/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.9.sdk/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/include/ruby-2.0.0/universal-darwin14/ruby/config.h&#x27;, needed by &#x60;autolink.o&#x27;.  Stop.

make failed, exit code 2

Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/redcarpet-3.1.2 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-14/2.0.0/redcarpet-3.1.2/gem_make.out
&#x60;&#x60;&#x60;

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-46107769) on: **Saturday, June 14, 2014**

You don&#x27;t have the ruby headers:

&#x60;&#x60;&#x60;text
ruby/config.h&#x27;, needed by &#x60;autolink.o&#x27;
&#x60;&#x60;&#x60;

If you have homebrew installed, run &#x60;brew install ruby&#x60; and try again! @mattr- / @robin850 might have more knowledge of this particular area.
---
> from: [**robin850**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-46109311) on: **Sunday, June 15, 2014**

To be honest, I don&#x27;t own a Mac so I don&#x27;t know how you could install Ruby headers but this seems to be [the answer](http://stackoverflow.com/a/19532316/2663959). You may also encounter vmg/redcarpet#357. I hope it helps!
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-46118334) on: **Sunday, June 15, 2014**

I feel compelled to point out and remind that executing &#x60;sudo gem install anything&#x60; on a *nix-based-derived machine is asking for and creating trouble, and is a poor approach to managing gems.

This link is python-based but described the general, consistent scenario, concerns, and pitfalls:

https://workaround.org/easy-install-debian

One would be better off, safer, and the system (especially the tinker-hostile OSX) positively intact by leveraging something such as &#x60;homebrew&#x60;, &#x60;rbenv&#x60;, &#x60;rvm&#x60;, &#x60;chruby&#x60;, etc. and or using &#x60;gem&#x60; with its native switches such as &#x60;--user-install&#x60;, etc.

---
> from: [**clayspencer**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-46130600) on: **Sunday, June 15, 2014**

parkr&#x27;s suggestiion is ultimately what fixed the problem. Fyi jaybe-jekyll I installed several gems without any issues, though in some cases I did seem to have to re-update. As for robin850, that method is deprecated on Xcode 5.1.1 as Command Line Tools are now built in. Thanks all for the recommendations. Actually, I was trying to figure out an issue I am having with Opera malfunctioning in Yosemite 10.10 DP1 and began compile dev.opera not realizing Opera is not open source! dev.opera is their discussion forum for Opera development. That looks like it may be a kernel issue (unrelated to this discussion however). Thanks. Also, sorry for getting back so late, I ended up rashly deleting my ruby.framework and had to reinstall DP1. Yosemite 10.10 has been a challenge for software compatibility so far. 

So, in conclusion in order to fix the error above just brew install ruby. 

-Clay Spencer
---
> from: [**tannerjt**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-60533811) on: **Sunday, October 26, 2014**

Parkr&#x27;s suggestion worked for me as well.  Thanks!
---
> from: [**Hasset**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-68479608) on: **Wednesday, December 31, 2014**

same issue here.
MAC OS X 10.9.5
ruby already installed by brew:
&#x60;&#x60;&#x60;
$ brew install ruby
Warning: ruby-2.2.0 already installed
$ sudo gem install jekyll
Password:
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
creating Makefile

make &quot;DESTDIR=&quot;
make: *** No rule to make target &#x60;/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.10.sdk/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/include/ruby-2.0.0/universal-darwin13/ruby/config.h&#x27;, needed by &#x60;yajl.o&#x27;.  Stop.


Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.1.0 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.1.0/ext/yajl/gem_make.out
&#x60;&#x60;&#x60;

update: problem soved by [adding ruby&#x27;s path to the env](http://hasset.github.io/it/2015/01/01/issue-of-ruby-while-installing-jekyll/).
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/73#issuecomment-68513299) on: **Friday, January 02, 2015**

Yep. It&#x27;s asking for the ruby header files. In this case, &#x60;ruby/config.h&#x60;.
