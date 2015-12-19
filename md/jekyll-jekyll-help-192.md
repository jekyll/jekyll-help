# [when running jekyll serve ,error occurred,hitimes/hitimes (LoadError)](https://github.com/jekyll/jekyll-help/issues/192)

> state: **closed** opened by: **** on: ****

Hi ,
I&#x27;m the greener to jekyll,after I installed Jekyll and run for it,error happened...
According to the clues, I installed hittimes,but the problems still exists.
&#x60;&#x60;&#x60;
/Users/supercafe/my-awsome-site
supercafes-MacBook-Pro:my-awsome-site supercafe$ jekyll serve
Configuration file: /Users/supercafe/my-awsome-site/_config.yml
            Source: /Users/supercafe/my-awsome-site
       Destination: /Users/supercafe/my-awsome-site/_site
      Generating...
                    done.
/usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;: cannot load such file -- hitimes/hitimes (LoadError)
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/gems/2.0.0/gems/hitimes-1.2.2/lib/hitimes.rb:37:in &#x60;rescue in &lt;top (required)&gt;&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/gems/2.0.0/gems/hitimes-1.2.2/lib/hitimes.rb:32:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/gems/2.0.0/gems/timers-4.0.1/lib/timers/group.rb:4:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /usr/local/Cellar/ruby/2.0.0-p0/lib/ruby/gems/2.0.0/gems/timers-4.0.1/lib/timers.rb:4:in &#x60;&lt;top (required)&gt;&#x27;

&#x60;&#x60;&#x60;




### Comments

---
> from: [**AlejandroCruz**](https://github.com/jekyll/jekyll-help/issues/192#issuecomment-70382823) on: **Saturday, January 17, 2015**

I read that you were successful in resolving this issue in https://github.com/jekyll/jekyll/issues/3092 , how did you do it. Did you uninstall and reinstall hitimes and timers specifying versions? I tried the uninstall of both gems and then installing them and it did not resolve my issue. I was wondering if I missed a step or something.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/192#issuecomment-70395996) on: **Saturday, January 17, 2015**

It seems like, I&#x27;ll close the issue! thanks.
