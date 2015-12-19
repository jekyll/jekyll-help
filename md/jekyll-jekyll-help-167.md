# [Problem installing jekyll](https://github.com/jekyll/jekyll-help/issues/167)

> state: **closed** opened by: **** on: ****

On my macbook air OSX, 10.9.4, I ran
sudo gem install jekyll
and got the following:

Successfully installed kramdown-1.4.2
Fetching: mercenary-0.3.4.gem (100%)
Successfully installed mercenary-0.3.4
Fetching: safe_yaml-1.0.4.gem (100%)
Successfully installed safe_yaml-1.0.4
Fetching: colorator-0.1.gem (100%)
Successfully installed colorator-0.1
Fetching: yajl-ruby-1.1.0.gem (100%)
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
mkmf.rb can&#x27;t find header files for ruby at /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/include/ruby.h


Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.1.0 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/gems/yajl-ruby-1.1.0/ext/yajl/gem_make.out

I followed your advice and ran successfully: 
sudo gem update --system
but this did not affect the error. The second time it gave the same error:

Fetching: fast-stemmer-1.0.2.gem (100%)
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby -r ./siteconf20141011-1666-1at401u.rb extconf.rb
mkmf.rb can&#x27;t find header files for ruby at /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/include/ruby.h

extconf failed, exit code 1

Gem files will remain installed in /Library/Ruby/Gems/2.0.0/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /Library/Ruby/Gems/2.0.0/extensions/universal-darwin-13/2.0.0/fast-stemmer-1.0.2/gem_make.out

Can you help??


### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/167#issuecomment-58766685) on: **Saturday, October 11, 2014**

It seems like you have to install the Xcode Command Line Tools first (run &#x60;xcode-select --install&#x60; in your terminal).

Hope that helps!
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/167#issuecomment-58775692) on: **Saturday, October 11, 2014**

Also running &#x60;gem install jekyll&#x60; with sudo should not be necessary and may lead to problems later on.
---
> from: [**dbmumford**](https://github.com/jekyll/jekyll-help/issues/167#issuecomment-58818226) on: **Sunday, October 12, 2014**

On Oct 11, 2014, at 6:22 PM, Alfred Xing &lt;notifications@github.com&gt; wrote:

&gt; It seems like you have to install the Xcode Command Line Tools first (run xcode-select --install in your terminal).
&gt; 
&gt; Hope that helps!
&gt; 

On Oct 12, 2014, at 2:59 AM, Philipp Rudloff &lt;notifications@github.com&gt; wrote:

&gt; Also running gem install jekyll with sudo should not be necessary and may lead to problems later on.
&gt; 
&gt; —
&gt; 
Thanks. I believe it installed fully this time. 

But if I don&#x27;t do it as superuser, I have to change permissions on dozens of folders one at a time (at least, I&#x27;m not sure of the command to do it wholesale on a tree of subfolders so was doing it with command-I, then unlocking each folder). I&#x27;m not sure what problems this may lead to. I&#x27;m working on my own macbook air and am my own admin so I&#x27;m hoping there are no security issues.

David Mumford 
