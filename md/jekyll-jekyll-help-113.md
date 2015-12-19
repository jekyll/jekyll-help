# [Development tools cant be found &amp; Configuration Issues](https://github.com/jekyll/jekyll-help/issues/113)

> state: **closed** opened by: **** on: ****

Hello I am trying to install Jekyll and I have scoured the internet for solutions but not made much of a dent.  I have the latest version of ruby installed along with homebrew and the developer tools.  When in try to install the xcode via &quot;xcode-select --install&quot; I receive the following even though my path is to &quot;/Applications/Xcode.app/Contents/Developer&quot;  
&#x60;&#x60;&#x60;
xcode-select: Error: unknown command option &#x27;--install&#x27;.

xcode-select: Report or change the path to the active
              Xcode installation for this machine.

Usage: xcode-select --print-path
           Prints the path of the active Xcode folder
   or: xcode-select --switch &lt;xcode_path&gt;
           Sets the path for the active Xcode folder
   or: xcode-select --version
           Prints the version of xcode-select
&#x60;&#x60;&#x60;


When I try to install Jekyll or anything similar I receive:
&#x60;&#x60;&#x60;
new-host:~ jameswaltz$ gem install jekyll
Building native extensions.  This could take a while...
/Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/site_ruby/2.1.0/rubygems/ext/builder.rb:73: warning: Insecure world writable dir /usr/local in PATH, mode 040777
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/bin/ruby -r ./siteconf20140807-99237-l32ov4.rb extconf.rb
/Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:1523: warning: Insecure world writable dir /usr/local in PATH, mode 040777
checking for ffi.h... *** extconf.rb failed ***
Could not create Makefile due to some reason, probably lack of necessary
libraries and/or headers.  Check the mkmf.log file for more details.  You may
need configuration options.

Provided configuration options:
	--with-opt-dir
	--without-opt-dir
	--with-opt-include
	--without-opt-include=${opt-dir}/include
	--with-opt-lib
	--without-opt-lib=${opt-dir}/lib
	--with-make-prog
	--without-make-prog
	--srcdir=.
	--curdir
	--ruby=/Users/jameswaltz/.rvm/rubies/ruby-2.1.2/bin/ruby
	--with-ffi_c-dir
	--without-ffi_c-dir
	--with-ffi_c-include
	--without-ffi_c-include=${ffi_c-dir}/include
	--with-ffi_c-lib
	--without-ffi_c-lib=${ffi_c-dir}/lib
	--with-libffi-config
	--without-libffi-config
	--with-pkg-config
	--without-pkg-config
/Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:456:in &#x60;try_do&#x27;: The compiler failed to generate an executable file. (RuntimeError)
You have to install development tools first.
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:587:in &#x60;try_cpp&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:1067:in &#x60;block in have_header&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:918:in &#x60;block in checking_for&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:351:in &#x60;block (2 levels) in postpone&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:321:in &#x60;open&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:351:in &#x60;block in postpone&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:321:in &#x60;open&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:347:in &#x60;postpone&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:917:in &#x60;checking_for&#x27;
	from /Users/jameswaltz/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/mkmf.rb:1066:in &#x60;have_header&#x27;
	from extconf.rb:16:in &#x60;&lt;main&gt;&#x27;

extconf failed, exit code 1

Gem files will remain installed in /Users/jameswaltz/.rvm/gems/ruby-2.1.2/gems/ffi-1.9.3 for inspection.
Results logged to /Users/jameswaltz/.rvm/gems/ruby-2.1.2/extensions/x86_64-darwin-13/2.1.0-static/ffi-1.9.3/gem_make.out
&#x60;&#x60;&#x60;

Thanks for any help!

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51561315) on: **Thursday, August 07, 2014**

It seems that you need to fix some permissions first. Go to Disk Utility, select your Macintosh HD and Press Repair Permissions. Lets see if we can get rid of &#x60;warning: Insecure world writable dir /usr/local in PATH, mode 040777&#x60;. After that, try to run &#x60;xcode-select --install&#x60;. This command should pop up something like a banner that you have to press install to get the command line tools. This will not install xcode though. Remember that you can also download the latest Command Line Tools from the 
Developer Apple site. BTW, if you type &#x60;brew doctor&#x60; what&#x27;s the output?
---
> from: [**j-waltz**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51562629) on: **Thursday, August 07, 2014**

Thanks! I will try to repair permissions when I am back at my computer tomorrow.  I originally installed the developer tools from the apple site and they are installed on my computer but they do not seem to get recognized.  My brew doctor is up to date and the system is ready to brew.
---
> from: [**j-waltz**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51812232) on: **Monday, August 11, 2014**

I have repaired my permissions but I am still receiving the same message as originally posted and when I install xcode i receive the following and my path is set to &#x27;/Applications/Xcode.app/Contents/Developer&#x27;:
&#x60;&#x60;&#x60;
xcode-select: Report or change the path to the active
              Xcode installation for this machine.

Usage: xcode-select --print-path
           Prints the path of the active Xcode folder
   or: xcode-select --switch &lt;xcode_path&gt;
           Sets the path for the active Xcode folder
   or: xcode-select --version
           Prints the version of xcode-select
&#x60;&#x60;&#x60;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51857821) on: **Monday, August 11, 2014**

Hi @j-waltz, My have my path is set to:

&#x60;&#x60;&#x60;
$ xcode-select --print-path
/Library/Developer/CommandLineTools
&#x60;&#x60;&#x60;

I don&#x27;t have xcode installed though, just the Command Line Tools. I don&#x27;t really know whats happening here with xcode or of your path is ok. I can recommend you try to reinstall your ruby. &#x60;rvm reinstall &lt;your-ruby-version&gt;&#x60; to see if you can solve the ruby related problem.
---
> from: [**j-waltz**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51859954) on: **Monday, August 11, 2014**

Success! I had tried reinstalling ruby before but I think repairing my permissions and then reinstalling seemed to do the trick. I do still receive the warning below, but since it is a personal computer it shouldn&#x27;t cause any issues, correct?

&#x60;&#x60;&#x60;
warning: Insecure world writable dir /usr/local in PATH, mode 040777

&#x60;&#x60;&#x60;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51863462) on: **Monday, August 11, 2014**

I&#x27;m glad it worked, It&#x27;s weird that you continue having that warning, you should try to Repair Permissions again and check if the warning goes away. If that doesn&#x27;t work try &#x60;chmod go-w /usr/local&#x60; good luck!
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/113#issuecomment-51864635) on: **Monday, August 11, 2014**

I&#x27;m going to close the issue, feel free to reopen it again or create a new one if you have any other problem!
