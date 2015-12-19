# [Build error after upgrading to Yosemite](https://github.com/jekyll/jekyll-help/issues/203)

> state: **open** opened by: **** on: ****

I get the following error on &#x60;jekyll build&#x60; and have tried reinstalling Ruby, Rails, gems etc. 

WARN: Unresolved specs during Gem::Specification.reset:
      posix-spawn (~&gt; 0.3.6)
      ffi (&gt;= 0.5.0)
WARN: Clearing out unresolved specs.
Please report a bug if this causes problems.
/Users/milesgrimshaw/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.5.2/bin/jekyll:21:in &#x60;block in &lt;top (required)&gt;&#x27;: cannot load such file -- jekyll/version (LoadError)
	from /Users/milesgrimshaw/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary.rb:18:in &#x60;program&#x27;
	from /Users/milesgrimshaw/.rbenv/versions/2.1.3/lib/ruby/gems/2.1.0/gems/jekyll-2.5.2/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/203#issuecomment-64261095) on: **Monday, November 24, 2014**

Note Yosemite has new command line tools/Xcode and related cruft.
---
> from: [**milesgrimshaw**](https://github.com/jekyll/jekyll-help/issues/203#issuecomment-64261311) on: **Monday, November 24, 2014**

I think I updated all the Xcode stuff and GCC etc. 
---
> from: [**milesgrimshaw**](https://github.com/jekyll/jekyll-help/issues/203#issuecomment-64261664) on: **Monday, November 24, 2014**

&#x60;bundle exec serve&#x60; seems to work though
