# [Trouble installing Jekyll on OS X 10.9.2](https://github.com/jekyll/jekyll-help/issues/11)

> state: **closed** opened by: **** on: ****

I had a bit of trouble installing Jekyll on OS X 10.9.2 on my Retina MacBook Pro. Here is the error I got when running &#x27;gem install jekyll&#x27;.

&#x60;&#x60;&#x60;bash
sudo gem install jekyll
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /Users/jessica/.rvm/rubies/ruby-1.9.3-p448/bin/ruby extconf.rb
creating Makefile

make  clean

make
compiling porter.c
porter.c:31:44: error: stdlib.h: No such file or directory
porter.c:32:47: error: string.h: No such file or directory
porter.c: In function ‘create_stemmer’:
porter.c:85: warning: implicit declaration of function ‘malloc’
porter.c:85: warning: incompatible implicit declaration of built-in function ‘malloc’
porter.c: In function ‘free_stemmer’:
porter.c:91: warning: implicit declaration of function ‘free’
porter.c: In function ‘ends’:
porter.c:188: warning: implicit declaration of function ‘memcmp’
porter.c: In function ‘setto’:
porter.c:199: warning: implicit declaration of function ‘memmove’
porter.c:199: warning: incompatible implicit declaration of built-in function ‘memmove’
porter.c: In function ‘step1ab’:
porter.c:233: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:234: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:234: warning: passing argument 2 of ‘setto’ discards qualifiers from pointer target type
porter.c:237: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:238: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:238: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:240: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:240: warning: passing argument 2 of ‘setto’ discards qualifiers from pointer target type
porter.c:241: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:241: warning: passing argument 2 of ‘setto’ discards qualifiers from pointer target type
porter.c:242: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:242: warning: passing argument 2 of ‘setto’ discards qualifiers from pointer target type
porter.c:249: warning: passing argument 2 of ‘setto’ discards qualifiers from pointer target type
porter.c: In function ‘step1c’:
porter.c:257: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c: In function ‘step2’:
porter.c:267: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:267: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:268: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:268: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:270: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:270: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:271: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:271: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:273: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:273: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:275: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:275: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:280: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:280: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:281: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:281: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:282: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:282: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:283: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:283: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:285: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:285: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:286: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:286: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:287: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:287: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:289: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:289: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:290: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:290: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:291: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:291: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:292: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:292: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:294: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:294: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:295: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:295: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:296: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:296: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:298: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:298: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c: In function ‘step3’:
porter.c:308: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:308: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:309: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:309: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:310: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:310: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:312: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:312: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:314: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:314: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:315: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:315: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c:317: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:317: warning: passing argument 2 of ‘r’ discards qualifiers from pointer target type
porter.c: In function ‘step4’:
porter.c:325: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:326: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:327: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:328: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:329: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:330: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:331: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:332: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:333: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:334: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:335: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:336: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:337: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:339: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:340: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:341: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:342: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:343: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
porter.c:344: warning: passing argument 2 of ‘ends’ discards qualifiers from pointer target type
make: *** [porter.o] Error 1

make failed, exit code 2

Gem files will remain installed in /Users/jessica/.rvm/rubies/ruby-1.9.3-p448/lib/ruby/gems/1.9.1/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /Users/jessica/.rvm/rubies/ruby-1.9.3-p448/lib/ruby/gems/1.9.1/extensions/x86_64-darwin-12/1.9.1/fast-stemmer-1.0.2/gem_make.out
&#x60;&#x60;&#x60;

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39695137) on: **Sunday, April 06, 2014**

Hello, this issue seems to be similar or the same as issue #9. I notice that you are running rvm and have Ruby version 1.9.3 installed. I will throw a blind guess and say that rvm is not working as it should or something went wrong with the installation. Are you using rvm and that ruby version for other projects? 
---
> from: [**jessicasideways**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39696787) on: **Sunday, April 06, 2014**

Yes, I am. The other gems I’ve installed seem to be working fine.  

--  
Jessica Sideways
« Je me souviendrai, même si je suis née sous le pavot de Californie, élevé sous la rose jaune du Texas et de renaître sous l&#x27;arbre de pluie d&#x27;or de la Thaïlande, je suis maintenant et j&#x27;ai toujours été un enfant de la fleur de lys, une fille de Québec. » ~ Mlle. Jessica Sideways.


Le dimanche 2014 avril 6 à 22:21, Alberto Grespan a écrit :

&gt; Hello, this issue seems to be similar or the same as issue #9 (https://github.com/jekyll/help/issues/9). I notice that you are running rvm and have Ruby version 1.9.3 installed. I will throw a blind guess and say that rvm is not working as it should or something went wrong with the installation. Are you using rvm and that ruby version for other projects?  
&gt;  
&gt; —
&gt; Reply to this email directly or view it on GitHub (https://github.com/jekyll/help/issues/11#issuecomment-39695137).
&gt;  
&gt;  
&gt;  
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39762731) on: **Monday, April 07, 2014**

Can you provide me the output of &#x60;gcc -v&#x60;?
---
> from: [**jessicasideways**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39810356) on: **Monday, April 07, 2014**

Here it is:

Configured with: --prefix=/Applications/Xcode.app/Contents/Developer/usr --with-gxx-include-dir=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.9.sdk/usr/include/c++/4.2.1  
Apple LLVM version 5.1 (clang-503.0.38) (based on LLVM 3.4svn)
Target: x86_64-apple-darwin13.1.0
Thread model: posix



--  
Jessica Sideways
« Je me souviendrai, même si je suis née sous le pavot de Californie, élevé sous la rose jaune du Texas et de renaître sous l&#x27;arbre de pluie d&#x27;or de la Thaïlande, je suis maintenant et j&#x27;ai toujours été un enfant de la fleur de lys, une fille de Québec. » ~ Mlle. Jessica Sideways.


Le lundi 2014 avril 7 à 12:03, Alberto Grespan a écrit :

&gt; Can you provide me the output of gcc -v?
&gt;  
&gt; —
&gt; Reply to this email directly or view it on GitHub (https://github.com/jekyll/help/issues/11#issuecomment-39762731).
&gt;  
&gt;  
&gt;  
&gt;  
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39907231) on: **Tuesday, April 08, 2014**

Well, the gcc version seem to be just fine... On the other hand I will suggest reinstalling your Ruby if you don&#x27;t mind. I haven&#x27;t used rvm for years now so I don&#x27;t really know it any more, but I think this command will work &#x60;rvm reinstall 1.9.3-p448 &amp;&amp; rvm docs generate all&#x60;, also avoid using &#x60;sudo&#x60; if you can! Last but not least if none of the above works you can try Installing a newer version of Ruby, maybe 2.0.0 &#x60;rvm  install ruby-2.0.0-p451&#x60;?
---
> from: [**jessicasideways**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39920545) on: **Tuesday, April 08, 2014**

It looks like 2.0.0 worked out well and 1.9.3 didn’t work out.

--  
Jessica Sideways
« Je me souviendrai, même si je suis née sous le pavot de Californie, élevé sous la rose jaune du Texas et de renaître sous l&#x27;arbre de pluie d&#x27;or de la Thaïlande, je suis maintenant et j&#x27;ai toujours été un enfant de la fleur de lys, une fille de Québec. » ~ Mlle. Jessica Sideways.


Le mardi 2014 avril 8 à 16:00, Alberto Grespan a écrit :

&gt; Well, the gcc version seem to be just fine... On the other hand I will suggest reinstalling your Ruby if you don&#x27;t mind. I haven&#x27;t used rvm for years now so I don&#x27;t really know it any more, but I think this command will work rvm reinstall 1.9.3-p448 &amp;&amp; rvm docs generate all, also avoid using sudo if you can! Last but not least if none of the above works you can try Installing a newer version of Ruby, maybe 2.0.0 rvm install ruby-2.0.0-p451?
&gt;  
&gt; —
&gt; Reply to this email directly or view it on GitHub (https://github.com/jekyll/help/issues/11#issuecomment-39907231).
&gt;  
&gt;  
&gt;  
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/11#issuecomment-39920589) on: **Tuesday, April 08, 2014**

Did you reinstalled 1.9.3? —
Alberto Grespan

On Tue, Apr 8, 2014 at 9:02 PM, Jessica Sideways &lt;notifications@github.com&gt;
wrote:

&gt; It looks like 2.0.0 worked out well and 1.9.3 didn’t work out.
&gt; --  
&gt; Jessica Sideways
&gt; « Je me souviendrai, même si je suis née sous le pavot de Californie, élevé sous la rose jaune du Texas et de renaître sous l&#x27;arbre de pluie d&#x27;or de la Thaïlande, je suis maintenant et j&#x27;ai toujours été un enfant de la fleur de lys, une fille de Québec. » ~ Mlle. Jessica Sideways.
&gt; Le mardi 2014 avril 8 à 16:00, Alberto Grespan a écrit :
&gt;&gt; Well, the gcc version seem to be just fine... On the other hand I will suggest reinstalling your Ruby if you don&#x27;t mind. I haven&#x27;t used rvm for years now so I don&#x27;t really know it any more, but I think this command will work rvm reinstall 1.9.3-p448 &amp;&amp; rvm docs generate all, also avoid using sudo if you can! Last but not least if none of the above works you can try Installing a newer version of Ruby, maybe 2.0.0 rvm install ruby-2.0.0-p451?
&gt;&gt;  
&gt;&gt; —
&gt;&gt; Reply to this email directly or view it on GitHub (https://github.com/jekyll/help/issues/11#issuecomment-39907231).
&gt;&gt;  
&gt;&gt;  
&gt;&gt;  
&gt; ---
&gt; Reply to this email directly or view it on GitHub:
&gt; https://github.com/jekyll/help/issues/11#issuecomment-39920545
