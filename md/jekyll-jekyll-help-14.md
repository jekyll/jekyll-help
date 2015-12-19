# [had problems with Quick Start Guide until...](https://github.com/jekyll/jekyll-help/issues/14)

> state: **closed** opened by: **** on: ****

Just wanted to note (in case others have similar issues) that I had to modify the instructions slightly on the quick-start guide page (http://jekyllrb.com/docs/quickstart/) to get it to work for me.

Instead of:
~ $ gem install jekyll
~ $ jekyll new myblog
~ $ cd myblog
~/myblog $ jekyll serve
 =&gt; Now browse to http://localhost:4000

I had to do:
~ $ sudo gem install jekyll
~ $ sudo gem install json
~ $ jekyll new myblog
~ $ cd myblog
~/myblog $ jekyll serve
 =&gt; Now browse to http://localhost:4000

As you can see, very minor differences.  Also, I got errors when the Jekyll documentation was being installed, but it still worked fine.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/14#issuecomment-40088211) on: **Thursday, April 10, 2014**

The necessity of using &#x60;sudo&#x60; for those commands depends completely upon how you have your version of Ruby installed on your system. The goal should be to not have to use &#x60;sudo&#x60; to install gems, since that is potentially very dangerous. (Ruby will execute the code as root, giving that code access to do anything at all.)
