# [jekyll serve](https://github.com/jekyll/jekyll-help/issues/200)

> state: **closed** opened by: **** on: ****

I am trying to load jekyll server to view using http://localhost:4000
But keep getting this.

&#x60;&#x60;&#x60;
jekyll serve
/var/lib/gems/2.1.0/gems/jekyll-2.5.1/lib/jekyll.rb:166:in &#x60;&lt;top (required)&gt;&#x27;: undefined method &#x60;gracefully_require&#x27; for Jekyll::Deprecator:Class (NoMethodError)
	from /usr/lib/ruby/2.1.0/rubygems/core_ext/kernel_require.rb:73:in &#x60;require&#x27;
	from /usr/lib/ruby/2.1.0/rubygems/core_ext/kernel_require.rb:73:in &#x60;require&#x27;
	from /var/lib/gems/2.1.0/gems/jekyll-2.5.1/bin/jekyll:6:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/local/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

I have checked all dependencies installed, running Ubuntu 14.10
Something I must have done incorrectly but not sure what.
Thanks

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63673901) on: **Wednesday, November 19, 2014**

Which version of [Node.js](http://nodejs.org/) do you have? (Use &#x60;node -v&#x60; to find out.)
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63677996) on: **Wednesday, November 19, 2014**

v0.10.33
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63720610) on: **Wednesday, November 19, 2014**

I just tried to get Jekyll going on a fresh VirtualBox machine on my Macbook Air using Vagrant and the [chef/ubuntu-14.10](https://vagrantcloud.com/chef/boxes/ubuntu-14.10) box.

After a quick &#x60;apt-get update&#x60; and &#x60;apt-get install -y build-essential python ruby ruby-dev nodejs&#x60;, I ran &#x60;gem install jekyll&#x60;.

&#x60;&#x60;&#x60;
vagrant@vagrant:~$ python --version &amp;&amp; ruby --version &amp;&amp; nodejs --version &amp;&amp; jekyll --version
Python 2.7.8
ruby 2.1.2p95 (2014-05-08) [x86_64-linux-gnu]
v0.10.25
jekyll 2.5.1
vagrant@vagrant:~$ jekyll new mysite
New jekyll site installed in /home/vagrant/mysite. 
vagrant@vagrant:~$ cd mysite/
vagrant@vagrant:~/mysite$ jekyll serve
Configuration file: /home/vagrant/mysite/_config.yml
            Source: /home/vagrant/mysite
       Destination: /home/vagrant/mysite/_site
      Generating... 
                    done.
 Auto-regeneration: enabled for &#x27;/home/vagrant/mysite&#x27;
Configuration file: /home/vagrant/mysite/_config.yml
    Server address: http://127.0.0.1:4000/
  Server running... press ctrl-c to stop.
&#x60;&#x60;&#x60;

Everything seems to be working...

@jekyll/help Anyone have any ideas as to what&#x27;s going on here?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63721325) on: **Wednesday, November 19, 2014**

@ragwing Do you have this file: https://github.com/jekyll/jekyll/blob/master/lib/jekyll/deprecator.rb#L43 at &#x60;/var/lib/gems/2.1.0/gems/jekyll-2.5.1/lib/jekyll/deprecator.rb&#x60;?
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63729663) on: **Wednesday, November 19, 2014**

Yes I have the file. I will try to build it.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63729829) on: **Wednesday, November 19, 2014**

@ragwing Open it up. Does it have a &#x60;gracefully_require&#x60; method at the bottom?
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63732192) on: **Wednesday, November 19, 2014**

Yes it&#x27;s exactly the same as yours
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63732283) on: **Wednesday, November 19, 2014**

I tried to build again and it said everything is at the latest level
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63732888) on: **Wednesday, November 19, 2014**

@ragwing did you try uninstalling and re-installing the gem?
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63733398) on: **Wednesday, November 19, 2014**

Ok I am runing this on my laptop, so ssh ed to my desktop and installed everything installed correctly. I will go in there and run it see how it goes.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63733516) on: **Wednesday, November 19, 2014**

@ragwing can you do a &#x60;gem list&#x60; and paste the output
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63733782) on: **Wednesday, November 19, 2014**

On second thought, do you have a Gemfile for that project? if so can you run it with &#x60;bundle exec&#x60;
---
> from: [**ragwing**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63734726) on: **Wednesday, November 19, 2014**

Ok it worked fine on the desktop. So I guess something is wrong on the laptop. I will just use it on the desktop.
Thanks for the help

Ok I ran gem list on the desktop and the laptop
Desktop*** LOCAL GEMS ***

blankslate (2.1.2.4)
celluloid (0.16.0)
classifier-reborn (2.0.2)
coffee-script (2.3.0)
coffee-script-source (1.8.0)
colorator (0.1)
execjs (2.2.2)
fast-stemmer (1.0.2)
ffi (1.9.6)
hitimes (1.2.2)
jekyll (2.5.1)
jekyll-coffeescript (1.0.1)
jekyll-gist (1.1.0)
jekyll-paginate (1.1.0)
jekyll-sass-converter (1.2.1)
jekyll-watch (1.1.2)
kramdown (1.5.0)
liquid (2.6.1)
listen (2.8.0)
mercenary (0.3.5)
parslet (1.5.0)
posix-spawn (0.3.9)
pygments.rb (0.6.0)
rb-fsevent (0.9.4)
rb-inotify (0.9.5)
redcarpet (3.2.0)
safe_yaml (1.0.4)
sass (3.4.8)
timers (4.0.1)
toml (0.1.2)
yajl-ruby (1.1.0)

laptop
*** LOCAL GEMS ***

activesupport (3.2.20)
addressable (2.3.6)
afm (0.2.2)
Ascii85 (1.0.2)
bigdecimal (1.2.4)
blankslate (2.1.2.4)
builder (3.2.2)
bundler (1.6.3)
celluloid (0.16.0, 0.15.2)
childprocess (0.5.2)
classifier (1.3.4)
classifier-reborn (2.0.2)
coffee-script (2.3.0, 2.2.0)
coffee-script-source (1.8.0, 1.3.3)
colorator (0.1)
colored (1.2)
commander (4.1.6)
cucumber (1.3.11)
diff-lcs (1.2.5)
docile (1.1.5)
erubis (2.7.0)
ethon (0.7.1)
execjs (2.2.2, 2.2.1)
fast-stemmer (1.0.2)
ffi (1.9.6, 1.9.3)
gherkin (2.12.2)
gsl (1.15.3)
hashery (2.1.1)
highline (1.6.21)
hitimes (1.2.2)
html-proofer (1.5.4)
i18n (0.6.11, 0.6.9)
io-console (0.4.2)
jekyll (2.5.1, 2.0.3)
jekyll-coffeescript (1.0.1, 1.0.0)
jekyll-gist (1.1.0)
jekyll-paginate (1.1.0)
jekyll-sass-converter (1.2.1, 1.0.0)
jekyll-watch (1.1.2)
jekyll_test_plugin (0.1.0)
jekyll_test_plugin_malicious (0.1.0)
json (1.8.1)
kramdown (1.5.0, 1.4.0)
launchy (2.4.3)
liquid (2.6.1)
listen (2.8.0, 2.4.0)
log4r (1.1.10)
maruku (0.7.2, 0.7.1)
mercenary (0.3.5)
mime-types (1.25.1)
minitest (4.7.5)
msgpack (0.5.9)
multi_json (1.10.1, 1.10.0)
multi_test (0.1.1)
mysql (2.8.2)
narray (0.6.0.9)
net-http-persistent (2.9)
net-scp (1.1.1)
net-ssh (2.6.8)
nokogiri (1.6.3.1)
oj (2.9.4)
parallel (1.3.3)
parslet (1.6.1, 1.5.0)
pdf-core (0.2.5)
pdf-reader (1.3.3)
pg (0.13.2)
posix-spawn (0.3.9, 0.3.8)
prawn (1.0.0)
psych (2.0.5)
pygments.rb (0.6.0, 0.5.4)
rake (10.3.2, 10.1.0)
rb-fsevent (0.9.4)
rb-inotify (0.9.5)
rbtrace (0.4.5)
rdiscount (2.1.7.1, 1.6.8)
rdoc (4.1.0, 3.12.2)
redcarpet (3.2.0, 3.1.2)
RedCloth (4.2.9)
redgreen (1.2.2)
rouge (1.7.3, 1.5.1)
rr (1.1.2)
ruby-rc4 (0.1.5)
safe_yaml (1.0.4, 1.0.3)
sass (3.4.8, 3.2.19)
sequel (4.11.0)
sequel_pg (1.6.9)
shoulda (3.5.0)
shoulda-context (1.2.1)
shoulda-matchers (2.7.0)
simplecov (0.9.1)
simplecov-gem-adapter (1.0.1)
simplecov-html (0.8.0)
stackprof (0.2.7)
stringex (2.5.2)
test-unit (2.1.2.0)
thor (0.19.1)
timers (4.0.1, 1.1.0)
toml (0.1.2, 0.1.1)
trollop (2.0)
ttfunk (1.1.0)
typhoeus (0.6.9)
vagrant (1.4.3)
yajl-ruby (1.2.0, 1.1.0)
yell (2.0.5)

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-63736707) on: **Wednesday, November 19, 2014**

uninstall jekyll 2.0.3 and try again, maybe?
---
> from: [**jiajie999**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-67602912) on: **Thursday, December 18, 2014**

Same probnlem~~
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-67603107) on: **Thursday, December 18, 2014**

You need to use the latest version of Jekyll or the one GitHub Pages uses.
---
> from: [**jiajie999**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-67605223) on: **Thursday, December 18, 2014**

@parkr thanks for reply.
The problem is not jekyii, i just solve it by:  
sudo apt-get remove ruby* 
and reinstall jekyii



---
> from: [**tommyjiang**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-67636355) on: **Friday, December 19, 2014**

I encountered this problem today and solved it with @parkr &#x27;s advice which is uninstalling jekyll 2.0.3. Maybe this version is installed with Ruby I suppose.
Thank you, @parkr.
---
> from: [**luispedro**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-68161125) on: **Friday, December 26, 2014**

On my machine, I solved this problem by &#x60;&#x60;sudo apt-get remove jekyll; sudo apt-get autoremove ; sudo gem install jekyll&#x60;&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/200#issuecomment-68660282) on: **Sunday, January 04, 2015**

Oh, yeah. Definitely don&#x27;t install Jekyll using distro specific package managers like &#x60;apt-get&#x60; or &#x60;yum&#x60;.

Use &#x60;gem&#x60; to manage your Jekyll installation.
