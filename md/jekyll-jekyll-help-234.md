# [Help... with errors](https://github.com/jekyll/jekyll-help/issues/234)

> state: **open** opened by: **** on: ****

Tried build the site, and I got this errors..

WARN: Unresolved specs during Gem::Specification.reset:
      redcarpet (~&gt; 3.1)
      jekyll-watch (~&gt; 1.1)
      classifier-reborn (~&gt; 2.0)
WARN: Clearing out unresolved specs.
Please report a bug if this causes problems.
/Users/name.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.5.3/bin/jekyll:21:in &#x60;block in &lt;top (required)&gt;&#x27;: cannot load such file -- jekyll/version (LoadError)
	from /Users/name.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary.rb:18:in &#x60;program&#x27;
	from /Users/name/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.5.3/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/name/.rbenv/versions/2.1.2/bin/jekyll:23:in &#x60;load&#x27;
	from /Users/name/.rbenv/versions/2.1.2/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;


### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/234#issuecomment-69382945) on: **Friday, January 09, 2015**

Not entirely sure, but maybe run &#x60;&#x60;&#x60;bundle update&#x60;&#x60;&#x60; and see if that helps.
---
> from: [**luckyluke007**](https://github.com/jekyll/jekyll-help/issues/234#issuecomment-69384629) on: **Friday, January 09, 2015**

Ran the update, It started up and builded fine, but still get this warning. 

WARN: Unresolved specs during Gem::Specification.reset:
      redcarpet (~&gt; 3.1)
      jekyll-watch (~&gt; 1.1)
      classifier-reborn (~&gt; 2.0)

Ran &quot;bundle exec jekyll serve&quot; to eliminate the WARN..

Server address is change to 127.0.0.1 instead 0.0.0.0: 4000.. Would VPN set up change the server address?

Thanks...


---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/234#issuecomment-69385178) on: **Friday, January 09, 2015**

wonder if you have &#x60;&#x60;&#x60;github-pages&#x60;&#x60;&#x60; in your gemfile. I see you have the latest Jekyll, but but Github Pages is still on 2.4 https://pages.github.com/versions/
