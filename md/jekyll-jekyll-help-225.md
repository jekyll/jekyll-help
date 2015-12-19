# [jekyll vs bundle exec jekyll](https://github.com/jekyll/jekyll-help/issues/225)

> state: **closed** opened by: **** on: ****

Hi

THis is more of a question rather an issue (I am using &#x60;bundle exec jekyll &lt;command&gt;&#x60; as per the GitHub pages instructions) but what does &#x60;bundler&#x60; do in relation to &#x60;jekyll&#x60;, and what is the difference between, say, &#x60;jekyll [build | serve]&#x60; vs &#x60;bundle exec jekyll [build | serve]&#x60;?

Right now I am able to build and serve the site locally with &#x60;bundle exec jekyll [build | serve]&#x60;, but when I &#x60;jekyll [build | serve]&#x60; it throws the following error

    /Library/Ruby/Gems/2.0.0/gems/safe_yaml-1.0.4/lib/safe_yaml/psych_resolver.rb:4:in &#x60;&lt;class:PsychResolver&gt;&#x27;: uninitialized constant Psych::Nodes (NameError)
	from /Library/Ruby/Gems/2.0.0/gems/safe_yaml-1.0.4/lib/safe_yaml/psych_resolver.rb:2:in &#x60;&lt;module:SafeYAML&gt;&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/safe_yaml-1.0.4/lib/safe_yaml/psych_resolver.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/safe_yaml-1.0.4/lib/safe_yaml/load.rb:131:in &#x60;&lt;module:SafeYAML&gt;&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/safe_yaml-1.0.4/lib/safe_yaml/load.rb:26:in &#x60;&lt;top (required)&gt;&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.3/lib/jekyll.rb:26:in &#x60;&lt;top (required)&gt;&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Site/2.0.0/rubygems/core_ext/kernel_require.rb:69:in &#x60;require&#x27;
	from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.3/bin/jekyll:6:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;


### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/225#issuecomment-68932031) on: **Tuesday, January 06, 2015**

[Bundler&#x27;s Purpose and Rationale](http://bundler.io/rationale.html)

---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/225#issuecomment-68960392) on: **Tuesday, January 06, 2015**

Thank you, this is very clear :)
