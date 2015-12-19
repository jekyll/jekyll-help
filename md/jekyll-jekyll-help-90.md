# [Error installing jekyll. Failed to build gem native extension](https://github.com/jekyll/jekyll-help/issues/90)

> state: **closed** opened by: **** on: ****

I&#x27;m trying to install jekyll on OSX 10.7.5 and not sure what&#x27;s wrong. I&#x27;ve updated RubyGems as described on the [troubleshooting page](http://jekyllrb.com/docs/troubleshooting/) but when I try to install jekyll I get errors:

&#x60;&#x60;&#x60;bash
skube:~$ sudo gem update --system
Latest version currently installed. Aborting.
skube:~$ sudo gem install jekyll
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/1.8/usr/bin/ruby extconf.rb
checking for ffi.h... no
checking for ffi.h in /usr/local/include,/usr/include/ffi... no
checking for rb_thread_blocking_region()... no
checking for rb_thread_call_with_gvl()... no
checking for rb_thread_call_without_gvl()... no
checking for ffi_prep_cif_var()... no
creating extconf.h
creating Makefile

make  clean

make
mkdir -p &quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c&quot;/libffi-i386; (if [ ! -f &quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c&quot;/libffi-i386/Makefile ]; then echo &quot;Configuring libffi for i386&quot;; cd &quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c&quot;/libffi-i386 &amp;&amp; env CC=&quot; xcrun cc&quot; CFLAGS=&quot;-arch i386 &quot; LDFLAGS=&quot;-arch i386&quot; &quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c/libffi&quot;/configure --disable-static --with-pic=yes --disable-dependency-tracking --host=i386-apple-darwin &gt; /dev/null; fi); env MACOSX_DEPLOYMENT_TARGET=10.4 make -C &quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c&quot;/libffi-i386
Configuring libffi for i386
configure: WARNING: if you wanted to set the --build type, don&#x27;t use --host.
    If a cross compiler is detected then cross compile mode will be used
configure: error: in &#x60;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c/libffi-i386&#x27;:
configure: error: C compiler cannot create executables
See &#x60;config.log&#x27; for more details
make[1]: *** No targets specified and no makefile found.  Stop.
make: *** [&quot;/Library/Ruby/Gems/1.8/gems/ffi-1.9.3/ext/ffi_c&quot;/libffi-i386/.libs/libffi_convenience.a] Error 2

make failed, exit code 2

Gem files will remain installed in /Library/Ruby/Gems/1.8/gems/ffi-1.9.3 for inspection.
Results logged to /Library/Ruby/Gems/1.8/extensions/universal-darwin-11/1.8/ffi-1.9.3/gem_make.out
&#x60;&#x60;&#x60;

### Comments

---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48097525) on: **Saturday, July 05, 2014**

The system version of ruby on Mac OS X 10.7.5 is too old. Your best bet is to upgrade to to Mac OS X Mavericks (10.9) which is free or install rvm on 10.7.5. 
---
> from: [**skube**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48109621) on: **Sunday, July 06, 2014**

Now that would be a more understandable error message. :confused:  
I appreciate the translation.
---
> from: [**skube**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48110484) on: **Sunday, July 06, 2014**

I tried updating rvm with &#x60;\curl -sSL https://get.rvm.io | bash -s stable&#x60; and it brought me to v1.25.28

Unfortunately I still get the above error when I try to install jekyll. I guess upgrading to Mavericks is my only option.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48111878) on: **Sunday, July 06, 2014**

If you&#x27;re still determined to get it running on 10.7
Try installing RVM with autolibs.
&#x60;\curl -L https://get.rvm.io | bash -s stable --autolibs=enable --ruby&#x60;

---
> from: [**skube**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48113901) on: **Sunday, July 06, 2014**

I tried your suggestion, yet after I run:
&#x60;\curl -L https://get.rvm.io | bash -s stable --autolibs=enable --ruby&#x60; and it completes, running &#x60;sudo gem install jekyll&#x60; still results in the above error.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48116984) on: **Sunday, July 06, 2014**

Are you sure your terminal session is running ruby v2.1.2?

You can verify this by running &#x60;ruby -v&#x60;

Try this:
&#x60;&#x60;&#x60;
source $HOME/.rvm/scripts/rvm
rvm use 2.1.2
gem install jekyll
&#x60;&#x60;&#x60;
---
> from: [**skube**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48164604) on: **Monday, July 07, 2014**

Yes! That was indeed the problem. Thank you :+1: 
---
> from: [**skube**](https://github.com/jekyll/jekyll-help/issues/90#issuecomment-48366758) on: **Tuesday, July 08, 2014**

One note: I had to use the &#x60;--default&#x60; flag to I guess persist it, So:  &#x60;rvm --default use 2.1.2&#x60;
