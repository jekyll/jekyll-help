# [Jekyll Not Auto Generating while using “serve -w”](https://github.com/jekyll/jekyll-help/issues/69)

> state: **closed** opened by: **** on: ****

I also posted about this issue at [Stackoverflow](http://stackoverflow.com/questions/24101378/jekyll-not-auto-generating-while-using-serve-w).

I&#x27;m able to run jekyll serve and jekyll serve -w without any output errors. It generates and serves the site fine. Though as I edit page the changes aren&#x27;t automatically detected and are only output on to the development server after I run jekyll serve again.

I was messing around trying to install Nokogiri and ended up reinstalling rvm recently I was wondering if that is playing a role?

I&#x27;ve tried Jekyll serve on a clean install and in a few different versions of Ruby. This is probably a quick fix, but the answer is way out of my realm of knowledge.

**Using**
- Jekyll 2.0.3
- ruby 2.1.2p95 (2014-05-08 revision 45877) [x86_64-darwin13.0]
- OS X 10.9.3

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/69#issuecomment-45443094) on: **Sunday, June 08, 2014**

What are your &#x60;source&#x60; and &#x60;destination&#x60;?
---
> from: [**ericthor**](https://github.com/jekyll/jekyll-help/issues/69#issuecomment-45458328) on: **Sunday, June 08, 2014**

I was able to get it to work by moving it out of the default &quot;Sites&quot; folder in OSX. After I moved it out of that folder it worked fine. Its definitely an issue not with the Jekyll code. I resolved it as of now.

    source: ./_source
    destination: ./_site


---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/69#issuecomment-45458530) on: **Sunday, June 08, 2014**

Glad you got it resolved!
