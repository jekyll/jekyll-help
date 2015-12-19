# [Installation error](https://github.com/jekyll/jekyll-help/issues/30)

> state: **closed** opened by: **** on: ****

When I try to install jekyll on my Mac (OSX 10.9.2), I get the following errors:

&#x60;&#x60;&#x60;
$ sudo gem install jekyll
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
creating Makefile

make &quot;DESTDIR=&quot; clean

make &quot;DESTDIR=&quot;
compiling porter.c
porter.c:359:27: warning: &#x27;&amp;&amp;&#x27; within &#x27;||&#x27; [-Wlogical-op-parentheses]
      if (a &gt; 1 || a == 1 &amp;&amp; !cvc(z, z-&gt;k - 1)) z-&gt;k--;
                ~~ ~~~~~~~^~~~~~~~~~~~~~~~~~~~
porter.c:359:27: note: place parentheses around the &#x27;&amp;&amp;&#x27; expression to silence this warning
      if (a &gt; 1 || a == 1 &amp;&amp; !cvc(z, z-&gt;k - 1)) z-&gt;k--;
                          ^
                   (                          )
1 warning generated.
compiling porter_wrap.c
linking shared-object stemmer.bundle
clang: error: unknown argument: &#x27;-multiply_definedsuppress&#x27; [-Wunused-command-line-argument-hard-error-in-future]
clang: note: this will be a hard error (cannot be downgraded to a warning) in the future
make: *** [stemmer.bundle] Error 1

make failed, exit code 2

Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-13/2.0.0/fast-stemmer-1.0.2/gem_make.out
&#x60;&#x60;&#x60;

I have command line developer tools 5.1.0.0 and XCode 5.1.1. I can bypass the error by typing &#x60;ARCHFLAGS=-Wno-error=unused-command-line-argument-hard-error-in-future&#x60; before the &#x60;gem install&#x60;, but not sure if that was a good idea!

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/30#issuecomment-42261002) on: **Monday, May 05, 2014**

Hi, I&#x27;m sorry about this problem, but it seems to be related with Mac OS X System Ruby. To solve this problem install Ruby using a version manager like rbenv, chruby or rvm. You can also take a look at issue #11 and #19. Remember that when using a Ruby version manager you don&#x27;t need to use &#x60;sudo&#x60; to install anything.

Hope this helps! 
---
> from: [**bkfunk**](https://github.com/jekyll/jekyll-help/issues/30#issuecomment-42349249) on: **Tuesday, May 06, 2014**

Yeah, I found an SO post (I think...somewhere, anyway) that said stemmer was the problem, and suggested the workaround. Anyway, seems to be working now, and hopefully I didn&#x27;t do too much damage.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/30#issuecomment-42356253) on: **Tuesday, May 06, 2014**

Thanks! it&#x27;s nice to know that work around like this exists. I have read a bit about that problem and the flag you used, it seems to be something related with the new version of clang. It also seems that it doesn&#x27;t do any damage, but I still suggest installing a fresh copy of ruby using a version manager.
