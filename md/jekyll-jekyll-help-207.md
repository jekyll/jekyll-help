# [Installation ](https://github.com/jekyll/jekyll-help/issues/207)

> state: **open** opened by: **** on: ****

Hey, 
i get the following error while install  jekyll via gem:

ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

        /usr/bin/ruby1.9.1 extconf.rb
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;: cannot load such file -- mkmf (LoadError)
	from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from extconf.rb:1:in &#x60;&lt;main&gt;&#x27;


Thanks
Fabian

### Comments

---
> from: [**redknitin**](https://github.com/jekyll/jekyll-help/issues/207#issuecomment-65910702) on: **Saturday, December 06, 2014**

@FabianCernota You are probably just missing the Ruby Dev package. You can install it on Ubuntu or other Debian-based Linux by typing in a terminal window: &#x60;sudo apt-get install ruby-dev&#x60; . On a Fedora/RedHat-based Linux, it would be a &#x60;sudo yum install ruby-devel&#x60;

After installing Ruby Dev, attempt to install the Jekyll gem again.
---
> from: [**bradhowes**](https://github.com/jekyll/jekyll-help/issues/207#issuecomment-67150158) on: **Tuesday, December 16, 2014**

You will also need &#x27;make&#x27; and &#x27;gcc&#x27;.
---
> from: [**KitN**](https://github.com/jekyll/jekyll-help/issues/207#issuecomment-72155122) on: **Thursday, January 29, 2015**

Thanks, bradhowes. I had trouble with &#x60;gem install jekyll&#x60; but &#x60;apt-get install make&#x60; worked for me on Debian.
