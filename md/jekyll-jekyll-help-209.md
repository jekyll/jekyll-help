# [Installing Jekyll: ERROR: Failed to build gem native extension.](https://github.com/jekyll/jekyll-help/issues/209)

> state: **closed** opened by: **** on: ****

Hello,

Completely new to Jekyll, so apologies if this is a stupid question. After running 

&#x60;&#x60;gem install jekyll&#x60;&#x60;

I get the following error message:

&#x60;&#x60;Building native extensions.  This could take a while...&#x60;&#x60;
&#x60;&#x60;ERROR:  Error installing jekyll:&#x60;&#x60;
	&#x60;&#x60;ERROR: Failed to build gem native extension.&#x60;&#x60;

    &#x60;&#x60;/usr/bin/ruby2.1 -r ./siteconf20141210-5038-1u40f3s.rb extconf.rb&#x60;&#x60;
&#x60;&#x60;mkmf.rb can&#x27;t find header files for ruby at /usr/lib/ruby/include/ruby.h&#x60;&#x60;

&#x60;&#x60;extconf failed, exit code 1&#x60;&#x60;

&#x60;&#x60;Gem files will remain installed in /var/lib/gems/2.1.0/gems/fast-stemmer-1.0.2 for inspection.&#x60;&#x60;
&#x60;&#x60;Results logged to /var/lib/gems/2.1.0/extensions/x86_64-linux/2.1.0/fast-stemmer-1.0.2/gem_make.out&#x60;&#x60;

Seems like the commonest solution to this is to run

&#x60;&#x60;sudo apt-get install ruby1.9.1-dev&#x60;&#x60;

But this ain&#x27;t solving the problem for me -- anyone got any tips..?

### Comments

---
> from: [**alexwoolley**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-66441677) on: **Wednesday, December 10, 2014**

PS this is on Ubuntu 
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-66548197) on: **Wednesday, December 10, 2014**

Sounds like you are missing the &#x60;&#x60;&#x60;ruby-dev&#x60;&#x60;&#x60; package, try installing it with &#x60;&#x60;&#x60;sudo apt-get install ruby-dev&#x60;&#x60;&#x60; and see if it works then with the newest version of Ruby.
---
> from: [**redknitin**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-66724039) on: **Thursday, December 11, 2014**

@alexwoolley , as @sondr3 mentioned, the issue is most definitely because you are missing the ruby-dev package. You may also have to install other dependencies (Javascript and Python).

There is a jekyll package available for Ubuntu (look it up on Synaptic Package Manager) though I&#x27;m not sure of what version of jekyll it is. The package version is 0.11.2-1, which could possibly refer to Jekyll version 2.1 - that&#x27;s just a guess from the package version and I could be wrong.
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-66850156) on: **Friday, December 12, 2014**

Don&#x27;t use the Jekyll package on Ubuntu, it is incredibly old and not supported at all, people have previously been burnt by it, stay with the version you get from the gem.
---
> from: [**alexwoolley**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-66912809) on: **Sunday, December 14, 2014**

Thanks so much guys -- so sorry I&#x27;ve taken a while to reply!

All&#x27;s working now


---
> from: [**dhanvi**](https://github.com/jekyll/jekyll-help/issues/209#issuecomment-98558050) on: **Sunday, May 03, 2015**

:+1: for the issue this saved my day too :smile: 
