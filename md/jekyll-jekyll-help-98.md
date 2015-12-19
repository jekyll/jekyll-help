# [Jekyll and Compass](https://github.com/jekyll/jekyll-help/issues/98)

> state: **closed** opened by: **** on: ****

I installed jekyll on Windows and everything works fine...

... until I tried to include compass. 

&#x60;&#x60;&#x60;SASS
@import &quot;compass&quot;
&#x60;&#x60;&#x60;

I get the error message &#x60;File to import not found or unreadable: compass. (Sass::SyntaxError)&#x60; But compass is installed!

Any ideas, why jekyll doesn&#x27;t find compass?


### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-49332601) on: **Thursday, July 17, 2014**

You have to add the Compass stylesheets path to your import path in order for this to work
---
> from: [**SimonWpt**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-49335416) on: **Thursday, July 17, 2014**

Something like this doesnt work:
&#x60;&#x60;&#x60;sass
@import &quot;C:/develop/Ruby200-x64/lib/ruby/gems/2.0.0/gems/compass-0.12.6/frameworks/compass/stylesheets/_compass.scss&quot;
&#x60;&#x60;&#x60;
I got the error message:
&#x60;&#x60;&#x60;
Conversion error: There was an error converting &#x27;css/screen.sass&#x27;.
jekyll 2.1.1 | Error:  File to import not found or unreadable: compass/reset/utilities.
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-49335940) on: **Thursday, July 17, 2014**

I meant in your config. It looks like [we&#x27;ve locked down import paths](https://github.com/jekyll/jekyll-sass-converter/blob/adb4bed65a6cee499fe57a692382aab88b68810c/lib/jekyll/converters/scss.rb#L54). We chose to go with Sass rather than Compass&#x27;s ruby libraries for Sass&#x27;s simplicity. I think we can open up import paths when not in safe mode. This would allow you to do:

&#x60;&#x60;&#x60;yaml
sass:
  sass_dir: _sass
  load_paths:
    - &quot;C:/develop/Ruby200-x64/lib/ruby/gems/2.0.0/gems/compass-0.12.6/frameworks/compass/stylesheets/&quot;
&#x60;&#x60;&#x60;

But it&#x27;s not currently available.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-49336048) on: **Thursday, July 17, 2014**

We&#x27;ll track this over in https://github.com/jekyll/jekyll-sass-converter/issues/13
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-49336076) on: **Thursday, July 17, 2014**

For now you can probably copy the relevant files to your &#x60;_sass&#x60; dir.
---
> from: [**ajkochanowicz**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-69097311) on: **Wednesday, January 07, 2015**

Where did this end up? Can one add muiltiple sass load_paths in Jekyll&#x27;s current state?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-69098304) on: **Wednesday, January 07, 2015**

&gt; Can one add muiltiple sass load_paths

Yes!
---
> from: [**ajkochanowicz**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-69099842) on: **Wednesday, January 07, 2015**

That&#x27;s good!

Is this documented somewhere? I&#x27;ve tried a few times but it&#x27;s not working for me.

I&#x27;ve also looked at PRs. They seem to indicate you just add a YAML array to &#x60;load_paths&#x60; in _config.yml. Is there more to it than that?

Thanks
---
> from: [**ajkochanowicz**](https://github.com/jekyll/jekyll-help/issues/98#issuecomment-69100275) on: **Wednesday, January 07, 2015**

Ah, I think I got it.

    sass:
      load_paths:
       - &quot;_sass&quot;
       - &quot;path1&quot;
       - &quot;path2&quot;
