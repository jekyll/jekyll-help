# [Unable to &#x60;bundle exec jekyll build&#x60; because of missing gems - &#x60;nokogiri&#x60; issue](https://github.com/jekyll/jekyll-help/issues/288)

> state: **closed** opened by: **** on: ****

Hi

I&#x27;m unable to &#x60;bundle exec jekyll build&#x60; my blog, it reports that the &#x60;github-pages&#x60; gem is missing:

    Could not find gem &#x27;github-pages (= 33) ruby&#x27; in the gems available on this machine.

So I tried &#x60;bundle install&#x60; but I run into the issue with &#x60;nokigori&#x60; again:

    Fetching gem metadata from https://rubygems.org/..........
    Resolving dependencies...
    Using RedCloth 4.2.9
    Using i18n 0.7.0
    Using json 1.8.2
    Using minitest 5.5.1
    Using thread_safe 0.3.5
    Using tzinfo 1.2.2
    Using activesupport 4.2.1
    Using blankslate 2.1.2.4
    Using hitimes 1.2.2
    Using timers 4.0.1
    Using celluloid 0.16.0
    Using fast-stemmer 1.0.2
    Using classifier-reborn 2.0.3
    Using coffee-script-source 1.9.1
    Using execjs 2.4.0
    Using coffee-script 2.3.0
    Using colorator 0.1
    Using ffi 1.9.8
    Using gemoji 2.1.0
    Using net-dns 0.8.0
    Using public_suffix 1.5.0
    Using github-pages-health-check 0.2.2
    Using jekyll-coffeescript 1.0.1
    Using jekyll-gist 1.2.1
    Using jekyll-paginate 1.1.0
    Using sass 3.4.13
    Using jekyll-sass-converter 1.2.0
    Using rb-fsevent 0.9.4
    Using rb-inotify 0.9.5
    Using listen 2.10.0
    Using jekyll-watch 1.2.1
    Using kramdown 1.5.0
    Using liquid 2.6.1
    Using mercenary 0.3.5
    Using posix-spawn 0.3.10
    Using yajl-ruby 1.2.1
    Using pygments.rb 0.6.1
    Using redcarpet 3.1.2
    Using safe_yaml 1.0.4
    Using parslet 1.5.0
    Using toml 0.1.2
    Using jekyll 2.4.0
    Using mini_portile 0.6.2

    Gem::Ext::BuildError: ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20150402-12603-e1uja9.rb extconf.rb --use-system-libraries
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

    Gem files will remain installed in /var/folders/md/qy7y7r3d2zn7wfmcpvc2mtk80000gn/T/bundler20150402-12603-19vuinn/nokogiri-1.6.6.2/gems/nokogiri-1.6.6.2 for inspection.
    Results logged to /var/folders/md/qy7y7r3d2zn7wfmcpvc2mtk80000gn/T/bundler20150402-12603-19vuinn/nokogiri-1.6.6.2/extensions/universal-darwin-14/2.0.0/nokogiri-1.6.6.2/gem_make.out
    An error occurred while installing nokogiri (1.6.6.2), and Bundler cannot continue.
    Make sure that &#x60;gem install nokogiri -v &#x27;1.6.6.2&#x27;&#x60; succeeds before bundling.

I tried &#x60;gem install nokogiri -v &#x27;1.6.6.2&#x27;&#x60; but I had the usual error message which is

    Building native extensions.  This could take a while...
    ERROR:  Error installing nokogiri:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20150402-7321-15frer3.rb extconf.rb
    checking if the C compiler accepts ... yes
    checking if the C compiler accepts -Wno-error=unused-command-line-argument-hard-error-in-future... no
    Building nokogiri using packaged libraries.
    checking for gzdopen() in -lz... yes
    checking for iconv... yes
    ************************************************************************
    IMPORTANT NOTICE:

    Building Nokogiri with a packaged version of libxml2-2.9.2
    with the following patches applied:
	- 0001-Revert-Missing-initialization-for-the-catalog-module.patch
	- 0002-Fix-missing-entities-after-CVE-2014-3660-fix.patch

    Team Nokogiri will keep on doing their best to provide security
    updates in a timely manner, but if this is a concern for you and want
    to use the system library instead; abort this installation process and
    reinstall nokogiri as follows:

    gem install nokogiri -- --use-system-libraries
        [--with-xml2-config=/path/to/xml2-config]
        [--with-xslt-config=/path/to/xslt-config]
    
    If you are using Bundler, tell it to use the option:

    bundle config build.nokogiri --use-system-libraries
    bundle install

    Note, however, that nokogiri is not fully compatible with arbitrary
    versions of libxml2 provided by OS/package vendors.
    ************************************************************************
    Extracting libxml2-2.9.2.tar.gz into tmp/x86_64-apple-darwin14/ports/libxml2/2.9.2... OK
    Running patch with /Library/Ruby/Gems/2.0.0/gems/nokogiri-1.6.6.2/ports/patches/libxml2/0001-    Revert-Missing-initialization-for-the-catalog-module.patch...
    Running &#x27;patch&#x27; for libxml2 2.9.2... OK
    Running patch with /Library/Ruby/Gems/2.0.0/gems/nokogiri-1.6.6.2/ports/patches/libxml2/0002-    Fix-missing-entities-after-CVE-2014-3660-fix.patch...
    Running &#x27;patch&#x27; for libxml2 2.9.2... OK
    Running &#x27;configure&#x27; for libxml2 2.9.2... OK
    Running &#x27;compile&#x27; for libxml2 2.9.2... ERROR, review &#x27;/Library/Ruby/Gems/2.0.0/gems/nokogiri-1.6.6.2/ext/nokogiri/tmp/x86_64-apple-darwin14/ports/libxml2/2.9.2/compile.log&#x27; to see what happened.
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
	--enable-static
	--disable-static
	--with-zlib-dir
	--without-zlib-dir
	--with-zlib-include
	--without-zlib-include=${zlib-dir}/include
	--with-zlib-lib
	--without-zlib-lib=${zlib-dir}/lib
	--enable-cross-build
	--disable-cross-build
    /Library/Ruby/Gems/2.0.0/gems/mini_portile-0.6.2/lib/mini_portile.rb:279:in &#x60;block in execute&#x27;: Failed to complete compile task (RuntimeError)
	from /Library/Ruby/Gems/2.0.0/gems/mini_portile-0.6.2/lib/mini_portile.rb:271:in &#x60;chdir&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/mini_portile-0.6.2/lib/mini_portile.rb:271:in &#x60;execute&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/mini_portile-0.6.2/lib/mini_portile.rb:70:in &#x60;compile&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/mini_portile-0.6.2/lib/mini_portile.rb:110:in &#x60;cook&#x27;
	from extconf.rb:278:in &#x60;block in process_recipe&#x27;
	from extconf.rb:177:in &#x60;tap&#x27;
	from extconf.rb:177:in &#x60;process_recipe&#x27;
	from extconf.rb:475:in &#x60;&lt;main&gt;&#x27;

    extconf failed, exit code 1

    Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/nokogiri-1.6.6.2 for inspection.
    Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-14/2.0.0/nokogiri-1.6.6.2    /gem_make.out


### Comments

---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/288#issuecomment-88943286) on: **Thursday, April 02, 2015**

I found a solution to the installation of the latest version of &#x60;nokogiri&#x60;, as reported here

https://github.com/sparklemotion/nokogiri/issues/1166

by @dwt, and &#x60;bunde install&#x60; and &#x60;bundle exec jekyll build&#x60; are working fine now.
