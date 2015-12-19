# [Can&#x27;t Install Jekyll - Updated Ruby Gems](https://github.com/jekyll/jekyll-help/issues/137)

> state: **open** opened by: **** on: ****

Here is the error I receive on the command line in my terminal on OSX 10.6.8 .

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/137#issuecomment-53151740) on: **Saturday, August 23, 2014**

I think you forgot something. :)
---
> from: [**mriccelli75**](https://github.com/jekyll/jekyll-help/issues/137#issuecomment-53152616) on: **Saturday, August 23, 2014**

/System/Library/Frameworks/Ruby.framework/Versions/1.8/usr/bin/ruby -r ./siteconf20140823-291-ngzkuf-0.rb extconf.rb
mkmf.rb can&#x27;t find header files for ruby at /System/Library/Frameworks/Ruby.framework/Versions/1.8/usr/lib/ruby/ruby.h

---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/137#issuecomment-53512780) on: **Tuesday, August 26, 2014**

What do you see in terminal when you type &#x60;ruby -v&#x60;?

Look into installing a ruby version manager like rvm (with Homebrew or MacPorts) if you are running the system installed ruby 1.8.7.

AJ
---
> from: [**redknitin**](https://github.com/jekyll/jekyll-help/issues/137#issuecomment-65910948) on: **Saturday, December 06, 2014**

@mriccelli75 The error seems to indicate that you don&#x27;t have the needed Ruby Development files on your system. Do you have Xcode installed?
