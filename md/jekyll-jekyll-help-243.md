# [&#x60;bundle exec jekyll&#x60; build fails after reboot](https://github.com/jekyll/jekyll-help/issues/243)

> state: **open** opened by: **** on: ****

Hi

I have a problem when I try to build my site locally using &#x60;bundle exec jekyll build&#x60;, I receive the following message

    Bundler could not find compatible versions for gem &quot;kramdown&quot;:
      In snapshot (Gemfile.lock):
        kramdown (1.3.1)

      In Gemfile:
        github-pages (= 32) ruby depends on
          kramdown (= 1.5.0) ruby

    Running &#x60;bundle update&#x60; will rebuild your snapshot from scratch, using only
    the gems in your Gemfile, which may resolve the conflict.

There is a conflict between the kramdown versions because github-pages requires a later version?  I have tried changing the Gemfile entry for github-pages to require kramdown version &#x27;~&gt; 1.3.1&#x27; but this does not work.

Is there an easy fix for this?


### Comments

---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-70947441) on: **Wednesday, January 21, 2015**

Could someone please help?
---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-71005774) on: **Thursday, January 22, 2015**

Sure,

1. What&#x27;s the output of &#x60;bundle update&#x60;?
2. Would you be up for posting your &#x60;Gemfile&#x60; and &#x60;Gemfile.lock&#x60;? (here, or link to a gist)

---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-71008459) on: **Thursday, January 22, 2015**

I deleted the Gemfile.lock and ran &#x60;bundle exec jekyll build&#x60; again, which created the right Gemfile again and it worked.

Separately &#x60;bundle update&#x60; still shows problems.  Here is the output.

    $ bundle update
    Fetching gem metadata from https://rubygems.org/..........
    Resolving dependencies...
    Using RedCloth 4.2.9
    Using i18n 0.7.0
    Using json 1.8.2
    Using minitest 5.5.1
    Using thread_safe 0.3.4
    Using tzinfo 1.2.2
    Using activesupport 4.2.0
    Using blankslate 2.1.2.4
    Using hitimes 1.2.2
    Using timers 4.0.1
    Using celluloid 0.16.0
    Using fast-stemmer 1.0.2
    Using classifier-reborn 2.0.3
    Using coffee-script-source 1.8.0
    Using execjs 2.2.2
    Using coffee-script 2.3.0
    Using colorator 0.1
    Using ffi 1.9.6
    Using gemoji 2.1.0
    Using net-dns 0.8.0
    Using public_suffix 1.4.6
    Using github-pages-health-check 0.2.1
    Using jekyll-coffeescript 1.0.1
    Using jekyll-gist 1.1.0
    Using jekyll-paginate 1.1.0
    Using sass 3.4.10
    Using jekyll-sass-converter 1.2.0
    Using rb-fsevent 0.9.4
    Using rb-inotify 0.9.5
    Using listen 2.8.5
    Using jekyll-watch 1.2.0
    Using kramdown 1.5.0
    Using liquid 2.6.1
    Using mercenary 0.3.5
    Using posix-spawn 0.3.9
    Using yajl-ruby 1.1.0
    Using pygments.rb 0.6.0
    Using redcarpet 3.1.2
    Using safe_yaml 1.0.4
    Using parslet 1.5.0
    Using toml 0.1.2
    Using jekyll 2.4.0
    Using mini_portile 0.6.2

    Gem::Ext::BuildError: ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20150122-1743-    d1uujk.rb extconf.rb --use-system-libraries
    checking if the C compiler accepts ... yes
    checking if the C compiler accepts -Wno-error=unused-command-line-argument-hard-error-in-future... no
    Building nokogiri using system libraries.
    checking for xmlParseDoc() in libxml/parser.h... no
    checking for xmlParseDoc() in -lxml2... no
    checking for xmlParseDoc() in -llibxml2... no
    -----
    libxml2 is missing.  Please locate mkmf.log to investigate how it is failing.
    -----
    *** extconf.rb failed ***
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
	--ruby=/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby
	--help
	--clean
	--use-system-libraries
	--with-zlib-dir
	--without-zlib-dir
	--with-zlib-include
	--without-zlib-include=${zlib-dir}/include
	--with-zlib-lib
	--without-zlib-lib=${zlib-dir}/lib
	--with-xml2-dir
	--without-xml2-dir
	--with-xml2-include
	--without-xml2-include=${xml2-dir}/include
	--with-xml2-lib
	--without-xml2-lib=${xml2-dir}/lib
	--with-libxml-2.0-config
	--without-libxml-2.0-config
	--with-pkg-config
	--without-pkg-config
	--with-xslt-dir
	--without-xslt-dir
	--with-xslt-include
	--without-xslt-include=${xslt-dir}/include
	--with-xslt-lib
	--without-xslt-lib=${xslt-dir}/lib
	--with-libxslt-config
	--without-libxslt-config
	--with-exslt-dir
	--without-exslt-dir
	--with-exslt-include
	--without-exslt-include=${exslt-dir}/include
	--with-exslt-lib
	--without-exslt-lib=${exslt-dir}/lib
	--with-libexslt-config
	--without-libexslt-config
	--with-xml2lib
	--without-xml2lib
	--with-libxml2lib
	--without-libxml2lib

    extconf failed, exit code 1

    Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/nokogiri-1.6.5 for inspection.
    Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-14/2.0.0/nokogiri-1.6.5/gem_make.out
    An error occurred while installing nokogiri (1.6.5), and Bundler cannot continue.
    Make sure that &#x60;gem install nokogiri -v &#x27;1.6.5&#x27;&#x60; succeeds before bundling.
---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-71017806) on: **Thursday, January 22, 2015**

Great! If you&#x27;re happy with everything working, fine, if not, read on:

Thanks for pasting the output, that&#x27;s useful. The trouble installing &#x60;nokogiri&#x60; could be for one of two reasons:

1. You don’t have the build toolchain installed
2. You don’t have the package &#x60;libxml2&#x60; installed (or you need to point to it correctly)

I would try installing the build tools via &#x60;xcode-select --install&#x60; on the command line, then re-running. If they are already installed, then there are a bunch of StackOverflow Questions about &#x60;libxml2&#x60; and &#x60;nokogiri&#x60; but be careful as instructions vary depending on the environment (e.g. Mavericks vs Yosemite, Homebrew vs preinstalled).
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-71025072) on: **Thursday, January 22, 2015**

Hi

I don&#x27;t want anything to affect my Jekyll setup, and I don&#x27;t use any XML parsing so I&#x27;m OK with nokogiri not installing but ...actually &#x60;libxml2&#x60; seems to be in several places:

    /opt/local/var/macports/software/libxml2
    /opt/local/share/gtk-doc/html/libxml2
    /opt/local/lib/cmake/libxml2
    /opt/local/include/libxml2
    /usr/include/libxml2

If I did need to use it for something else, what is the right location and how would I point the nokogiri installation to the right location?

---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/243#issuecomment-71035966) on: **Thursday, January 22, 2015**

To specify the path to &#x60;libxml2&#x60; explicitly you can follow the instructions on the [Nokogiri Install Page](http://www.nokogiri.org/tutorials/installing_nokogiri.html#mac_os_x), especially the **Using Nonstandard libxml2 / libxslt installations** section. It involves installing &#x60;nokogiri&#x60; via &#x60;gem install&#x60; and passing in some configuration options. In terms of what the correct location is, I&#x27;m not sure, that&#x27;s where my knowledge begins to end.

Of course it may be the build tools, that would be my first port of call; what&#x27;s the result of running &#x60;make -v&#x60; on the command line? If it&#x27;s similar to&#x60;GNU Make 3.81&#x60; you have them. If not, run &#x60;xcode-select --install&#x60;.

