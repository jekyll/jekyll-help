# [FreeBSD: Could not find kramdown](https://github.com/jekyll/jekyll-help/issues/46)

> state: **closed** opened by: **** on: ****

Hi,
When I try to run jekyll &#x60;new bloglocation&#x60; on FreeBSD 10.0, I get this error:
&#x60;&#x60;&#x60;
/usr/local/lib/ruby/site_ruby/2.1/rubygems/dependency.rb:247:in &#x60;to_specs&#x27;: Could not find kramdown (~&gt; 1.0.2) amongst [classifier-1.3.3, colorator-0.1, commander-4.1.4, directory_watcher-1.4.1, fast-stemmer-1.0.2, highline-1.6.21, jekyll-1.0.3, kramdown-1.3.3, liquid-2.5.0, maruku-0.6.1, posix-spawn-0.3.6, pygments.rb-0.5.4, safe_yaml-0.9.4, syntax-1.0.0, yajl-ruby-1.2.0] (Gem::LoadError)
        from /usr/local/lib/ruby/site_ruby/2.1/rubygems/specification.rb:778:in &#x60;block in activate_dependencies&#x27;
        from /usr/local/lib/ruby/site_ruby/2.1/rubygems/specification.rb:767:in &#x60;each&#x27;
        from /usr/local/lib/ruby/site_ruby/2.1/rubygems/specification.rb:767:in &#x60;activate_dependencies&#x27;
        from /usr/local/lib/ruby/site_ruby/2.1/rubygems/specification.rb:751:in &#x60;activate&#x27;
        from /usr/local/lib/ruby/site_ruby/2.1/rubygems.rb:1232:in &#x60;gem&#x27;
        from /usr/local/bin/jekyll:22:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;


Ruby says I have &#x60;kramdown&#x60; on my computer, so why is it giving an error that &#x60;kramdown&#x60; isn&#x27;t installed on my computer? I installed Jekyll from FreeBSD ports, and set the default Ruby version to &#x60;2.1&#x60;.
BTW I have posted this issue on [Jekyll&#x27;s main GitHub repo](https://github.com/jekyll/jekyll/issues/2403), but @parktr said to come here.

### Comments

---
> from: [**JesseHerrick**](https://github.com/jekyll/jekyll-help/issues/46#issuecomment-44323833) on: **Tuesday, May 27, 2014**

What version of Kramdown do you have? &#x60;kramdown -v&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/46#issuecomment-44333721) on: **Tuesday, May 27, 2014**

It&#x27;s looking for Kramdown ~&gt; 1.0.2 (&gt;= 1.0.2, &lt; 1.1), but you have 1.3.3 installed.
---
> from: [**neelchauhan**](https://github.com/jekyll/jekyll-help/issues/46#issuecomment-44342681) on: **Tuesday, May 27, 2014**

@parkr I have a question.
Why can&#x27;t Jekyll work with Kramdown 1.3.3? I know how to code, but I don&#x27;t know Ruby (have no interest, to the extent that I&#x27;d rather learn assembly language!)
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/46#issuecomment-44344253) on: **Tuesday, May 27, 2014**

You&#x27;re running an old version of Jekyll there. The latest gemspec is &#x60;~&gt; 1.3&#x60;, which would include anything above 1.3.0 before 2.0.0.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/46#issuecomment-48134263) on: **Sunday, July 06, 2014**

Going to close this one out since it looks like @parkr&#x27;s solution ought to work. Feel free to add another comment if you need more help, or just open a new issue if something new comes up.
