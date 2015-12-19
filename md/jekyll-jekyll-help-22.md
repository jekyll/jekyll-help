# [Installation fails](https://github.com/jekyll/jekyll-help/issues/22)

> state: **closed** opened by: **** on: ****

Hi,
just wante to give jekyll a try and can&#x27;t pass the installation on MacOSX.
I did a &quot;sudo gem update --system&quot; and then tried to install jekyll with &quot;sudo gem install jekyll&quot;, which results in:
&#x60;&#x60;&#x60;
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
What am I doing wrong?

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/22#issuecomment-41373572) on: **Friday, April 25, 2014**

Have a look through jekyll/jekyll#2125
---
> from: [**kevinpapst**](https://github.com/jekyll/jekyll-help/issues/22#issuecomment-41376150) on: **Friday, April 25, 2014**

Thanks for the link :) could install it now, closing this issue.
