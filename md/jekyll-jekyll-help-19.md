# [Error Installing Jekyll with Native Mavericks Ruby](https://github.com/jekyll/jekyll-help/issues/19)

> state: **closed** opened by: **** on: ****

Had native Mavericks ruby, and updated gems, and XCode Command LIne tools. Ran into this error:
&#x60;&#x60;&#x60;
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
&#x60;&#x60;&#x60;

Was able to solve the problem with the guidance of http://stackoverflow.com/questions/22479246/how-to-install-jekyll-on-mac-osx-10-9-with-ruby-2-0-0, which suggested installing &#x60;rvm&#x60; and then doing &#x60;rvm use ruby-2.1.1&#x60; and then &#x60;sudo gem install jekyll&#x60;

Figured other people could benefit from the rvm/version suggestion. Cheers!


### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/19#issuecomment-40559781) on: **Tuesday, April 15, 2014**

Thank you @jennazee, It will be very useful!  
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/19#issuecomment-40584456) on: **Wednesday, April 16, 2014**

:+1: 
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/19#issuecomment-40598918) on: **Wednesday, April 16, 2014**

Note, if you&#x27;re using rvm there&#x27;s never a reason to install gems using
&#x60;sudo&#x60;. Please don&#x27;t do that, it&#x27;ll only cause problems.
---
> from: [**jennazee**](https://github.com/jekyll/jekyll-help/issues/19#issuecomment-40599872) on: **Wednesday, April 16, 2014**

Oh okay, noted. Without rvm it seemed I didn&#x27;t have permissions so I
figured I needed to use it. Any problems I should look out for for this
case?

On Wednesday, April 16, 2014, Matthew Scharley &lt;notifications@github.com&gt;
wrote:

&gt; Note, if you&#x27;re using rvm there&#x27;s never a reason to install gems using
&gt; &#x60;sudo&#x60;. Please don&#x27;t do that, it&#x27;ll only cause problems.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/19#issuecomment-40598918&gt;
&gt; .
&gt;


-- 
Jenna Zeigen
jennazeigen.com
@zeigenvector
215.833.5104
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/19#issuecomment-40600162) on: **Wednesday, April 16, 2014**

Without RVM, ruby is installed into the system locations and so it needs
sudo access. RVM installs ruby into your home folder, &#x60;~/.rvm&#x60;, so you
don&#x27;t need any special sort of access ever. One of the many benefits of RVM.
​
