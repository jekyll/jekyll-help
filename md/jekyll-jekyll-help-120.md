# [Error installing Jekyll - Ubuntu](https://github.com/jekyll/jekyll-help/issues/120)

> state: **closed** opened by: **** on: ****

When I try to install Jekyll I get this error:
&#x60;&#x60;&#x60;
 sudo gem install jekyll
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

        /usr/bin/ruby1.9.1 extconf.rb
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;: cannot load such file -- mkmf (LoadError)
	from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from extconf.rb:1:in &#x60;&lt;main&gt;&#x27;


Gem files will remain installed in /var/lib/gems/1.9.1/gems/fast-stemmer-1.0.2 for inspection.
Results logged to /var/lib/gems/1.9.1/gems/fast-stemmer-1.0.2/ext/gem_make.out
&#x60;&#x60;&#x60;
And here is a copy of the file that it links to:
&#x60;&#x60;&#x60;
/usr/bin/ruby1.9.1 extconf.rb
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;: cannot load such file -- mkmf (LoadError)
	from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from extconf.rb:1:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

Thanks in advance!


### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/120#issuecomment-52431318) on: **Sunday, August 17, 2014**

Try this before:

&#x60;&#x60;&#x60;
apt-get install ruby1.9.1-dev
&#x60;&#x60;&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/120#issuecomment-55646850) on: **Monday, September 15, 2014**

Closing this since I&#x27;m pretty sure that should have fixed it. @jdguy17, if you need more help, feel free to re-open this issue and let us know what&#x27;s wrong!
---
> from: [**thawn**](https://github.com/jekyll/jekyll-help/issues/120#issuecomment-91532309) on: **Friday, April 10, 2015**

I had the same error and can confirm that the suggestion from penibelst helped.
However, I did run apt-get build-dep jekyll before. Should this not have taken care of installing ruby1.9.1-dev? 
