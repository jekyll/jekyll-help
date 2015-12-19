# [problem installing jekyll](https://github.com/jekyll/jekyll-help/issues/194)

> state: **closed** opened by: **** on: ****

I&#x27;m on a MAC OSX 10.9.5
I successfully executed sudo gem update --system
but then on executing
sudo gem install jekyll

I got the following error. Any advice appreciated.
(It seems from the error that it might have something to do with my system being 10.9.5 
and not 10.10, but I am reluctant to update to 10.10 yet)

$sudo gem install jekyll
Fetching: fast-stemmer-1.0.2.gem (100%)
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20141116-21190-1elw968.rb extconf.rb
creating Makefile

make &quot;DESTDIR=&quot; clean

make &quot;DESTDIR=&quot;
make: *** No rule to make target &#x60;/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.10.sdk/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/include/ruby-2.0.0/universal-darwin13/ruby/config.h&#x27;, needed by &#x60;porter.o&#x27;.  Stop.

make failed, exit code 2

Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-13/2.0.0/fast-stemmer-1.0.2/gem_make.out

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/194#issuecomment-63232363) on: **Sunday, November 16, 2014**

Hi, @stephens999 this is a known problem with the OSX ruby. Take a look at issue #167 and #19 and see if any of those work. I recommend that you install Ruby through a version manager like rbenv, chruby or rvm and avoid using the system Ruby. If you plan using rbenv [I wrote a blog post](http://albertogrespan.com/blog/installing-ruby-the-right-way-on-os-x-using-rbenv/). Hope it helps.
---
> from: [**stephens999**](https://github.com/jekyll/jekyll-help/issues/194#issuecomment-63249467) on: **Sunday, November 16, 2014**

thanks for the pointers - I got this working now by following instructions in issue #19  
