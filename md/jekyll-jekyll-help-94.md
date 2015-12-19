# [Can’t update the gem](https://github.com/jekyll/jekyll-help/issues/94)

> state: **closed** opened by: **** on: ****

During my recent contribution to Jekyll I’ve installed the master branch version (2.0.3) with &#x60;bundle exec install&#x60;. This version is globally available on my system. Now I want to update it to the latest *stable*. I can’t figure out how. This is what I have tried:

&#x60;&#x60;&#x60;bash
$ ruby -v
ruby 2.0.0p384 (2014-01-12) [i386-linux-gnu]

$ jekyll -v
jekyll 2.0.3

$ gem update jekyll
Updating installed gems
Nothing to update

$ gem list jekyll

*** LOCAL GEMS ***

jekyll-coffeescript (1.0.0)
jekyll-gist (1.1.0)
jekyll-paginate (1.0.0)
jekyll-sass-converter (1.0.0)
jekyll_test_plugin (0.1.0)
jekyll_test_plugin_malicious (0.1.0)

$ ll /var/lib/gems/1.9.1/gems/jekyll-2.0.3/

drwxr-xr-x  9 root root  4096 Jun 20 20:28 ./
drwxr-xr-x 29 root root  4096 Jun 20 20:28 ../
drwxr-xr-x  2 root root  4096 Jun 20 20:28 bin/
-rw-r--r--  1 root root  3826 Jun 20 20:28 CONTRIBUTING.markdown
-rw-r--r--  1 root root   123 Jun 20 20:28 cucumber.yml
drwxr-xr-x  3 root root  4096 Jun 20 20:28 docs/
drwxr-xr-x  4 root root  4096 Jun 20 20:28 features/
-rw-r--r--  1 root root    38 Jun 20 20:28 Gemfile
-rw-r--r--  1 root root   132 Jun 20 20:28 .gitignore
-rw-r--r--  1 root root 51154 Jun 20 20:28 History.markdown
-rw-r--r--  1 root root  2680 Jun 20 20:28 jekyll.gemspec
drwxr-xr-x  4 root root  4096 Jun 20 20:28 lib/
-rw-r--r--  1 root root  1086 Jun 20 20:28 LICENSE
-rw-r--r--  1 root root  6899 Jun 20 20:28 Rakefile
-rw-r--r--  1 root root  3328 Jun 20 20:28 README.markdown
drwxr-xr-x  2 root root  4096 Jun 20 20:28 script/
drwxr-xr-x 12 root root  4096 Jun 20 20:28 site/
drwxr-xr-x  4 root root  4096 Jun 20 20:28 test/
-rw-r--r--  1 root root   842 Jun 20 20:28 .travis.yml
&#x60;&#x60;&#x60;

Jekyll doesn’t exist in the &#x60;/var/lib/gems/2.0.0/gems&#x60; directory.

:warning: After upgrading to Ubuntu 14.04 I did [restored Ruby 2.0 like suggested in this blog post](http://blog.costan.us/2014/04/restoring-ruby-20-on-ubuntu-1404.html). Maybe, I don’t remember, I did it *after* I installed Jekyll from the master branch. This may cause my issue.

Any suggestions?

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48427063) on: **Tuesday, July 08, 2014**

Hey @penibelst did you tried updating the gem in the Gemfile definition and then &#x60;bundle update jekyll&#x60;?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48436213) on: **Tuesday, July 08, 2014**

@albertogg This is the output:

&#x60;&#x60;&#x60;bash

$ cd jekyll
$ git pull upstream master
# Jekyll::VERSION switches to 2.1.0

$ bundle update jekyll
…
Using jekyll (2.1.0) from source at .
…
Your bundle is updated!

$ jekyll -v
jekyll 2.0.3
&#x60;&#x60;&#x60;

No luck.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48485451) on: **Wednesday, July 09, 2014**

Try this?

&#x60;&#x60;&#x60;
$ cd jekyll
$ git pull upstream master
$ bundle
$ bundle exec rake build
$ gem install -l pkg/jekyll-2.1.0.gem
&#x60;&#x60;&#x60;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48493243) on: **Wednesday, July 09, 2014**

@penibelst and what if you &#x60;bundle exec jekyll -v&#x60; ? it should work. Also, if you &#x60;gem list --local | grep jekyll&#x60; do you see jekyll latest stable version installed along with the previous ones?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48518208) on: **Wednesday, July 09, 2014**

@parkr 
&#x60;&#x60;&#x60;bash
$ sudo gem install -l pkg/jekyll-2.1.0.gem
ERROR:  Error installing pkg/jekyll-2.1.0.gem:
	ERROR: Failed to build gem native extension.
&#x60;&#x60;&#x60;

@albertogg 

&#x60;&#x60;&#x60;bash
$ bundle exec jekyll -v
jekyll 2.1.0

$ jekyll -v
jekyll 2.0.3

$ gem list --local | grep jekyll
jekyll-coffeescript (1.0.0)
jekyll-gist (1.1.0)
jekyll-paginate (1.0.0)
jekyll-sass-converter (1.0.0)
jekyll-watch (1.0.0)
jekyll_test_plugin (0.1.0)
jekyll_test_plugin_malicious (0.1.0)
&#x60;&#x60;&#x60;


---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48519146) on: **Wednesday, July 09, 2014**

@pwnall, I think my issue is related to your [post about restoring Ruby 2.0](http://blog.costan.us/2014/04/restoring-ruby-20-on-ubuntu-1404.html).
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48521025) on: **Wednesday, July 09, 2014**

Try &#x60;which jekyll&#x60;, &#x60;which jekyll-gist&#x60; and compare the output. Also related to the link you are showing is the symlink for the &#x60;/usr/bin/gem&#x60; working? what happens when you try to run &#x60;/usr/bin/gem2.0 list --local&#x60;? 
---
> from: [**pwnall**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48547563) on: **Wednesday, July 09, 2014**

@parkr  Your comment above shows that when you do &#x60;sudo gem install -l pkg/jekyll-2.1.0.gem&#x60; it says &#x60;ERROR: Failed to build gem native extension.&#x60; I think that&#x27;s what&#x27;s causing your troubles.

Rubygems usually gives you a link to an error log. If you gist it and post a link, I can try to take a look.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48570698) on: **Wednesday, July 09, 2014**

@albertogg 

&#x60;&#x60;&#x60;bash
$ which jekyll
/usr/local/bin/jekyll

$ which jekyll-gist
# nothing

$ ll /usr/bin/gem
lrwxrwxrwx 1 root root 15 Jun 20 20:35 /usr/bin/gem -&gt; /usr/bin/gem2.0

$ /usr/bin/gem2.0 list --local

*** LOCAL GEMS ***

# …
i18n (0.6.9)
io-console (0.4.2)
jekyll-coffeescript (1.0.0)
jekyll-gist (1.1.0)
# …
&#x60;&#x60;&#x60;

@pwnall It’s not @parkr, but me who has the issue.

&gt; Rubygems usually gives you a link to an error log.

I can’t find the link.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48625273) on: **Thursday, July 10, 2014**

@penibelst have you tried removing jekyll and installing it again?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48628691) on: **Thursday, July 10, 2014**

@albertogg Removing &#x60;/var/lib/gems/1.9.1/gems/jekyll-2.0.3/&#x60; with &#x60;rm -r&#x60;? No.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48630127) on: **Thursday, July 10, 2014**

@penibelst no no, using &#x60;gem uninstall&#x60;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48630384) on: **Thursday, July 10, 2014**

@albertogg 

&#x60;&#x60;&#x60;bash
$ gem uninstall jekyll
# nothing happens
$ jekyll -v
jekyll 2.0.3
&#x60;&#x60;&#x60;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48799630) on: **Friday, July 11, 2014**

@penibelst can you track where is this jekyll gem installed? and then check if there are other gems too. If you install any other gem does the gem list command shows it? 
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48805608) on: **Saturday, July 12, 2014**

@albertogg Turns out I can’t install gems anymore:

&#x60;&#x60;&#x60;bash
$ sudo gem install html-proofer
ERROR:  Error installing html-proofer:
	ERROR: Failed to build gem native extension.

    /usr/bin/ruby2.0 extconf.rb
checking for ffi.h... no
checking for ffi.h in /usr/local/include,/usr/include/ffi... no
checking for rb_thread_blocking_region()... yes
checking for rb_thread_call_with_gvl()... yes
checking for rb_thread_call_without_gvl()... yes
checking for ffi_prep_cif_var()... no
creating extconf.h
creating Makefile

make &quot;DESTDIR=&quot; clean

make &quot;DESTDIR=&quot;
make -C &quot;/var/lib/gems/2.0.0/gems/ffi-1.9.3/ext/ffi_c/libffi-i386-linux-gnu&quot;

# more errors

make failed, exit code 2

Gem files will remain installed in /var/lib/gems/2.0.0/gems/ffi-1.9.3 for inspection.
Results logged to /var/lib/gems/2.0.0/extensions/x86-linux/2.0.0/ffi-1.9.3/gem_make.out
&#x60;&#x60;&#x60;

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48805805) on: **Saturday, July 12, 2014**

Can I delete *all* gems inside &#x60;/var/lib/gems/&#x60; and just start again?
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/94#issuecomment-48838638) on: **Sunday, July 13, 2014**

@pwnall You are right, native extensions were the problem. After reinstalling some gems with native extensions everything works as expected.
