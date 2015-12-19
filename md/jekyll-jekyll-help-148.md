# [gem install jekyll fails to build after reinstalling gems](https://github.com/jekyll/jekyll-help/issues/148)

> state: **closed** opened by: **** on: ****

Because &#x60;gem update&#x60; failed on several gems with &#x60;ERROR: Failed to build gem native extension.&#x60; I tried to reinstall all gems (using Ruby 2.0.0 on Windows 7).

&#x60;&#x60;&#x60;
for i in &#x60;gem list --no-versions&#x60;; do gem uninstall -aIx $i; done
&#x60;&#x60;&#x60;

Now installing Jekyll fails as well:

&#x60;&#x60;&#x60;
$ gem install jekyll --verbose
HEAD https://api.rubygems.org/api/v1/dependencies
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=jekyll
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=classifier-reborn,colorat
r,jekyll-coffeescript,jekyll-gist,jekyll-paginate,jekyll-sass-converter,jekyll-
atch,kramdown,liquid,mercenary,pygments.rb,redcarpet,safe_yaml,toml
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=coffee-script,fast-stemme
,listen,parslet,posix-spawn,sass,yajl-ruby
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=blankslate,celluloid,coff
e-script-source,execjs,rb-fsevent,rb-inotify
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=ffi,timers
200 OK
GET https://api.rubygems.org/api/v1/dependencies?gems=hitimes
200 OK
Temporarily enhancing PATH to include DevKit...
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/LICENSE
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/README
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/Rakefile
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/VERSION.yml
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/ext/Makefile
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/ext/extconf.rb
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/ext/porter.c
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/ext/porter_wrap.c
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/lib/fast-stemmer.rb
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/lib/fast_stemmer.rb
c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast-stemmer-1.0.2/test/fast_stemmer_te
t.rb
Building native extensions.  This could take a while...
c:/Ruby200-x64/bin/ruby.exe -r ./siteconf20140913-6112-ztbkzl.rb extconf.rb
creating Makefile
make &quot;DESTDIR=&quot; clean
make &quot;DESTDIR=&quot;
ERROR:  Error installing jekyll:
        ERROR: Failed to build gem native extension.

    Building has failed. See above output for more information on the failure.
Building has failed. See above output for more information on the failure.
make failed, exit code 127

Gem files will remain installed in c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/fast
stemmer-1.0.2 for inspection.
Results logged to c:/Ruby200-x64/lib/ruby/gems/2.0.0/extensions/x64-mingw32/2.0
0/fast-stemmer-1.0.2/gem_make.out
&#x60;&#x60;&#x60;

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/148#issuecomment-55486770) on: **Saturday, September 13, 2014**

Okay, turns out reinstalling Ruby and the Ruby DevKit entirely resolves both issues I&#x27;ve had.
